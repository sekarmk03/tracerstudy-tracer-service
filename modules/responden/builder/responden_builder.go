package builder

import (
	"tracerstudy-tracer-service/common/config"
	mhsbSvc "tracerstudy-tracer-service/modules/mhsbiodata/service"
	"tracerstudy-tracer-service/modules/responden/handler"
	"tracerstudy-tracer-service/modules/responden/repository"
	respSvc "tracerstudy-tracer-service/modules/responden/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildRespondenHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.RespondenHandler {
	respondenRepo := repository.NewRespondenRepository(db)
	respondenSvc := respSvc.NewRespondenService(cfg, respondenRepo)

	mhsbiodataSvc := mhsbSvc.NewMhsBiodataService(cfg)

	return handler.NewRespondenHandler(cfg, respondenSvc, mhsbiodataSvc)
}
