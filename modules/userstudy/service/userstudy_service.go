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
	"tracerstudy-tracer-service/modules/userstudy/entity"
	"tracerstudy-tracer-service/modules/userstudy/repository"

	"github.com/xuri/excelize/v2"
)

type UserStudyService struct {
	cfg                 config.Config
	userStudyRepository repository.UserStudyRepositoryUseCase
}

type UserStudyServiceUseCase interface {
	FindAll(ctx context.Context, limit, page uint32) ([]*entity.UserStudy, int64, error)
	FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error)
	Update(ctx context.Context, nim, emailResponden, hpResponden string, fields *entity.UserStudy) (*entity.UserStudy, error)
	Create(ctx context.Context, namaResponden, emailResponden, hpResponden, namaInstansi, jabatan, alamatInstansi, nimLulusan, namaLulusan, prodiLulusan, tahunLulusan string) (*entity.UserStudy, error)
	ExportUSReport(ctx context.Context) (*bytes.Buffer, error)
	FindUserStudyRekap(ctx context.Context) ([]*entity.UserStudyRekap, error)
	FindUserStudyRekapByProdi(ctx context.Context, kodeProdi string) ([]*entity.UserStudyRekapByProdi, error)
}

func NewUserStudyService(cfg config.Config, userStudyRepository repository.UserStudyRepositoryUseCase) *UserStudyService {
	return &UserStudyService{
		cfg:                 cfg,
		userStudyRepository: userStudyRepository,
	}
}

func (svc *UserStudyService) FindAll(ctx context.Context, limit, page uint32) ([]*entity.UserStudy, int64, error) {
	offset := (page - 1) * limit

	res, totalRecords, err := svc.userStudyRepository.FindAll(ctx, int(limit), int(offset))
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - FindAll] Error while find all user study:", parseError.Message)
		return nil, 0, err
	}

	return res, totalRecords, nil
}

