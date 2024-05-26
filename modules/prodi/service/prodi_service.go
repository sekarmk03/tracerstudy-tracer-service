package service

import (
	"context"
	"log"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	"tracerstudy-tracer-service/modules/prodi/entity"
	"tracerstudy-tracer-service/modules/prodi/repository"
)

type ProdiService struct {
	cfg             config.Config
	prodiRepository repository.ProdiRepositoryUseCase
}

type ProdiServiceUseCase interface {
	FindAll(ctx context.Context) ([]*entity.Prodi, error)
	FindAllFakultas(ctx context.Context) ([]*entity.Fakultas, error)
	FindProdiByKode(ctx context.Context, kode string) (*entity.Prodi, error)
	Create(ctx context.Context, kode, kodeDikti, kodeFakultas, kodeIntegrasi, nama, jenjang, namaFakultas, akronimFakultas string) (*entity.Prodi, error)
	Update(ctx context.Context, kode string, fields *entity.Prodi) (*entity.Prodi, error)
	Delete(ctx context.Context, kode string) error
}

func NewProdiService(cfg config.Config, prodiRepository repository.ProdiRepositoryUseCase) *ProdiService {
	return &ProdiService{
		cfg:             cfg,
		prodiRepository: prodiRepository,
	}
}

func (svc *ProdiService) FindAll(ctx context.Context) ([]*entity.Prodi, error) {
	res, err := svc.prodiRepository.FindAll(ctx)
	if err != nil {
		log.Println("ERROR: [ProdiService - FindAll] Error while find all prodi: ", err)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) FindAllFakultas(ctx context.Context) ([]*entity.Fakultas, error) {
	res, err := svc.prodiRepository.FindAllFakultas(ctx)
	if err != nil {
		log.Println("ERROR: [ProdiService - FindAllFakultas] Error while find all fakultas: ", err)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) FindProdiByKode(ctx context.Context, kode string) (*entity.Prodi, error) {
	res, err := svc.prodiRepository.FindProdiByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - FindProdiByKode] Error while find prodi by kode: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) Create(ctx context.Context, kode, kodeDikti, kodeFakultas, kodeIntegrasi, nama, jenjang, namaFakultas, akronimFakultas string) (*entity.Prodi, error) {
	prodi := &entity.Prodi{
		Kode:            kode,
		KodeDikti:       kodeDikti,
		KodeIntegrasi:   kodeIntegrasi,
		Nama:            nama,
		Jenjang:         jenjang,
		KodeFakultas:    kodeFakultas,
		NamaFakultas:    namaFakultas,
		AkronimFakultas: akronimFakultas,
	}

	res, err := svc.prodiRepository.Create(ctx, prodi)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Create] Error while create prodi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) Update(ctx context.Context, kode string, fields *entity.Prodi) (*entity.Prodi, error) {
	prodi, err := svc.prodiRepository.FindProdiByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Update] Error while find prodi by kode: ", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "kode_dikti", fields.KodeDikti)
	utils.AddItemToMap(updatedMap, "kode_integrasi", fields.KodeIntegrasi)
	utils.AddItemToMap(updatedMap, "nama", fields.Nama)
	utils.AddItemToMap(updatedMap, "jenjang", fields.Jenjang)
	utils.AddItemToMap(updatedMap, "kode_fakultas", fields.KodeFakultas)
	utils.AddItemToMap(updatedMap, "nama_fakultas", fields.NamaFakultas)
	utils.AddItemToMap(updatedMap, "akronim_fakultas", fields.AkronimFakultas)

	res, err := svc.prodiRepository.Update(ctx, prodi, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Update] Error while update prodi: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *ProdiService) Delete(ctx context.Context, kode string) error {
	_, err := svc.prodiRepository.FindProdiByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Delete] Error while find prodi by kode: ", parseError.Message)
		return err
	}

	err = svc.prodiRepository.Delete(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [ProdiService - Delete] Error while delete prodi: ", parseError.Message)
		return err
	}

	return nil
}
