package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/mhsbiodata/handler"
	"tracerstudy-tracer-service/modules/mhsbiodata/service"

	"google.golang.org/grpc"
)

func BuildMhsBiodataHandler(cfg config.Config, grpcConn *grpc.ClientConn) *handler.MhsBiodataHandler {
	mhsbiodataSvc := service.NewMhsBiodataService(cfg)
	return handler.NewMhsBiodataHandler(cfg, mhsbiodataSvc)
}
