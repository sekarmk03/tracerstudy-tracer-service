package service

import (
	"context"
	"log"
	"time"
	"tracerstudy-tracer-service/common/config"
	"tracerstudy-tracer-service/common/errors"
	"tracerstudy-tracer-service/common/utils"
	"tracerstudy-tracer-service/modules/kabkota/entity"
	"tracerstudy-tracer-service/modules/kabkota/repository"
)

type KabKotaService struct {
	cfg               config.Config
	kabkotaRepository repository.KabKotaRepositoryUseCase
}

type KabKotaServiceUseCase interface {
	FindAll(ctx context.Context, page, limit uint32) ([]*entity.KabKota, int64, error)
	FindByIdWil(ctx context.Context, idWil string) (*entity.KabKota, error)
	Create(ctx context.Context, idWil, nama, idIndukWil string) (*entity.KabKota, error)
	Update(ctx context.Context, idWil string, fields *entity.KabKota) (*entity.KabKota, error)
	Delete(ctx context.Context, idWil string) error
}

func NewKabKotaService(cfg config.Config, kabkotaRepository repository.KabKotaRepositoryUseCase) *KabKotaService {
	return &KabKotaService{
		cfg:               cfg,
		kabkotaRepository: kabkotaRepository,
	}
}

func (svc *KabKotaService) FindAll(ctx context.Context, page, limit uint32) ([]*entity.KabKota, int64, error) {
	offset := (page - 1) * limit

	res, rows, err := svc.kabkotaRepository.FindAll(ctx, int(limit), int(offset))
	if err != nil {
		log.Println("ERROR: [KabKotaService - FindAll] Error while find all kabkota:", err)
		return nil, 0, err
	}

	return res, rows, nil
}

func (svc *KabKotaService) FindByIdWil(ctx context.Context, idWil string) (*entity.KabKota, error) {
	res, err := svc.kabkotaRepository.FindByIdWil(ctx, idWil)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaService - FindByIdWil] Error while find KabKota by IdWil:", parseError.Message)
		return nil, err
	}

	return res, nil
}

func (svc *KabKotaService) Create(ctx context.Context, idWil, nama, idIndukWil string) (*entity.KabKota, error) {
	kabkota := &entity.KabKota{
		IdWil:      idWil,
		Nama:       nama,
		IdIndukWil: idIndukWil,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	res, err := svc.kabkotaRepository.Create(ctx, kabkota)
	if err != nil {
		log.Println("ERROR: [KabKotaService - Create] Error while create KabKota:", err)
		return nil, err
	}

	return res, nil
}

func (svc *KabKotaService) Update(ctx context.Context, idWil string, fields *entity.KabKota) (*entity.KabKota, error) {
	kabkota, err := svc.kabkotaRepository.FindByIdWil(ctx, idWil)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaService - Update] Error while find KabKota by IdWil:", parseError.Message)
		return nil, err
	}

	updateMap := make(map[string]interface{})

	utils.AddItemToMap(updateMap, "nama", fields.Nama)
	utils.AddItemToMap(updateMap, "id_induk_wil", fields.IdIndukWil)

	res, err := svc.kabkotaRepository.Update(ctx, kabkota, updateMap)
	if err != nil {
		log.Println("ERROR: [KabKotaService - Update] Error while update KabKota:", err)
		return nil, err
	}

	return res, nil
}

func (svc *KabKotaService) Delete(ctx context.Context, idWil string) error {
	err := svc.kabkotaRepository.Delete(ctx, idWil)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [KabKotaService - Delete] Error while delete KabKota:", parseError.Message)
		return err
	}

	return nil
}
