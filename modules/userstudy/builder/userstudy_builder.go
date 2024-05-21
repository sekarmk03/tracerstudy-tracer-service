package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/userstudy/handler"
	"tracerstudy-tracer-service/modules/userstudy/repository"
	"tracerstudy-tracer-service/modules/userstudy/service"
	pkRepo "tracerstudy-tracer-service/modules/pkts/repository"
	pkSvc "tracerstudy-tracer-service/modules/pkts/service"
	respRepo "tracerstudy-tracer-service/modules/responden/repository"
	respSvc "tracerstudy-tracer-service/modules/responden/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildUserStudyHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.UserStudyHandler {
	userStudyRepo := repository.NewUserStudyRepository(db)
	userStudySvc := service.NewUserStudyService(cfg, userStudyRepo)
	pktsRepo := pkRepo.NewPKTSRepository(db)
	pktsSvc := pkSvc.NewPKTSService(cfg, pktsRepo)
	respondenRepo := respRepo.NewRespondenRepository(db)
	respondenSvc := respSvc.NewRespondenService(cfg, respondenRepo)

	return handler.NewUserStudyHandler(cfg, userStudySvc, respondenSvc, pktsSvc)
}
