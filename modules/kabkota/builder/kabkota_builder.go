package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/kabkota/handler"
	"tracerstudy-tracer-service/modules/kabkota/repository"
	"tracerstudy-tracer-service/modules/kabkota/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildKabKotaHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.KabKotaHandler {
	kabkotaRepo := repository.NewKabKotaRepository(db)
	kabkotaSvc := service.NewKabKotaService(cfg, kabkotaRepo)

	return handler.NewKabKotaHandler(cfg, kabkotaSvc)
}
