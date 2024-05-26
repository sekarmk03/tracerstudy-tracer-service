package handler

import (
	"context"
	"log"
	"net/http"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	mhsbSvc "tracerstudy-tracer-service/modules/mhsbiodataapi/service"
	"tracerstudy-tracer-service/modules/responden/entity"
	resSvc "tracerstudy-tracer-service/modules/responden/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RespondenHandler struct {
	pb.UnimplementedRespondenServiceServer
	config           config.Config
	respondenSvc     resSvc.RespondenServiceUseCase
	mhsbiodataapiSvc mhsbSvc.MhsBiodataApiServiceUseCase
}

func NewRespondenHandler(config config.Config, respondenService resSvc.RespondenServiceUseCase, mhsbiodataapiService mhsbSvc.MhsBiodataApiServiceUseCase) *RespondenHandler {
	return &RespondenHandler{
		config:           config,
		respondenSvc:     respondenService,
		mhsbiodataapiSvc: mhsbiodataapiService,
	}
}

func (rh *RespondenHandler) GetAllResponden(ctx context.Context, req *emptypb.Empty) (*pb.GetAllRespondenResponse, error) {
	responden, err := rh.respondenSvc.FindAll(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenHandler - GetAllResponden] Error while get all responden:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetAllRespondenResponse{
			Code:    uint32(http.StatusOK),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var respondenArr []*pb.Responden
	for _, r := range responden {
		respondenProto := entity.ConvertEntityToProto(r)
		respondenArr = append(respondenArr, respondenProto)
	}

	return &pb.GetAllRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all responden success",
		Data:    respondenArr,
	}, nil
}

func (rh *RespondenHandler) GetRespondenByNim(ctx context.Context, req *pb.GetRespondenByNimRequest) (*pb.GetRespondenByNimResponse, error) {
	responden, err := rh.respondenSvc.FindByNim(ctx, req.GetNim())
	if err != nil {
		if responden == nil {
			log.Println("WARNING: [RespondenHandler - GetRespondenByNim] Resource responden not found for nim:", req.GetNim())
			// return nil, status.Errorf(codes.NotFound, "responden not found")
			return &pb.GetRespondenByNimResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "responden not found",
			}, status.Errorf(codes.NotFound, "responden not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenHandler - GetRespondenByNim] Error while get responden by nim:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetRespondenByNimResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.GetRespondenByNimResponse{
		Code:    uint32(http.StatusOK),
		Message: "get responden by nim success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) UpdateRespondenFromSiak(ctx context.Context, req *pb.UpdateRespondenFromSiakRequest) (*pb.UpdateRespondenResponse, error) {
	mhsbiodata, err := rh.mhsbiodataapiSvc.FetchMhsBiodataByNimFromSiakApi(req.GetNim())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: RespondenHandler-UpdateRespondenFromSiak] Error while fetch mhsbiodata from siak:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.UpdateRespondenResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var kodeProdi string
	if len(mhsbiodata.KODEPST) > 4 {
		kodeProdi = mhsbiodata.KODEPST[:4]
	} else {
		kodeProdi = mhsbiodata.KODEPST
	}

	var thnSidang string
	if len(mhsbiodata.TGLSIDANG) > 4 {
		thnSidang = mhsbiodata.TGLSIDANG[:4]
	} else {
		thnSidang = ""
	}

	convertedProto := &entity.Responden{
		Nama:          mhsbiodata.NAMA,
		StatusUpdate:  0,
		JalurMasuk:    mhsbiodata.JLRMASUK,
		TahunMasuk:    mhsbiodata.THNMASUK,
		LamaStudi:     uint32(utils.ConvStrToUint(mhsbiodata.LAMASTD, "lama_studi")),
		KodeFakultas:  mhsbiodata.KODEFAK,
		KodeProdi:     kodeProdi,
		JenisKelamin:  mhsbiodata.KODEJK,
		Email:         mhsbiodata.EMAIL,
		Hp:            utils.FormatPhoneNumber(mhsbiodata.HP),
		Ipk:           mhsbiodata.IPK,
		TanggalSidang: mhsbiodata.TGLSIDANG,
		TahunSidang:   thnSidang,
		// tanggal wisuda?
	}

	responden, err := rh.respondenSvc.Update(ctx, req.GetNim(), convertedProto)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenHandler - UpdateRespondenFromSiak] Error while update responden from siak:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.UpdateRespondenResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.UpdateRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "update responden from siak success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) CreateResponden(ctx context.Context, req *pb.CreateRespondenRequest) (*pb.CreateRespondenResponse, error) {
	responden, err := rh.respondenSvc.Create(ctx, req.GetNim())

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenHandler - CreateResponden] Error while create responden:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.CreateRespondenResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.CreateRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "create responden success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) UpdateResponden(ctx context.Context, req *pb.Responden) (*pb.UpdateRespondenResponse, error) {
	respDataUpdate := &entity.Responden{
		StatusUpdate:  1,
		Email:         req.GetEmail(),
		Hp:            utils.FormatPhoneNumber(req.GetHp()),
		Nik:           utils.CleanNumberText(req.GetNik()),
		Npwp:          utils.CleanNumberText(req.GetNpwp()),
		TanggalWisuda: req.GetTanggalWisuda(),
	}

	responden, err := rh.respondenSvc.Update(ctx, req.GetNim(), respDataUpdate)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenHandler - UpdateResponden] Error while update responden:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.UpdateRespondenResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	respondenProto := entity.ConvertEntityToProto(responden)

	return &pb.UpdateRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "update responden success",
		Data:    respondenProto,
	}, nil
}

func (rh *RespondenHandler) GetRespondenByNimList(ctx context.Context, req *pb.GetRespondenByNimListRequest) (*pb.GetAllRespondenResponse, error) {
	responden, err := rh.respondenSvc.FindByNimList(ctx, req.GetNims())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [RespondenHandler - GetRespondenByNimList] Error while get responden by nim list:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetAllRespondenResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var respondenArr []*pb.Responden
	for _, r := range responden {
		respondenProto := entity.ConvertEntityToProto(r)
		respondenArr = append(respondenArr, respondenProto)
	}

	return &pb.GetAllRespondenResponse{
		Code:    uint32(http.StatusOK),
		Message: "get responden by nim list success",
		Data:    respondenArr,
	}, nil
}
