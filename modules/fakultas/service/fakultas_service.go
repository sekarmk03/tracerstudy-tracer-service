package service

import (
	"context"
	"log"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	"tracerstudy-tracer-service/modules/fakultas/entity"
	"tracerstudy-tracer-service/modules/fakultas/repository"
)

type FakultasService struct {
	cfg                config.Config
	fakultasRepository repository.FakultasRepositoryUseCase
}

type FakultasServiceUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Fakultas, error)
	FindFakultasByKode(ctx context.Context, kode string) (*entity.Fakultas, error)
	Create(ctx context.Context, kode, nama, akronim string) (*entity.Fakultas, error)
	Update(ctx context.Context, kode string, fields *entity.Fakultas) (*entity.Fakultas, error)
	Delete(ctx context.Context, kode string) error
}

func NewFakultasService(cfg config.Config, fakultasRepository repository.FakultasRepositoryUseCase) *FakultasService {
	return &FakultasService{
		cfg:                cfg,
		fakultasRepository: fakultasRepository,
	}
}

func (svc *FakultasService) FindAll(ctx context.Context, req any) ([]*entity.Fakultas, error) {
	res, err := svc.fakultasRepository.FindAll(ctx, req)
	if err != nil {
		log.Println("ERROR: [FakultasService - FindAll] Error while find all fakultas: ", err)
		return nil, err
	}

	return res, nil
}

func (svc *FakultasService) FindFakultasByKode(ctx context.Context, kode string) (*entity.Fakultas, error) {
	res, err := svc.fakultasRepository.FindFakultasByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasService - FindFakultasByKode] Error while find fakultas by kode: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *FakultasService) Create(ctx context.Context, kode, nama, akronim string) (*entity.Fakultas, error) {
	fakultas := &entity.Fakultas{
		Kode:    kode,
		Nama:    nama,
		Akronim: akronim,
	}

	res, err := svc.fakultasRepository.Create(ctx, fakultas)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasService - Create] Error while create fakultas: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *FakultasService) Update(ctx context.Context, kode string, fields *entity.Fakultas) (*entity.Fakultas, error) {
	fakultas, err := svc.fakultasRepository.FindFakultasByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasService - Update] Error while find fakultas by kode: ", parseError.Message)
		return nil, err
	}

	updatedMap := make(map[string]interface{})

	utils.AddItemToMap(updatedMap, "nama", fields.Nama)
	utils.AddItemToMap(updatedMap, "akronim", fields.Akronim)

	res, err := svc.fakultasRepository.Update(ctx, fakultas, updatedMap)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasService - Update] Error while update fakultas: ", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *FakultasService) Delete(ctx context.Context, kode string) error {
	_, err := svc.fakultasRepository.FindFakultasByKode(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasService - Delete] Error while find fakultas by kode: ", parseError.Message)
		return err
	}

	err = svc.fakultasRepository.Delete(ctx, kode)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [FakultasService - Delete] Error while delete fakultas: ", parseError.Message)
		return err
	}

	return nil
}
