package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/modules/mhsbiodata/entity"
	"tracerstudy-tracer-service/modules/mhsbiodata/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MhsBiodataHandler struct {
	pb.UnimplementedMhsBiodataServiceServer
	config        config.Config
	mhsbiodataSvc service.MhsBiodataServiceUseCase
}

func NewMhsBiodataHandler(config config.Config, mhsbiodataService service.MhsBiodataServiceUseCase) *MhsBiodataHandler {
	return &MhsBiodataHandler{
		config:        config,
		mhsbiodataSvc: mhsbiodataService,
	}
}

func (mbh *MhsBiodataHandler) FetchMhsBiodataByNim(ctx context.Context, req *pb.MhsBiodataRequest) (*pb.MhsBiodataResponse, error) {
	nim := req.GetNim()

	var apiResponse *entity.MhsBiodata
	apiResponse, err := mbh.mhsbiodataSvc.FetchMhsBiodataByNimFromSiakApi(nim)
	if err != nil {
		if apiResponse == nil {
			log.Println("WARNING: [MhsBiodataHandler - FetchMhsBiodataByNim] Resource not found: nim ", nim)
			// return nil, status.Errorf(codes.NotFound, "resource not found")
			return &pb.MhsBiodataResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "mhsbiodata not found",
			}, status.Errorf(codes.NotFound, "resource not found")
		}

		parseError := errors.ParseError(err)
		log.Println("ERROR: [MhsBiodataHandler - FetchMhsBiodataByNim] Error while fetching mhs biodata:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.MhsBiodataResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var mhsBiodata = entity.ConvertEntityToProto(apiResponse)

	return &pb.MhsBiodataResponse{
		Code:    uint32(http.StatusOK),
		Message: "get mhs biodata success",
		Data:    mhsBiodata,
	}, nil
}
