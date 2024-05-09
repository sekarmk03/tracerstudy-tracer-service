package mhsbiodata

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/mhsbiodata/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	mhsbiodata := builder.BuildMhsBiodataHandler(cfg, grpcConn)
	pb.RegisterMhsBiodataServiceServer(server, mhsbiodata)
}
