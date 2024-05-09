package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/fakultas/handler"
	"tracerstudy-tracer-service/modules/fakultas/repository"
	"tracerstudy-tracer-service/modules/fakultas/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildFakultasHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.FakultasHandler {
	fakultasRepo := repository.NewFakultasRepository(db)
	fakultasSvc := service.NewFakultasService(cfg, fakultasRepo)

	return handler.NewFakultasHandler(cfg, fakultasSvc)
}
