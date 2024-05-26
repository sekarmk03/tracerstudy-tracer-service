package handler

import (
	"context"
	"log"
	"math"
	"net/http"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	"tracerstudy-tracer-service/modules/pkts/entity"
	"tracerstudy-tracer-service/modules/pkts/service"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PKTSHandler struct {
	pb.UnimplementedPKTSServiceServer
	config  config.Config
	PKTSSvc service.PKTSServiceUseCase
}

func NewPKTSHandler(config config.Config, pktsService service.PKTSServiceUseCase) *PKTSHandler {
	return &PKTSHandler{
		config:  config,
		PKTSSvc: pktsService,
	}
}

func (ph *PKTSHandler) GetAllPKTS(ctx context.Context, req *pb.GetAllPKTSRequest) (*pb.GetAllPKTSResponse, error) {
	pkts, totalRecords, err := ph.PKTSSvc.FindAll(ctx, req.Pagination.Limit, req.Pagination.Page)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - GetAllPKTS] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetAllPKTSResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var pktsArr []*pb.PKTS
	for _, p := range pkts {
		pktsProto := entity.ConvertEntityToProto(p)
		pktsArr = append(pktsArr, pktsProto)
	}

	totalPages := uint32(math.Ceil(float64(totalRecords) / float64(req.Pagination.Limit)))

	pagination := &pb.Pagination{
		TotalRows:   uint32(totalRecords),
		TotalPages:  totalPages,
		CurrentPage: req.Pagination.Page,
		CurrentRows: uint32(len(pkts)),
	}

	return &pb.GetAllPKTSResponse{
		Code:       uint32(http.StatusOK),
		Message:    "get all pkts success",
		Pagination: pagination,
		Data:       pktsArr,
	}, nil
}

func (ph *PKTSHandler) GetPKTSByNim(ctx context.Context, req *pb.GetPKTSByNimRequest) (*pb.GetPKTSResponse, error) {
	pkts, err := ph.PKTSSvc.FindByNim(ctx, req.Nim)
	if err != nil {
		if pkts == nil {
			log.Println("WARNING: [PKTSHandler - GetPKTSByNim] Resource pkts not found for nim:", req.Nim)
			// return nil, status.Errorf(codes.NotFound, "pkts not found")
			return &pb.GetPKTSResponse{
				Code:    uint32(http.StatusNotFound),
				Message: "pkts not found",
			}, status.Errorf(codes.NotFound, "pkts not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - GetPKTSByNim] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetPKTSResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	pktProto := entity.ConvertEntityToProto(pkts)

	return &pb.GetPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "get pkts by nim success",
		Data:    pktProto,
	}, nil
}

func (ph *PKTSHandler) CreatePKTS(ctx context.Context, req *pb.PKTS) (*pb.GetPKTSResponse, error) {
	// get data from api siak 1 -> kodeprodi, thn_sidang

	pkts, err := ph.PKTSSvc.Create(ctx, req.GetNim(), req.GetKodeProdi(), req.GetTahunSidang())

	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - CreatePKTS] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetPKTSResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	pktProto := entity.ConvertEntityToProto(pkts)

	return &pb.GetPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "create pkts success",
		Data:    pktProto,
	}, nil
}

func (ph *PKTSHandler) UpdatePKTS(ctx context.Context, req *pb.PKTS) (*pb.GetPKTSResponse, error) {
	pktsDataUpdate := &entity.PKTS{
		F8:                  req.GetF8(),
		F504:                req.GetF504(),
		F502:                req.GetF502(),
		F506:                req.GetF506(),
		F505:                req.GetF505(),
		F5a1:                req.GetF5A1(),
		F5a2:                req.GetF5A2(),
		F1101:               req.GetF1101(),
		F1102:               req.GetF1102(),
		F5b:                 req.GetF5B(),
		F5c:                 req.GetF5C(),
		F5d:                 req.GetF5D(),
		F18a:                req.GetF18A(),
		F18b:                req.GetF18B(),
		F18c:                req.GetF18C(),
		F18d:                req.GetF18D(),
		F1201:               req.GetF1201(),
		F1202:               req.GetF1202(),
		F14:                 req.GetF14(),
		F15:                 req.GetF15(),
		F1761:               req.GetF1761(),
		F1762:               req.GetF1762(),
		F1763:               req.GetF1763(),
		F1764:               req.GetF1764(),
		F1765:               req.GetF1765(),
		F1766:               req.GetF1766(),
		F1767:               req.GetF1767(),
		F1768:               req.GetF1768(),
		F1769:               req.GetF1769(),
		F1770:               req.GetF1770(),
		F1771:               req.GetF1771(),
		F1772:               req.GetF1772(),
		F1773:               req.GetF1773(),
		F1774:               req.GetF1774(),
		F21:                 req.GetF21(),
		F22:                 req.GetF22(),
		F23:                 req.GetF23(),
		F24:                 req.GetF24(),
		F25:                 req.GetF25(),
		F26:                 req.GetF26(),
		F27:                 req.GetF27(),
		F301:                req.GetF301(),
		F302:                req.GetF302(),
		F303:                req.GetF303(),
		F401:                req.GetF401(),
		F402:                req.GetF402(),
		F403:                req.GetF403(),
		F404:                req.GetF404(),
		F405:                req.GetF405(),
		F406:                req.GetF406(),
		F407:                req.GetF407(),
		F408:                req.GetF408(),
		F409:                req.GetF409(),
		F410:                req.GetF410(),
		F411:                req.GetF411(),
		F412:                req.GetF412(),
		F413:                req.GetF413(),
		F414:                req.GetF414(),
		F415:                req.GetF415(),
		F416:                req.GetF416(),
		F6:                  req.GetF6(),
		F7:                  req.GetF7(),
		F7a:                 req.GetF7A(),
		F1001:               req.GetF1001(),
		F1002:               req.GetF1002(),
		F1601:               req.GetF1601(),
		F1602:               req.GetF1602(),
		F1603:               req.GetF1603(),
		F1604:               req.GetF1604(),
		F1605:               req.GetF1605(),
		F1606:               req.GetF1606(),
		F1607:               req.GetF1607(),
		F1608:               req.GetF1608(),
		F1609:               req.GetF1609(),
		F1610:               req.GetF1610(),
		F1611:               req.GetF1611(),
		F1612:               req.GetF1612(),
		F1613:               req.GetF1613(),
		F1614:               req.GetF1614(),
		NamaAtasan:          req.GetNamaAtasan(),
		HpAtasan:            utils.FormatPhoneNumber(req.GetHpAtasan()),
		EmailAtasan:         req.GetEmailAtasan(),
		TinggalSelamaKuliah: req.GetTinggalSelamaKuliah(),
	}

	pkt, err := ph.PKTSSvc.Update(ctx, req.GetNim(), pktsDataUpdate)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - UpdatePKTS] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetPKTSResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	pktsProto := entity.ConvertEntityToProto(pkt)

	return &pb.GetPKTSResponse{
		Code:    uint32(http.StatusOK),
		Message: "update pkts success",
		Data:    pktsProto,
	}, nil
}

