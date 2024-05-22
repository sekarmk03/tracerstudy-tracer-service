package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	"tracerstudy-tracer-service/modules/pkts/entity"
	"tracerstudy-tracer-service/modules/pkts/repository"

	"github.com/xuri/excelize/v2"
)

type PKTSService struct {
	cfg            config.Config
	pktsRepository repository.PKTSRepositoryUseCase
}

type PKTSServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.PKTS, error)
	FindByNim(ctx context.Context, nim string) (*entity.PKTS, error)
	Create(ctx context.Context, nim, kodeProdi, tahunSidang string) (*entity.PKTS, error)
	Update(ctx context.Context, nim string, fields *entity.PKTS) (*entity.PKTS, error)
	FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error)
	ExportPKTSReport(ctx context.Context, tahunSidang string) (*bytes.Buffer, error)
	FindPKTSRekapByProdi(ctx context.Context, kodeprodi, tahunSidang string) ([]*entity.PKTSRekapByProdi, error)
}

func NewPKTSService(cfg config.Config, pktsRepository repository.PKTSRepositoryUseCase) *PKTSService {
	return &PKTSService{
		cfg:            cfg,
		pktsRepository: pktsRepository,
	}
}

func (svc *PKTSService) FindAll(ctx context.Context, req any) ([]*entity.PKTS, error) {
	res, err := svc.pktsRepository.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService - FindAll] Error while find all pkts:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) FindByNim(ctx context.Context, nim string) (*entity.PKTS, error) {
	res, err := svc.pktsRepository.FindByNim(ctx, nim)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService - FindByNim] Error while find pkts by nim:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) Create(ctx context.Context, nim, kodeProdi, tahunSidang string) (*entity.PKTS, error) {
	reqEntity := &entity.PKTS{
		Nim:         nim,
		KodeProdi:   kodeProdi,
		TahunSidang: tahunSidang,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	res, err := svc.pktsRepository.Create(ctx, reqEntity)
	if err != nil {
		log.Println("ERROR: [PKTSService - Create] Error while create pkts:", err)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) Update(ctx context.Context, nim string, fields *entity.PKTS) (*entity.PKTS, error) {
	pkts, err := svc.pktsRepository.FindByNim(ctx, nim)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService - Update] Error while find pkts by nim:", parseError.Message)
		return nil, err
	}

	updateMap := make(map[string]interface{})

	utils.AddItemToMap(updateMap, "f8", fields.F8)
	utils.AddItemToMap(updateMap, "f504", fields.F504)
	utils.AddItemToMap(updateMap, "f502", fields.F502)
	utils.AddItemToMap(updateMap, "f506", fields.F506)
	utils.AddItemToMap(updateMap, "f505", fields.F505)
	utils.AddItemToMap(updateMap, "f5a1", fields.F5a1)
	utils.AddItemToMap(updateMap, "f5a2", fields.F5a2)
	utils.AddItemToMap(updateMap, "f1101", fields.F1101)
	utils.AddItemToMap(updateMap, "f1102", fields.F1102)
	utils.AddItemToMap(updateMap, "f5b", fields.F5b)
	utils.AddItemToMap(updateMap, "f5c", fields.F5c)
	utils.AddItemToMap(updateMap, "f5d", fields.F5d)
	utils.AddItemToMap(updateMap, "f18a", fields.F18a)
	utils.AddItemToMap(updateMap, "f18b", fields.F18b)
	utils.AddItemToMap(updateMap, "f18c", fields.F18c)
	utils.AddItemToMap(updateMap, "f18d", fields.F18d)
	utils.AddItemToMap(updateMap, "f1201", fields.F1201)
	utils.AddItemToMap(updateMap, "f1202", fields.F1202)
	utils.AddItemToMap(updateMap, "f14", fields.F14)
	utils.AddItemToMap(updateMap, "f15", fields.F15)
	utils.AddItemToMap(updateMap, "f1761", fields.F1761)
	utils.AddItemToMap(updateMap, "f1762", fields.F1762)
	utils.AddItemToMap(updateMap, "f1763", fields.F1763)
	utils.AddItemToMap(updateMap, "f1764", fields.F1764)
	utils.AddItemToMap(updateMap, "f1765", fields.F1765)
	utils.AddItemToMap(updateMap, "f1766", fields.F1766)
	utils.AddItemToMap(updateMap, "f1767", fields.F1767)
	utils.AddItemToMap(updateMap, "f1768", fields.F1768)
	utils.AddItemToMap(updateMap, "f1769", fields.F1769)
	utils.AddItemToMap(updateMap, "f1770", fields.F1770)
	utils.AddItemToMap(updateMap, "f1771", fields.F1771)
	utils.AddItemToMap(updateMap, "f1772", fields.F1772)
	utils.AddItemToMap(updateMap, "f1773", fields.F1773)
	utils.AddItemToMap(updateMap, "f1774", fields.F1774)
	utils.AddItemToMap(updateMap, "f21", fields.F21)
	utils.AddItemToMap(updateMap, "f22", fields.F22)
	utils.AddItemToMap(updateMap, "f23", fields.F23)
	utils.AddItemToMap(updateMap, "f24", fields.F24)
	utils.AddItemToMap(updateMap, "f25", fields.F25)
	utils.AddItemToMap(updateMap, "f26", fields.F26)
	utils.AddItemToMap(updateMap, "f27", fields.F27)
	utils.AddItemToMap(updateMap, "f301", fields.F301)
	utils.AddItemToMap(updateMap, "f302", fields.F302)
	utils.AddItemToMap(updateMap, "f303", fields.F303)
	utils.AddItemToMap(updateMap, "f401", fields.F401)
	utils.AddItemToMap(updateMap, "f402", fields.F402)
	utils.AddItemToMap(updateMap, "f403", fields.F403)
	utils.AddItemToMap(updateMap, "f404", fields.F404)
	utils.AddItemToMap(updateMap, "f405", fields.F405)
	utils.AddItemToMap(updateMap, "f406", fields.F406)
	utils.AddItemToMap(updateMap, "f407", fields.F407)
	utils.AddItemToMap(updateMap, "f408", fields.F408)
	utils.AddItemToMap(updateMap, "f409", fields.F409)
	utils.AddItemToMap(updateMap, "f410", fields.F410)
	utils.AddItemToMap(updateMap, "f411", fields.F411)
	utils.AddItemToMap(updateMap, "f412", fields.F412)
	utils.AddItemToMap(updateMap, "f413", fields.F413)
	utils.AddItemToMap(updateMap, "f414", fields.F414)
	utils.AddItemToMap(updateMap, "f415", fields.F415)
	utils.AddItemToMap(updateMap, "f416", fields.F416)
	utils.AddItemToMap(updateMap, "f6", fields.F6)
	utils.AddItemToMap(updateMap, "f7", fields.F7)
	utils.AddItemToMap(updateMap, "f7a", fields.F7a)
	utils.AddItemToMap(updateMap, "f1001", fields.F1001)
	utils.AddItemToMap(updateMap, "f1002", fields.F1002)
	utils.AddItemToMap(updateMap, "f1601", fields.F1601)
	utils.AddItemToMap(updateMap, "f1602", fields.F1602)
	utils.AddItemToMap(updateMap, "f1603", fields.F1603)
	utils.AddItemToMap(updateMap, "f1604", fields.F1604)
	utils.AddItemToMap(updateMap, "f1605", fields.F1605)
	utils.AddItemToMap(updateMap, "f1606", fields.F1606)
	utils.AddItemToMap(updateMap, "f1607", fields.F1607)
	utils.AddItemToMap(updateMap, "f1608", fields.F1608)
	utils.AddItemToMap(updateMap, "f1609", fields.F1609)
	utils.AddItemToMap(updateMap, "f1610", fields.F1610)
	utils.AddItemToMap(updateMap, "f1611", fields.F1611)
	utils.AddItemToMap(updateMap, "f1612", fields.F1612)
	utils.AddItemToMap(updateMap, "f1613", fields.F1613)
	utils.AddItemToMap(updateMap, "f1614", fields.F1614)
	utils.AddItemToMap(updateMap, "nama_atasan", fields.NamaAtasan)
	utils.AddItemToMap(updateMap, "hp_atasan", fields.HpAtasan)
	utils.AddItemToMap(updateMap, "email_atasan", fields.EmailAtasan)
	utils.AddItemToMap(updateMap, "tinggal_selama_kuliah", fields.TinggalSelamaKuliah)

	res, err := svc.pktsRepository.Update(ctx, pkts, updateMap)
	if err != nil {
		log.Println("ERROR: [PKTSService - Update] Error while update pkts:", err)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error) {
	res, err := svc.pktsRepository.FindByAtasan(ctx, namaA, hpA, emailA)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService - FindByAtasan] Error while find pkts by atasan:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *PKTSService) ExportPKTSReport(ctx context.Context, tahunSidang string) (*bytes.Buffer, error) {
	rows, err := svc.pktsRepository.FindAllReport(ctx, tahunSidang)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService - FindAll] Error while find all pkts:", parseError.Message)
		return nil, err
	}

	file := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := file.NewSheet(sheetName)
	if err != nil {
		log.Println("ERROR: [PKTSService - ExportPKTSReport] Error while create new sheet:", err)
		return nil, err
	}

	file.SetCellValue(sheetName, "A1", "created_at")
	file.SetCellValue(sheetName, "B1", "updated_at")
	file.SetCellValue(sheetName, "C1", "kodeprodi")
	file.SetCellValue(sheetName, "D1", "nama_prodi")
	file.SetCellValue(sheetName, "E1", "kdptimsmh")
	file.SetCellValue(sheetName, "F1", "kdpstmsmh")
	file.SetCellValue(sheetName, "G1", "jenjang")
	file.SetCellValue(sheetName, "H1", "nimhsmsmh")
	file.SetCellValue(sheetName, "I1", "nmmhsmsmh")
	file.SetCellValue(sheetName, "J1", "jenis_kelamin")
	file.SetCellValue(sheetName, "K1", "telpomsmh")
	file.SetCellValue(sheetName, "L1", "emailmsmh")
	file.SetCellValue(sheetName, "M1", "tahun_lulus")
	file.SetCellValue(sheetName, "N1", "nik")
	file.SetCellValue(sheetName, "O1", "npwp")
	file.SetCellValue(sheetName, "P1", "f8")
	file.SetCellValue(sheetName, "Q1", "f504")
	file.SetCellValue(sheetName, "R1", "f502")
	file.SetCellValue(sheetName, "S1", "f505")
	file.SetCellValue(sheetName, "T1", "f506")
	file.SetCellValue(sheetName, "U1", "f5a1")
	file.SetCellValue(sheetName, "V1", "f5a2")
	file.SetCellValue(sheetName, "W1", "f1101")
	file.SetCellValue(sheetName, "X1", "f1102")
	file.SetCellValue(sheetName, "Y1", "f5b")
	file.SetCellValue(sheetName, "Z1", "f5c")
	file.SetCellValue(sheetName, "AA1", "f5d")
	file.SetCellValue(sheetName, "AB1", "f18a")
	file.SetCellValue(sheetName, "AC1", "f18b")
	file.SetCellValue(sheetName, "AD1", "f18c")
	file.SetCellValue(sheetName, "AE1", "f18d")
	file.SetCellValue(sheetName, "AF1", "f1201")
	file.SetCellValue(sheetName, "AG1", "f1202")
	file.SetCellValue(sheetName, "AH1", "f14")
	file.SetCellValue(sheetName, "AI1", "f15")
	file.SetCellValue(sheetName, "AJ1", "f1761")
	file.SetCellValue(sheetName, "AK1", "f1762")
	file.SetCellValue(sheetName, "AL1", "f1763")
	file.SetCellValue(sheetName, "AM1", "f1764")
	file.SetCellValue(sheetName, "AN1", "f1765")
	file.SetCellValue(sheetName, "AO1", "f1766")
	file.SetCellValue(sheetName, "AP1", "f1767")
	file.SetCellValue(sheetName, "AQ1", "f1768")
	file.SetCellValue(sheetName, "AR1", "f1769")
	file.SetCellValue(sheetName, "AS1", "f1770")
	file.SetCellValue(sheetName, "AT1", "f1771")
	file.SetCellValue(sheetName, "AU1", "f1772")
	file.SetCellValue(sheetName, "AV1", "f1773")
	file.SetCellValue(sheetName, "AW1", "f1774")
	file.SetCellValue(sheetName, "AX1", "f21")
	file.SetCellValue(sheetName, "AY1", "f22")
	file.SetCellValue(sheetName, "AZ1", "f23")
	file.SetCellValue(sheetName, "BA1", "f24")
	file.SetCellValue(sheetName, "BB1", "f25")
	file.SetCellValue(sheetName, "BC1", "f26")
	file.SetCellValue(sheetName, "BD1", "f27")
	file.SetCellValue(sheetName, "BE1", "f301")
	file.SetCellValue(sheetName, "BF1", "f302")
	file.SetCellValue(sheetName, "BG1", "f303")
	file.SetCellValue(sheetName, "BH1", "f401")
	file.SetCellValue(sheetName, "BI1", "f402")
	file.SetCellValue(sheetName, "BJ1", "f403")
	file.SetCellValue(sheetName, "BK1", "f404")
	file.SetCellValue(sheetName, "BL1", "f405")
	file.SetCellValue(sheetName, "BM1", "f406")
	file.SetCellValue(sheetName, "BN1", "f407")
	file.SetCellValue(sheetName, "BO1", "f408")
	file.SetCellValue(sheetName, "BP1", "f409")
	file.SetCellValue(sheetName, "BQ1", "f410")
	file.SetCellValue(sheetName, "BR1", "f411")
	file.SetCellValue(sheetName, "BS1", "f412")
	file.SetCellValue(sheetName, "BT1", "f413")
	file.SetCellValue(sheetName, "BU1", "f414")
	file.SetCellValue(sheetName, "BV1", "f415")
	file.SetCellValue(sheetName, "BW1", "f416")
	file.SetCellValue(sheetName, "BX1", "f6")
	file.SetCellValue(sheetName, "BY1", "f7")
	file.SetCellValue(sheetName, "BZ1", "f7a")
	file.SetCellValue(sheetName, "CA1", "f1001")
	file.SetCellValue(sheetName, "CB1", "f1002")
	file.SetCellValue(sheetName, "CC1", "f1601")
	file.SetCellValue(sheetName, "CD1", "f1602")
	file.SetCellValue(sheetName, "CE1", "f1603")
	file.SetCellValue(sheetName, "CF1", "f1604")
	file.SetCellValue(sheetName, "CG1", "f1605")
	file.SetCellValue(sheetName, "CH1", "f1606")
	file.SetCellValue(sheetName, "CI1", "f1607")
	file.SetCellValue(sheetName, "CJ1", "f1608")
	file.SetCellValue(sheetName, "CK1", "f1609")
	file.SetCellValue(sheetName, "CL1", "f1610")
	file.SetCellValue(sheetName, "CM1", "f1611")
	file.SetCellValue(sheetName, "CN1", "f1612")
	file.SetCellValue(sheetName, "CO1", "f1613")
	file.SetCellValue(sheetName, "CP1", "f1614")
	file.SetCellValue(sheetName, "CQ1", "nama_atasan")
	file.SetCellValue(sheetName, "CR1", "hp_atasan")
	file.SetCellValue(sheetName, "CS1", "email_atasan")

	kodePT := "001034"

	for i, row := range rows {
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), row.CreatedAt)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), row.UpdatedAt)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), row.KodeProdi)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), row.NamaProdi)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), kodePT)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), row.KodeDikti)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), row.Jenjang)
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), row.Nim)
		file.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), row.Nama)
		file.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), row.JK)
		file.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), row.Hp)
		file.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), row.Email)
		file.SetCellValue(sheetName, fmt.Sprintf("M%d", i+2), row.TahunSidang)
		file.SetCellValue(sheetName, fmt.Sprintf("N%d", i+2), row.Nik)
		file.SetCellValue(sheetName, fmt.Sprintf("O%d", i+2), row.Npwp)
		file.SetCellValue(sheetName, fmt.Sprintf("P%d", i+2), row.F8)
		file.SetCellValue(sheetName, fmt.Sprintf("Q%d", i+2), row.F504)
		file.SetCellValue(sheetName, fmt.Sprintf("R%d", i+2), row.F502)
		file.SetCellValue(sheetName, fmt.Sprintf("S%d", i+2), row.F505)
		file.SetCellValue(sheetName, fmt.Sprintf("T%d", i+2), row.F506)
		file.SetCellValue(sheetName, fmt.Sprintf("U%d", i+2), row.F5a1)
		file.SetCellValue(sheetName, fmt.Sprintf("V%d", i+2), row.F5a2)
		file.SetCellValue(sheetName, fmt.Sprintf("W%d", i+2), row.F1101)
		file.SetCellValue(sheetName, fmt.Sprintf("X%d", i+2), row.F1102)
		file.SetCellValue(sheetName, fmt.Sprintf("Y%d", i+2), row.F5b)
		file.SetCellValue(sheetName, fmt.Sprintf("Z%d", i+2), row.F5c)
		file.SetCellValue(sheetName, fmt.Sprintf("AA%d", i+2), row.F5d)
		file.SetCellValue(sheetName, fmt.Sprintf("AB%d", i+2), row.F18a)
		file.SetCellValue(sheetName, fmt.Sprintf("AC%d", i+2), row.F18b)
		file.SetCellValue(sheetName, fmt.Sprintf("AD%d", i+2), row.F18c)
		file.SetCellValue(sheetName, fmt.Sprintf("AE%d", i+2), row.F18d)
		file.SetCellValue(sheetName, fmt.Sprintf("AF%d", i+2), row.F1201)
		file.SetCellValue(sheetName, fmt.Sprintf("AG%d", i+2), row.F1202)
		file.SetCellValue(sheetName, fmt.Sprintf("AH%d", i+2), row.F14)
		file.SetCellValue(sheetName, fmt.Sprintf("AI%d", i+2), row.F15)
		file.SetCellValue(sheetName, fmt.Sprintf("AJ%d", i+2), row.F1761)
		file.SetCellValue(sheetName, fmt.Sprintf("AK%d", i+2), row.F1762)
		file.SetCellValue(sheetName, fmt.Sprintf("AL%d", i+2), row.F1763)
		file.SetCellValue(sheetName, fmt.Sprintf("AM%d", i+2), row.F1764)
		file.SetCellValue(sheetName, fmt.Sprintf("AN%d", i+2), row.F1765)
		file.SetCellValue(sheetName, fmt.Sprintf("AO%d", i+2), row.F1766)
		file.SetCellValue(sheetName, fmt.Sprintf("AP%d", i+2), row.F1767)
		file.SetCellValue(sheetName, fmt.Sprintf("AQ%d", i+2), row.F1768)
		file.SetCellValue(sheetName, fmt.Sprintf("AR%d", i+2), row.F1769)
		file.SetCellValue(sheetName, fmt.Sprintf("AS%d", i+2), row.F1770)
		file.SetCellValue(sheetName, fmt.Sprintf("AT%d", i+2), row.F1771)
		file.SetCellValue(sheetName, fmt.Sprintf("AU%d", i+2), row.F1772)
		file.SetCellValue(sheetName, fmt.Sprintf("AV%d", i+2), row.F1773)
		file.SetCellValue(sheetName, fmt.Sprintf("AW%d", i+2), row.F1774)
		file.SetCellValue(sheetName, fmt.Sprintf("AX%d", i+2), row.F21)
		file.SetCellValue(sheetName, fmt.Sprintf("AY%d", i+2), row.F22)
		file.SetCellValue(sheetName, fmt.Sprintf("AZ%d", i+2), row.F23)
		file.SetCellValue(sheetName, fmt.Sprintf("BA%d", i+2), row.F24)
		file.SetCellValue(sheetName, fmt.Sprintf("BB%d", i+2), row.F25)
		file.SetCellValue(sheetName, fmt.Sprintf("BC%d", i+2), row.F26)
		file.SetCellValue(sheetName, fmt.Sprintf("BD%d", i+2), row.F27)
		file.SetCellValue(sheetName, fmt.Sprintf("BE%d", i+2), row.F301)
		file.SetCellValue(sheetName, fmt.Sprintf("BF%d", i+2), row.F302)
		file.SetCellValue(sheetName, fmt.Sprintf("BG%d", i+2), row.F303)
		file.SetCellValue(sheetName, fmt.Sprintf("BH%d", i+2), row.F401)
		file.SetCellValue(sheetName, fmt.Sprintf("BI%d", i+2), row.F402)
		file.SetCellValue(sheetName, fmt.Sprintf("BJ%d", i+2), row.F403)
		file.SetCellValue(sheetName, fmt.Sprintf("BK%d", i+2), row.F404)
		file.SetCellValue(sheetName, fmt.Sprintf("BL%d", i+2), row.F405)
		file.SetCellValue(sheetName, fmt.Sprintf("BM%d", i+2), row.F406)
		file.SetCellValue(sheetName, fmt.Sprintf("BN%d", i+2), row.F407)
		file.SetCellValue(sheetName, fmt.Sprintf("BO%d", i+2), row.F408)
		file.SetCellValue(sheetName, fmt.Sprintf("BP%d", i+2), row.F409)
		file.SetCellValue(sheetName, fmt.Sprintf("BQ%d", i+2), row.F410)
		file.SetCellValue(sheetName, fmt.Sprintf("BR%d", i+2), row.F411)
		file.SetCellValue(sheetName, fmt.Sprintf("BS%d", i+2), row.F412)
		file.SetCellValue(sheetName, fmt.Sprintf("BT%d", i+2), row.F413)
		file.SetCellValue(sheetName, fmt.Sprintf("BU%d", i+2), row.F414)
		file.SetCellValue(sheetName, fmt.Sprintf("BV%d", i+2), row.F415)
		file.SetCellValue(sheetName, fmt.Sprintf("BW%d", i+2), row.F416)
		file.SetCellValue(sheetName, fmt.Sprintf("BX%d", i+2), row.F6)
		file.SetCellValue(sheetName, fmt.Sprintf("BY%d", i+2), row.F7)
		file.SetCellValue(sheetName, fmt.Sprintf("BZ%d", i+2), row.F7a)
		file.SetCellValue(sheetName, fmt.Sprintf("CA%d", i+2), row.F1001)
		file.SetCellValue(sheetName, fmt.Sprintf("CB%d", i+2), row.F1002)
		file.SetCellValue(sheetName, fmt.Sprintf("CC%d", i+2), row.F1601)
		file.SetCellValue(sheetName, fmt.Sprintf("CD%d", i+2), row.F1602)
		file.SetCellValue(sheetName, fmt.Sprintf("CE%d", i+2), row.F1603)
		file.SetCellValue(sheetName, fmt.Sprintf("CF%d", i+2), row.F1604)
		file.SetCellValue(sheetName, fmt.Sprintf("CG%d", i+2), row.F1605)
		file.SetCellValue(sheetName, fmt.Sprintf("CH%d", i+2), row.F1606)
		file.SetCellValue(sheetName, fmt.Sprintf("CI%d", i+2), row.F1607)
		file.SetCellValue(sheetName, fmt.Sprintf("CJ%d", i+2), row.F1608)
		file.SetCellValue(sheetName, fmt.Sprintf("CK%d", i+2), row.F1609)
		file.SetCellValue(sheetName, fmt.Sprintf("CL%d", i+2), row.F1610)
		file.SetCellValue(sheetName, fmt.Sprintf("CM%d", i+2), row.F1611)
		file.SetCellValue(sheetName, fmt.Sprintf("CN%d", i+2), row.F1612)
		file.SetCellValue(sheetName, fmt.Sprintf("CO%d", i+2), row.F1613)
		file.SetCellValue(sheetName, fmt.Sprintf("CP%d", i+2), row.F1614)
		file.SetCellValue(sheetName, fmt.Sprintf("CQ%d", i+2), row.NamaAtasan)
		file.SetCellValue(sheetName, fmt.Sprintf("CR%d", i+2), row.HpAtasan)
		file.SetCellValue(sheetName, fmt.Sprintf("CS%d", i+2), row.EmailAtasan)
	}

	file.SetActiveSheet(index)

	excelFilePath := "pkts_report_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	if err := file.SaveAs(excelFilePath); err != nil {
		log.Println("ERROR: [PKTSService - ExportPKTSReport] Error while save excel file:", err)
		return nil, err
	}

	buff, err := file.WriteToBuffer()
	if err != nil {
		log.Println("ERROR: [PKTSService - ExportPKTSReport] Error while write to buffer:", err)
		return nil, err
	}

	return buff, nil
}

func (svc *PKTSService) FindPKTSRekapByProdi(ctx context.Context, kodeprodi, tahunSidang string) ([]*entity.PKTSRekapByProdi, error) {
	res, err := svc.pktsRepository.FindPKTSRekapByProdi(ctx, kodeprodi, tahunSidang)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PKTSService - FindPKTSRekap] Error while find pkts rekap:", parseError.Message)
		return nil, err
	}

	return res, nil
}
