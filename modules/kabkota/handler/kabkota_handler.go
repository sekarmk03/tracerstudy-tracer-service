package handler

import (
	"context"
	"log"
	"math"
	"net/http"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/modules/kabkota/entity"
	"tracerstudy-tracer-service/modules/kabkota/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KabKotaHandler struct {
	pb.UnimplementedKabKotaServiceServer
	config     config.Config
	kabkotaSvc service.KabKotaServiceUseCase
}

func NewKabKotaHandler(config config.Config, kabkotaService service.KabKotaServiceUseCase) *KabKotaHandler {
	return &KabKotaHandler{
		config:     config,
		kabkotaSvc: kabkotaService,
	}
}

func (kh *KabKotaHandler) GetAllKabKota(ctx context.Context, req *pb.GetAllKabKotaRequest) (*pb.GetAllKabKotaResponse, error) {
	if req == nil || req.Pagination == nil {
        return &pb.GetAllKabKotaResponse{
			Code:    uint32(http.StatusBadRequest),
			Message: "request cannot be nil",
		}, status.Errorf(codes.InvalidArgument, "request cannot be nil")
    }

	kabkota, totalRows, err := kh.kabkotaSvc.FindAll(ctx, req.Pagination.Page, req.Pagination.Limit)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler - GetAllKabKota] Error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetAllKabKotaResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var kabkotaArr []*pb.KabKota
	for _, k := range kabkota {
		kabkotaProto := entity.ConvertEntityToProto(k)
		kabkotaArr = append(kabkotaArr, kabkotaProto)
	}

	totalPages := uint32(math.Ceil(float64(totalRows) / float64(req.Pagination.Limit)))

	pagination := &pb.Pagination{
		TotalRows:   uint32(totalRows),
		TotalPages:  totalPages,
		CurrentPage: req.Pagination.Page,
		CurrentRows: uint32(len(kabkotaArr)),
	}

	return &pb.GetAllKabKotaResponse{
		Code:       uint32(http.StatusOK),
		Message:    "get all kabkota success",
		Pagination: pagination,
		Data:       kabkotaArr,
	}, nil
}

func (kh *KabKotaHandler) GetKabKotaByIdWil(ctx context.Context, req *pb.GetKabKotaByIdWilRequest) (*pb.GetKabKotaResponse, error) {
	kabkota, err := kh.kabkotaSvc.FindByIdWil(ctx, req.GetIdWil())
	if err != nil {
		if kabkota == nil {
			log.Println("WARNING: [KabKotaHandler - GetKabKotaByIdWil] Resource kabkota not found for idWil:", req.GetIdWil())
			// return nil, status.Errorf(codes.NotFound, "kabkota not found")
			return &pb.GetKabKotaResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "kabkota not found",
			}, status.Errorf(codes.NotFound, "kabkota not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler - GetKabKotaByIdWil] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetKabKotaResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	kabkotaProto := entity.ConvertEntityToProto(kabkota)

	return &pb.GetKabKotaResponse{
		Code:    uint32(http.StatusOK),
		Message: "get kabkota by idWil success",
		Data:    kabkotaProto,
	}, nil
}

func (kh *KabKotaHandler) CreateKabKota(ctx context.Context, req *pb.KabKota) (*pb.GetKabKotaResponse, error) {
	kabkota, err := kh.kabkotaSvc.Create(ctx, req.GetIdWil(), req.GetNama(), req.GetIdIndukWil())

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler - CreateKabKota] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetKabKotaResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	kabkotaProto := entity.ConvertEntityToProto(kabkota)

	return &pb.GetKabKotaResponse{
		Code:    uint32(http.StatusOK),
		Message: "create kabkota success",
		Data:    kabkotaProto,
	}, nil
}

func (kh *KabKotaHandler) UpdateKabKota(ctx context.Context, req *pb.KabKota) (*pb.GetKabKotaResponse, error) {
	kabkotaDataUpdate := &entity.KabKota{
		Nama:       req.GetNama(),
		IdIndukWil: req.GetIdIndukWil(),
	}

	kabkota, err := kh.kabkotaSvc.Update(ctx, req.GetIdWil(), kabkotaDataUpdate)

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler - UpdateKabKota] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetKabKotaResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	kabkotaProto := entity.ConvertEntityToProto(kabkota)

	return &pb.GetKabKotaResponse{
		Code:    uint32(http.StatusOK),
		Message: "update kabkota success",
		Data:    kabkotaProto,
	}, nil
}

func (kh *KabKotaHandler) DeleteKabKota(ctx context.Context, req *pb.GetKabKotaByIdWilRequest) (*pb.DeleteKabKotaResponse, error) {
	err := kh.kabkotaSvc.Delete(ctx, req.GetIdWil())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaHandler - DeleteKabKota] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.DeleteKabKotaResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.DeleteKabKotaResponse{
		Code:    uint32(http.StatusOK),
		Message: "delete kabkota success",
	}, nil
}
