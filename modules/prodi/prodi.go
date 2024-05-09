package prodi

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/prodi/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	prodi := builder.BuildProdiHandler(cfg, db, grpcConn)
	pb.RegisterProdiServiceServer(server, prodi)
}
