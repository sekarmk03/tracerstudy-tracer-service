package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/modules/fakultas/entity"
	"tracerstudy-tracer-service/modules/fakultas/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FakultasHandler struct {
	pb.UnimplementedFakultasServiceServer
	config      config.Config
	fakultasSvc service.FakultasServiceUseCase
}

func NewFakultasHandler(config config.Config, fakultasService service.FakultasServiceUseCase) *FakultasHandler {
	return &FakultasHandler{
		config:      config,
		fakultasSvc: fakultasService,
	}
}

func (fh *FakultasHandler) GetAllFakultas(ctx context.Context, req *emptypb.Empty) (*pb.GetAllFakultasResponse, error) {
	fakultas, err := fh.fakultasSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasHandler - GetAllFakultas] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetAllFakultasResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var fakultasArr []*pb.Fakultas
	for _, f := range fakultas {
		// convert each f to pb.Fakultas
		fakultasProto := entity.ConvertEntityToProto(f)
		fakultasArr = append(fakultasArr, fakultasProto)
	}

	return &pb.GetAllFakultasResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all fakultas success",
		Data:    fakultasArr,
	}, nil
}

func (fh *FakultasHandler) GetFakultasByKode(ctx context.Context, req *pb.GetFakultasByKodeRequest) (*pb.GetFakultasResponse, error) {
	fakultas, err := fh.fakultasSvc.FindFakultasByKode(ctx, req.GetKode())
	if err != nil {
		if fakultas == nil {
			log.Println("WARNING: [FakultasHandler - GetFakultasByKode] Resource fakultas not found for kode:", req.GetKode())
			// return nil, status.Errorf(codes.NotFound, "fakultas not found")
			return &pb.GetFakultasResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "fakultas not found",
			}, status.Errorf(codes.NotFound, "fakultas not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasHandler - GetFakultasByKode] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetFakultasResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	fakultasProto := entity.ConvertEntityToProto(fakultas)

	return &pb.GetFakultasResponse{
		Code:    uint32(http.StatusOK),
		Message: "get fakultas by kode success",
		Data:    fakultasProto,
	}, nil
}

func (fh *FakultasHandler) CreateFakultas(ctx context.Context, req *pb.Fakultas) (*pb.GetFakultasResponse, error) {
	fakultas, err := fh.fakultasSvc.Create(ctx, req.GetKode(), req.GetNama(), req.GetAkronim())

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasHandler - CreateFakultas] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetFakultasResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	fakultasProto := entity.ConvertEntityToProto(fakultas)

	return &pb.GetFakultasResponse{
		Code:    uint32(http.StatusOK),
		Message: "create fakultas success",
		Data:    fakultasProto,
	}, nil
}

func (fh *FakultasHandler) UpdateFakultas(ctx context.Context, req *pb.Fakultas) (*pb.GetFakultasResponse, error) {
	fakultas, err := fh.fakultasSvc.Update(ctx, req.GetKode(), &entity.Fakultas{
		Nama:    req.GetNama(),
		Akronim: req.GetAkronim(),
	})

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasHandler - UpdateFakultas] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetFakultasResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	fakultasProto := entity.ConvertEntityToProto(fakultas)

	return &pb.GetFakultasResponse{
		Code:    uint32(http.StatusOK),
		Message: "update fakultas success",
		Data:    fakultasProto,
	}, nil
}

func (fh *FakultasHandler) DeleteFakultas(ctx context.Context, req *pb.GetFakultasByKodeRequest) (*pb.DeleteFakultasResponse, error) {
	err := fh.fakultasSvc.Delete(ctx, req.GetKode())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasHandler - DeleteFakultas] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.DeleteFakultasResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.DeleteFakultasResponse{
		Code:    uint32(http.StatusOK),
		Message: "delete fakultas success",
	}, nil
}