func (svc *UserStudyService) FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error) {
	res, err := svc.userStudyRepository.FindByNim(ctx, nim, emailResponden, hpResponden)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - FindByNim] Error while find user study by nim:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) Update(ctx context.Context, nim, emailResponden, hpResponden string, fields *entity.UserStudy) (*entity.UserStudy, error) {
	usrStd, err := svc.userStudyRepository.FindByNim(ctx, nim, emailResponden, hpResponden)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - FindByNim] Error while find user study by nim:", parseError.Message)
		return nil, err
	}

	updateMap := map[string]interface{}{
		"updated_at": time.Now(),
	}

	utils.AddItemToMap(updateMap, "nama_instansi", fields.NamaInstansi)
	utils.AddItemToMap(updateMap, "jabatan", fields.Jabatan)
	utils.AddItemToMap(updateMap, "alamat_instansi", fields.AlamatInstansi)
	utils.AddItemToMap(updateMap, "lama_mengenal_lulusan", fields.LamaMengenalLulusan)
	utils.AddItemToMap(updateMap, "etika", fields.Etika)
	utils.AddItemToMap(updateMap, "keahlian_bid_ilmu", fields.KeahlianBidIlmu)
	utils.AddItemToMap(updateMap, "bahasa_inggris", fields.BahasaInggris)
	utils.AddItemToMap(updateMap, "penggunaan_ti", fields.PenggunaanTi)
	utils.AddItemToMap(updateMap, "komunikasi", fields.Komunikasi)
	utils.AddItemToMap(updateMap, "kerjasama_tim", fields.KerjasamaTim)
	utils.AddItemToMap(updateMap, "pengembangan_diri", fields.PengembanganDiri)
	utils.AddItemToMap(updateMap, "kesiapan_terjun_masy", fields.KesiapanTerjunMasy)
	utils.AddItemToMap(updateMap, "keunggulan_lulusan", fields.KeunggulanLulusan)
	utils.AddItemToMap(updateMap, "kelemahan_lulusan", fields.KelemahanLulusan)
	utils.AddItemToMap(updateMap, "saran_peningkatan_kompetensi_lulusan", fields.SaranPeningkatanKompetensiLulusan)
	utils.AddItemToMap(updateMap, "saran_perbaikan_kurikulum", fields.SaranPerbaikanKurikulum)

	res, err := svc.userStudyRepository.Update(ctx, usrStd, updateMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - Update] Error while update user study:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) Create(ctx context.Context, namaResponden, emailResponden, hpResponden, namaInstansi, jabatan, alamatInstansi, nimLulusan, namaLulusan, prodiLulusan, tahunLulusan string) (*entity.UserStudy, error) {
	reqEntity := &entity.UserStudy{
		NamaResponden:  namaResponden,
		EmailResponden: emailResponden,
		HpResponden:    hpResponden,
		NamaInstansi:   namaInstansi,
		Jabatan:        jabatan,
		AlamatInstansi: alamatInstansi,
		NimLulusan:     nimLulusan,
		NamaLulusan:    namaLulusan,
		ProdiLulusan:   prodiLulusan,
		TahunLulusan:   tahunLulusan,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	res, err := svc.userStudyRepository.Create(ctx, reqEntity)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - Create] Error while create user study:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) ExportUSReport(ctx context.Context) (*bytes.Buffer, error) {
	rows, _, err := svc.userStudyRepository.FindAll(ctx, 0, 0)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - FindAll] Error while find all pkts:", parseError.Message)
		return nil, err
	}

	file := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := file.NewSheet(sheetName)
	if err != nil {
		log.Println("ERROR: [UserStudyService - ExportPKTSReport] Error while create new sheet:", err)
		return nil, err
	}

	file.SetCellValue(sheetName, "A1", "nama_responden")
	file.SetCellValue(sheetName, "B1", "email_responden")
	file.SetCellValue(sheetName, "C1", "hp_responden")
	file.SetCellValue(sheetName, "D1", "nama_instansi")
	file.SetCellValue(sheetName, "E1", "jabatan")
	file.SetCellValue(sheetName, "F1", "alamat_instansi")
	file.SetCellValue(sheetName, "G1", "nim_lulusan")
	file.SetCellValue(sheetName, "H1", "nama_lulusan")
	file.SetCellValue(sheetName, "I1", "tahun_lulusan")
	file.SetCellValue(sheetName, "J1", "lama_mengenal_lulusan")
	file.SetCellValue(sheetName, "K1", "etika")
	file.SetCellValue(sheetName, "L1", "keahlian_bid_ilmu")
	file.SetCellValue(sheetName, "M1", "bahasa_inggris")
	file.SetCellValue(sheetName, "N1", "penggunaan_ti")
	file.SetCellValue(sheetName, "O1", "komunikasi")
	file.SetCellValue(sheetName, "P1", "kerjasama_tim")
	file.SetCellValue(sheetName, "Q1", "pengembangan_diri")
	file.SetCellValue(sheetName, "R1", "kesiapan_terjun_masy")
	file.SetCellValue(sheetName, "S1", "keunggulan_lulusan")
	file.SetCellValue(sheetName, "T1", "kelemahan_lulusan")
	file.SetCellValue(sheetName, "U1", "saran_peningkatan_kompetensi_lulusan")
	file.SetCellValue(sheetName, "V1", "saran_perbaikan_kurikulum")
	file.SetCellValue(sheetName, "W1", "created_at")
	file.SetCellValue(sheetName, "X1", "updated_at")

	for i, row := range rows {
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), row.NamaResponden)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), row.EmailResponden)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), row.HpResponden)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), row.NamaInstansi)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), row.Jabatan)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), row.AlamatInstansi)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), row.NimLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), row.NamaLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), row.TahunLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), row.LamaMengenalLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), row.Etika)
		file.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), row.KeahlianBidIlmu)
		file.SetCellValue(sheetName, fmt.Sprintf("M%d", i+2), row.BahasaInggris)
		file.SetCellValue(sheetName, fmt.Sprintf("N%d", i+2), row.PenggunaanTi)
		file.SetCellValue(sheetName, fmt.Sprintf("O%d", i+2), row.Komunikasi)
		file.SetCellValue(sheetName, fmt.Sprintf("P%d", i+2), row.KerjasamaTim)
		file.SetCellValue(sheetName, fmt.Sprintf("Q%d", i+2), row.PengembanganDiri)
		file.SetCellValue(sheetName, fmt.Sprintf("R%d", i+2), row.KesiapanTerjunMasy)
		file.SetCellValue(sheetName, fmt.Sprintf("S%d", i+2), row.KeunggulanLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("T%d", i+2), row.KelemahanLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("U%d", i+2), row.SaranPeningkatanKompetensiLulusan)
		file.SetCellValue(sheetName, fmt.Sprintf("V%d", i+2), row.SaranPerbaikanKurikulum)
		file.SetCellValue(sheetName, fmt.Sprintf("W%d", i+2), (row.CreatedAt).Format("2006-01-02 15:04:05"))
		file.SetCellValue(sheetName, fmt.Sprintf("X%d", i+2), (row.UpdatedAt).Format("2006-01-02 15:04:05"))
	}

	file.SetActiveSheet(index)

	excelFilePath := "us_report_" + strconv.FormatInt(time.Now().Unix(), 10) + ".xlsx"
	if err := file.SaveAs(excelFilePath); err != nil {
		log.Println("ERROR: [UserStudyService - ExportUSReport] Error while save excel file:", err)
		return nil, err
	}

	buff, err := file.WriteToBuffer()
	if err != nil {
		log.Println("ERROR: [UserStudyService - ExportUSReport] Error while write to buffer:", err)
		return nil, err
	}

	return buff, nil
}

func (svc *UserStudyService) FindUserStudyRekap(ctx context.Context) ([]*entity.UserStudyRekap, error) {
	res, err := svc.userStudyRepository.FindUserStudyRekap(ctx)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - FindUserStudyRekap] Error while find user study rekap:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *UserStudyService) FindUserStudyRekapByProdi(ctx context.Context, kodeProdi string) ([]*entity.UserStudyRekapByProdi, error) {
	res, err := svc.userStudyRepository.FindUserStudyRekapByProdi(ctx, kodeProdi)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [UserStudyService - FindUserStudyRekapByProdi] Error while find user study rekap by prodi:", parseError.Message)
		return nil, err
	}

	return res, nil
}
