package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	// commonJwt "tracerstudy-tracer-service/common/jwt"
	// roles "tracerstudy-tracer-service/common/authorization"
	// "tracerstudy-tracer-service/server/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	connProtocol  = "tcp"
	maxMsgSize    = 1024 * 1024 * 150
	tokenDuration = 5 * time.Minute
	secretKey     = "secret"
)

type Grpc struct {
	Server   *grpc.Server
	listener net.Listener
	Port     string
}

func NewGrpc(port string, options ...grpc.ServerOption) *Grpc {
	options = append(options, grpc.MaxSendMsgSize(maxMsgSize))
	options = append(options, grpc.MaxRecvMsgSize(maxMsgSize))

	server := grpc.NewServer(options...)

	return &Grpc{
		Server: server,
		Port:   port,
	}
}

func NewGrpcServer(
	port string,
	// jwtManager *commonJwt.JWT,
) *Grpc {
	// var options grpc.ServerOption
	// options := grpc_middleware.WithUnaryServerChain()
	// add option unary interceptor
	// jwtManager := commonJwt.NewJWT(secretKey, tokenDuration)
	// authInterceptor := interceptor.NewAuthInterceptor(jwtManager, accessibleRoles())
	// authInterceptor := interceptor.NewAuthInterceptor(jwtManager, roles.GetAccessibleRoles())
	options := []grpc.ServerOption{
		// grpc.UnaryInterceptor(authInterceptor.Unary()),
	}
	server := NewGrpc(port, options...)
	return server
}

func (g *Grpc) Run() error {
	var err error
	g.listener, err = net.Listen(connProtocol, fmt.Sprintf(":%s", g.Port))
	if err != nil {
		return status.Errorf(codes.Internal, "ERROR: Failed to listen on port %s: %v", g.Port, err)
	}

	go g.serve()
	log.Printf("grpc server is running on port %s\n", g.Port)
	return nil
}

func (g *Grpc) serve() {
	if err := g.Server.Serve(g.listener); err != nil {
		panic(err)
	}
}

func (g *Grpc) AwaitTermination() error {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	<-sign

	g.Server.GracefulStop()
	return g.listener.Close()
}

// func defaultUnaryServerInterceptor() []grpc.UnaryServerInterceptor {
// 	jwtManager := commonJwt.NewJWT(secretKey, tokenDuration)
// 	authInterceptor := interceptor.NewAuthInterceptor(jwtManager, accessibleRoles())

// 	options := []grpc.ServerOption{
// 		grpc.UnaryInterceptor(authInterceptor.Unary()),
// 	}

// 	return options
// }

// func accessibleRoles() map[string][]uint32 {
// 	const prodiService = "/tracer_study_grpc.ProdiService/"
// 	// 1 = admin
// 	// 2 = user
// 	return map[string][]uint32{
// 		prodiService + "GetAllProdi": {1, 2},
// 		prodiService + "GetProdiByKodeprodi": {1, 2},
// 		prodiService + "CreateProdi": {1},
// 	}
// }
