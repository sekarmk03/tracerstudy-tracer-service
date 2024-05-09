package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/prodi/handler"
	"tracerstudy-tracer-service/modules/prodi/repository"
	"tracerstudy-tracer-service/modules/prodi/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildProdiHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.ProdiHandler {
	prodiRepo := repository.NewProdiRepository(db)
	prodiSvc := service.NewProdiService(cfg, prodiRepo)

	return handler.NewProdiHandler(cfg, prodiSvc)
}
