package interceptor

import (
	"context"
	"log"

	commonJwt "tracerstudy-tracer-service/common/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	jwtManager      *commonJwt.JWT
	accessibleRoles map[string][]uint32
}

func NewAuthInterceptor(jwtManager *commonJwt.JWT, accessibleRoles map[string][]uint32) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager:      jwtManager,
		accessibleRoles: accessibleRoles,
	}
}

func (a *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		log.Println("INFO [Auth Interceptor - Unary Server Interceptor] Method:", info.FullMethod)

		if err := a.authorize(ctx, info.FullMethod); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (a *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := a.accessibleRoles[method]
	if !ok {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("ERROR: [Auth Interceptor - Authorize] Metadata is not provided")
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values, ok := md["authorization"]
	if !ok || len(values) == 0 {
		log.Println("ERROR: [Auth Interceptor - Authorize] Authorization token is not provided")
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := a.jwtManager.Verify(accessToken)
	if err != nil {
		log.Println("ERROR: [Auth Interceptor - Authorize] Access token is invalid:", err)
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	for _, role := range accessibleRoles {
		if role == claims.Role {
			return nil
		}
	}

	log.Println("ERROR: [Auth Interceptor - Authorize] No permission to access this RPC")
	return status.Errorf(codes.PermissionDenied, "no permission to access this RPC")
}
