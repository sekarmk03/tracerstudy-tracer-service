package pkts

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/pkts/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	pkts := builder.BuildPKTSHandler(cfg, db, grpcConn)
	pb.RegisterPKTSServiceServer(server, pkts)
}
