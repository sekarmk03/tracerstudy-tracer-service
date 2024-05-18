package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/mhsbiodataapi/handler"
	"tracerstudy-tracer-service/modules/mhsbiodataapi/service"

	"google.golang.org/grpc"
)

func BuildMhsBiodataApiHandler(cfg config.Config, grpcConn *grpc.ClientConn) *handler.MhsBiodataApiHandler {
	mhsbiodataapiSvc := service.NewMhsBiodataApiService(cfg)
	return handler.NewMhsBiodataHandler(cfg, mhsbiodataapiSvc)
}
