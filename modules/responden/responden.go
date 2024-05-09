package responden

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/responden/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	responden := builder.BuildRespondenHandler(cfg, db, grpcConn)
	pb.RegisterRespondenServiceServer(server, responden)
}
