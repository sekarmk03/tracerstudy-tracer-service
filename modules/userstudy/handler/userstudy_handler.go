package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/modules/userstudy/entity"
	"tracerstudy-tracer-service/modules/userstudy/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserStudyHandler struct {
	pb.UnimplementedUserStudyServiceServer
	config       config.Config
	userStudySvc service.UserStudyServiceUseCase
}

func NewUserStudyHandler(config config.Config, userStudyService service.UserStudyServiceUseCase) *UserStudyHandler {
	return &UserStudyHandler{
		config:       config,
		userStudySvc: userStudyService,
	}
}

func (uh *UserStudyHandler) GetAllUserStudy(ctx context.Context, req *emptypb.Empty) (*pb.MultipleUserStudyResponse, error) {
	userStudy, err := uh.userStudySvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyHandler - GetAllUserStudy] Error while get all user study:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.MultipleUserStudyResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var userStudyArr []*pb.UserStudy
	for _, u := range userStudy {
		userStudyProto := entity.ConvertEntityToProto(u)
		userStudyArr = append(userStudyArr, userStudyProto)
	}

	return &pb.MultipleUserStudyResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all user study success",
		Data:    userStudyArr,
	}, nil
}

func (uh *UserStudyHandler) GetUserStudyByNim(ctx context.Context, req *pb.GetUserStudyByNimRequest) (*pb.SingleUserStudyResponse, error) {
	userStudy, err := uh.userStudySvc.FindByNim(ctx, req.GetNim(), req.GetEmailResponden(), req.GetHpResponden())
	if err != nil {
		if userStudy == nil {
			log.Println("WARNING: [UserStudyHandler - GetUserStudyByNim] Resource user study not found for nim:", req.GetNim())
			// return nil, status.Errorf(codes.NotFound, "user study not found")
			return &pb.SingleUserStudyResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "user study not found",
			}, status.Errorf(codes.NotFound, "user study not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyHandler - GetUserStudyByNim] Error while get user study by nim:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.SingleUserStudyResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	userStudyProto := entity.ConvertEntityToProto(userStudy)

	return &pb.SingleUserStudyResponse{
		Code:    uint32(http.StatusOK),
		Message: "get user study by nim success",
		Data:    userStudyProto,
	}, nil
}

func (uh *UserStudyHandler) UpdateUserStudy(ctx context.Context, req *pb.UserStudy) (*pb.SingleUserStudyResponse, error) {
	convertedProto := &entity.UserStudy{
		NamaInstansi:                      req.GetNamaInstansi(),
		Jabatan:                           req.GetJabatan(),
		AlamatInstansi:                    req.GetAlamatInstansi(),
		LamaMengenalLulusan:               req.GetLamaMengenalLulusan(),
		Etika:                             req.GetEtika(),
		KeahlianBidIlmu:                   req.GetKeahlianBidIlmu(),
		BahasaInggris:                     req.GetBahasaInggris(),
		PenggunaanTi:                      req.GetPenggunaanTi(),
		Komunikasi:                        req.GetKomunikasi(),
		KerjasamaTim:                      req.GetKerjasamaTim(),
		PengembanganDiri:                  req.GetPengembanganDiri(),
		KesiapanTerjunMasy:                req.GetKesiapanTerjunMasy(),
		KeunggulanLulusan:                 req.GetKeunggulanLulusan(),
		KelemahanLulusan:                  req.GetKelemahanLulusan(),
		SaranPeningkatanKompetensiLulusan: req.GetSaranPeningkatanKompetensiLulusan(),
		SaranPerbaikanKurikulum:           req.GetSaranPerbaikanKurikulum(),
	}

	userStudy, err := uh.userStudySvc.Update(ctx, req.GetNimLulusan(), req.GetEmailResponden(), req.GetHpResponden(), convertedProto)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyHandler - UpdateUserStudy] Error while update user study:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.SingleUserStudyResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	userStudyProto := entity.ConvertEntityToProto(userStudy)

	return &pb.SingleUserStudyResponse{
		Code:    uint32(http.StatusOK),
		Message: "update user study success",
		Data:    userStudyProto,
	}, nil
}

func (uh *UserStudyHandler) CreateUserStudy(ctx context.Context, req *pb.UserStudy) (*pb.SingleUserStudyResponse, error) {
	userStudy, err := uh.userStudySvc.Create(ctx, req.GetNamaResponden(), req.GetEmailResponden(), req.GetHpResponden(), req.GetNamaInstansi(), req.GetJabatan(), req.GetAlamatInstansi(), req.GetNimLulusan(), req.GetNamaLulusan(), req.GetProdiLulusan(), req.GetTahunLulusan())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyHandler - CreateUserStudy] Error while create user study:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.SingleUserStudyResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	userStudyProto := entity.ConvertEntityToProto(userStudy)

	return &pb.SingleUserStudyResponse{
		Code:    uint32(http.StatusOK),
		Message: "create user study success",
		Data:    userStudyProto,
	}, nil
}

func (uh *UserStudyHandler) ExportUSReport(ctx context.Context, req *emptypb.Empty) (*pb.ExportUSReportResponse, error) {
	us, err := uh.userStudySvc.ExportUSReport(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyHandler - ExportUSReport] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.ExportUSReportResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	usBytes := us.Bytes()

	return &pb.ExportUSReportResponse{
		Code:    uint32(http.StatusOK),
		Message: "export user study report success",
		Data:    usBytes,
	}, nil
}
