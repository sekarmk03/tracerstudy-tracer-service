package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/pkts/handler"
	"tracerstudy-tracer-service/modules/pkts/repository"
	"tracerstudy-tracer-service/modules/pkts/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildPKTSHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.PKTSHandler {
	pktsRepo := repository.NewPKTSRepository(db)
	pktsSvc := service.NewPKTSService(cfg, pktsRepo)

	return handler.NewPKTSHandler(cfg, pktsSvc)
}
