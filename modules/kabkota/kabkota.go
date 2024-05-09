package kabkota

import (
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/modules/kabkota/builder"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitGrpc(server *grpc.Server, cfg config.Config, db *gorm.DB, grpcConn *grpc.ClientConn) {
	kabkota := builder.BuildKabKotaHandler(cfg, db, grpcConn)
	pb.RegisterKabKotaServiceServer(server, kabkota)
}
