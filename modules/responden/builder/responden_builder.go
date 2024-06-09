package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/responden/client"
	"tracerstudy-tracer-service/modules/responden/handler"
	"tracerstudy-tracer-service/modules/responden/repository"
	respSvc "tracerstudy-tracer-service/modules/responden/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildRespondenHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.RespondenHandler {
	respondenRepo := repository.NewRespondenRepository(db)
	respondenSvc := respSvc.NewRespondenService(cfg, respondenRepo)

	mhsApiSvc := client.BuildMhsBiodataServiceClient(cfg.ClientURL.MhsBiodata)

	return handler.NewRespondenHandler(cfg, respondenSvc, mhsApiSvc)
}
