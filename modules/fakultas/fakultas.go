package fakultas

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/fakultas/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	fakultas := builder.BuildFakultasHandler(cfg, db, grpcConn)
	pb.RegisterFakultasServiceServer(server, fakultas)
}
