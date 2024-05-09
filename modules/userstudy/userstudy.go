package userstudy

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/userstudy/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	userstudy := builder.BuildUserStudyHandler(cfg, db, grpcConn)
	pb.RegisterUserStudyServiceServer(server, userstudy)
}
