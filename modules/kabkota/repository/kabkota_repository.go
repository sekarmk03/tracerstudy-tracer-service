package repository

import (
	"context"
	"errors"
	"log"
	"time"
	kkentity "tracerstudy-tracer-service/modules/kabkota/entity"
	pentity "tracerstudy-tracer-service/modules/provinsi/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type KabKotaRepository struct {
	db *gorm.DB
}

func NewKabKotaRepository(db *gorm.DB) *KabKotaRepository {
	return &KabKotaRepository{
		db: db,
	}
}

type KabKotaRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*kkentity.KabKota, error)
	FindByIdWil(ctx context.Context, idWil string) (*kkentity.KabKota, error)
	Create(ctx context.Context, req *kkentity.KabKota) (*kkentity.KabKota, error)
	Update(ctx context.Context, kabkota *kkentity.KabKota, updatedFields map[string]interface{}) (*kkentity.KabKota, error)
	Delete(ctx context.Context, idWil string) error
}

func (k *KabKotaRepository) FindAll(ctx context.Context, req any) ([]*kkentity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - FindAll")
	defer span.End()

	var kabkota []*kkentity.KabKota
	if err := k.db.Debug().WithContext(ctxSpan).Find(&kabkota).Error; err != nil {
		log.Println("ERROR: [KabKotaRepository - FindAll] Internal error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	for _, kk := range kabkota {
		var provinsi pentity.Provinsi
		if err := k.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", kk.IdIndukWil).First(&provinsi).Error; err != nil {
			log.Println("ERROR: [KabKotaRepository - FindAll] Failed to fetch Provinsi data:", err)
			return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
		}

		kk.Provinsi = provinsi
	}

	return kabkota, nil
}

func (k *KabKotaRepository) FindByIdWil(ctx context.Context, idWil string) (*kkentity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - FindByIdWil")
	defer span.End()

	var kabkota kkentity.KabKota
	if err := k.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", idWil).First(&kabkota).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [KabKotaRepository - FindByIdWil] Record not found for idWil", idWil)
			return nil, status.Errorf(codes.NotFound, "record not found for idWil %s", idWil)
		}
		log.Println("ERROR: [KabKotaRepository - FindByIdWil] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &kabkota, nil
}

func (k *KabKotaRepository) Create(ctx context.Context, req *kkentity.KabKota) (*kkentity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - Create")
	defer span.End()

	if err := k.db.Debug().WithContext(ctxSpan).Create(&req).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("ERROR: [KabKotaRepository - Create] Duplicated key:", err)
			return nil, status.Errorf(codes.AlreadyExists, "duplicated key: %v", err)
		}
		log.Println("ERROR: [KabKotaRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (k *KabKotaRepository) Update(ctx context.Context, kabkota *kkentity.KabKota, updatedFields map[string]interface{}) (*kkentity.KabKota, error) {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()
	if err := k.db.Debug().WithContext(ctxSpan).Model(&kabkota).Where("id_wil", kabkota.IdWil).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidValue) {
			log.Println("ERROR: [KabKotaRepository - Update] Invalid value:", err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid value: %v", err)
		}
		log.Println("ERROR: [KabKotaRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return kabkota, nil
}

func (k *KabKotaRepository) Delete(ctx context.Context, idWil string) error {
	ctxSpan, span := trace.StartSpan(ctx, "KabKotaRepository - Delete")
	defer span.End()

	var kabkota kkentity.KabKota
	if err := k.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", idWil).First(&kabkota).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [KabKotaRepository - DeleteByIdWil] Record not found for idWil", idWil)
			return status.Errorf(codes.NotFound, "record not found for idWil %s", idWil)
		}
		log.Println("ERROR: [KabKotaRepository - DeleteByIdWil] Internal server error:", err)
		return status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	// Delete the record
	if err := k.db.Debug().WithContext(ctxSpan).Where("id_wil = ?", idWil).Delete(&kabkota).Error; err != nil {
		log.Println("ERROR: [KabKotaRepository - DeleteByIdWil] Internal server error:", err)
		return status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return nil
}