func (ph *PKTSHandler) GetNimByDataAtasan(ctx context.Context, req *pb.GetNimByDataAtasanRequest) (*pb.GetNimByDataAtasanResponse, error) {
	nims, err := ph.PKTSSvc.FindByAtasan(
		ctx,
		req.GetNamaAtasan(),
		utils.FormatPhoneNumber(req.GetHpAtasan()),
		req.GetEmailAtasan(),
	)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - GetNimByDataAtasan] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetNimByDataAtasanResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	result := make([]string, len(nims))
	for i, v := range nims {
		result[i] = *v
	}

	return &pb.GetNimByDataAtasanResponse{
		Code:    uint32(http.StatusOK),
		Message: "get pkts by atasan success",
		Nims:    result,
	}, nil
}

func (ph *PKTSHandler) ExportPKTSReport(ctx context.Context, req *pb.ExportPKTSReportRequest) (*pb.ExportPKTSReportResponse, error) {
	pkts, err := ph.PKTSSvc.ExportPKTSReport(ctx, req.GetTahunSidang())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - ExportPKTSReport] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.ExportPKTSReportResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	pktsBytes := pkts.Bytes()

	return &pb.ExportPKTSReportResponse{
		Code:    uint32(http.StatusOK),
		Message: "export pkts report success",
		Data:    pktsBytes,
	}, nil
}

func (ph *PKTSHandler) GetPKTSRekapByProdi(ctx context.Context, req *pb.GetPKTSRekapByProdiRequest) (*pb.GetPKTSRekapByProdiResponse, error) {
	pktsRekap, totalRecords, err := ph.PKTSSvc.FindPKTSRekapByProdi(ctx, req.Pagination.Limit, req.Pagination.Page, req.GetKodeprodi(), req.GetTahunSidang())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - GetPKTSRekapByProdi] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetPKTSRekapByProdiResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var pktsRekapArr []*pb.PKTSRekapByProdi
	for _, p := range pktsRekap {
		pktsRekapProto := entity.ConvertEntityPKTSRekapByProdiToProto(p)
		pktsRekapArr = append(pktsRekapArr, pktsRekapProto)
	}

	totalPages := uint32(math.Ceil(float64(totalRecords) / float64(req.Pagination.Limit)))

	pagination := &pb.Pagination{
		TotalRows:   uint32(totalRecords),
		TotalPages:  totalPages,
		CurrentPage: req.Pagination.Page,
		CurrentRows: uint32(len(pktsRekap)),
	}

	return &pb.GetPKTSRekapByProdiResponse{
		Code:       uint32(http.StatusOK),
		Message:    "get pkts rekap by prodi success",
		Pagination: pagination,
		Data:       pktsRekapArr,
	}, nil
}

func (ph *PKTSHandler) GetPKTSRekapByYear(ctx context.Context, req *pb.GetPKTSRekapByYearRequest) (*pb.GetPKTSRekapByYearResponse, error) {
	pktsRekap, totalRecords, err := ph.PKTSSvc.FindPKTSRekapByYear(ctx, req.Pagination.Limit, req.Pagination.Page, req.GetTahunSidang())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSHandler - GetPKTSRekapByYear] Internal server error:", parseError.Message)
		// return nil, status.Errorf(parseError.Code, parseError.Message)
		return &pb.GetPKTSRekapByYearResponse{
			Code:    uint32(http.StatusInternalServerError),
			Message: parseError.Message,
		}, status.Errorf(parseError.Code, parseError.Message)
	}

	var pktsRekapArr []*pb.PKTSRekapByYear
	for _, p := range pktsRekap {
		pktsRekapProto := entity.ConvertEntityPKTSRekapByYearToProto(p)
		pktsRekapArr = append(pktsRekapArr, pktsRekapProto)
	}

	totalPages := uint32(math.Ceil(float64(totalRecords) / float64(req.Pagination.Limit)))

	pagination := &pb.Pagination{
		TotalRows:   uint32(totalRecords),
		TotalPages:  totalPages,
		CurrentPage: req.Pagination.Page,
		CurrentRows: uint32(len(pktsRekap)),
	}

	return &pb.GetPKTSRekapByYearResponse{
		Code:       uint32(http.StatusOK),
		Message:    "get pkts rekap by year success",
		Pagination: pagination,
		Data:       pktsRekapArr,
	}, nil
}
