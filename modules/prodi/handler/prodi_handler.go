package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-tracer-service/common/config"

	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/modules/prodi/entity"
	"tracerstudy-tracer-service/modules/prodi/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProdiHandler struct {
	pb.UnimplementedProdiServiceServer
	config   config.Config
	prodiSvc service.ProdiServiceUseCase
}

func NewProdiHandler(config config.Config, prodiService service.ProdiServiceUseCase) *ProdiHandler {
	return &ProdiHandler{
		config:   config,
		prodiSvc: prodiService,
	}
}

func (ph *ProdiHandler) GetAllProdi(ctx context.Context, req *emptypb.Empty) (*pb.GetAllProdiResponse, error) {
	prodi, err := ph.prodiSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiHandler - GetAllProdi] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetAllProdiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var prodiArr []*pb.Prodi
	for _, p := range prodi {
		// convert each p to pb.Prodi
		prodiProto := entity.ConvertEntityToProto(p)
		prodiArr = append(prodiArr, prodiProto)
	}

	return &pb.GetAllProdiResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all prodi success",
		Data:    prodiArr,
	}, nil
}

func (ph *ProdiHandler) GetProdiByKode(ctx context.Context, req *pb.GetProdiByKodeRequest) (*pb.GetProdiResponse, error) {
	prodi, err := ph.prodiSvc.FindProdiByKode(ctx, req.GetKode())
	if err != nil {
		if prodi == nil {
			log.Println("WARNING: [ProdiHandler - GetProdiByKode] Resource prodi not found for kode:", req.GetKode())
			// return nil, status.Errorf(codes.NotFound, "prodi not found")
			return &pb.GetProdiResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "prodi not found",
			}, status.Errorf(codes.NotFound, "prodi not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiHandler - GetProdiByKode] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetProdiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	prodiProto := entity.ConvertEntityToProto(prodi)

	return &pb.GetProdiResponse{
		Code:    uint32(http.StatusOK),
		Message: "get prodi by kode success",
		Data:    prodiProto,
	}, nil
}

func (ph *ProdiHandler) CreateProdi(ctx context.Context, req *pb.Prodi) (*pb.GetProdiResponse, error) {
	prodi, err := ph.prodiSvc.Create(ctx, req.GetKode(), req.GetKodeDikti(), req.GetKodeFakultas(), req.GetKodeIntegrasi(), req.GetNama(), req.GetJenjang(), req.GetNamaFakultas(), req.GetAkronimFakultas())

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiHandler - CreateProdi] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetProdiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	prodiProto := entity.ConvertEntityToProto(prodi)

	return &pb.GetProdiResponse{
		Code:    uint32(http.StatusOK),
		Message: "create prodi success",
		Data:    prodiProto,
	}, nil
}

func (ph *ProdiHandler) UpdateProdi(ctx context.Context, req *pb.Prodi) (*pb.GetProdiResponse, error) {
	prodiDataUpdate := &entity.Prodi{
		KodeDikti:       req.GetKodeDikti(),
		KodeIntegrasi:   req.GetKodeIntegrasi(),
		Nama:            req.GetNama(),
		Jenjang:         req.GetJenjang(),
		KodeFakultas:    req.GetKodeFakultas(),
		NamaFakultas:    req.GetNamaFakultas(),
		AkronimFakultas: req.GetAkronimFakultas(),
	}

	prodi, err := ph.prodiSvc.Update(ctx, req.GetKode(), prodiDataUpdate)

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiHandler - UpdateProdi] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetProdiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	prodiProto := entity.ConvertEntityToProto(prodi)

	return &pb.GetProdiResponse{
		Code:    uint32(http.StatusOK),
		Message: "update prodi success",
		Data:    prodiProto,
	}, nil
}

func (ph *ProdiHandler) DeleteProdi(ctx context.Context, req *pb.GetProdiByKodeRequest) (*pb.DeleteProdiResponse, error) {
	err := ph.prodiSvc.Delete(ctx, req.GetKode())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiHandler - DeleteProdi] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.DeleteProdiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.DeleteProdiResponse{
		Code:    uint32(http.StatusOK),
		Message: "delete prodi success",
	}, nil
}
