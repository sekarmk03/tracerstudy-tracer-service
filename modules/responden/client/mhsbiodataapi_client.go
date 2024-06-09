package client

import (
	"context"
	"tracerstudy-tracer-service/pb"
	"tracerstudy-tracer-service/server"

	"google.golang.org/grpc/metadata"
)

type MhsBiodataApiServiceClient struct {
	Client pb.MhsBiodataApiServiceClient
}

func BuildMhsBiodataServiceClient(url string) MhsBiodataApiServiceClient {
	cc := server.InitGRPCConn(url, false, "")

	c := MhsBiodataApiServiceClient{
		Client: pb.NewMhsBiodataApiServiceClient(cc),
	}

	return c
}

func (mc *MhsBiodataApiServiceClient) FetchMhsBiodataByNim(ctx context.Context, token, nim string) (*pb.MhsBiodataApiResponse, error) {
	md := metadata.New(map[string]string{"authorization": token})
	ctxNew := metadata.NewOutgoingContext(ctx, md)

	req := &pb.MhsBiodataApiRequest{
		Nim: nim,
	}

	return mc.Client.FetchMhsBiodataByNim(ctxNew, req)
}
