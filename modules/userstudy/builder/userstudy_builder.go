package builder

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/userstudy/handler"
	"tracerstudy-tracer-service/modules/userstudy/repository"
	"tracerstudy-tracer-service/modules/userstudy/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func BuildUserStudyHandler(cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) *handler.UserStudyHandler {
	userStudyRepo := repository.NewUserStudyRepository(db)
	userStudySvc := service.NewUserStudyService(cfg, userStudyRepo)

	return handler.NewUserStudyHandler(cfg, userStudySvc)
}
