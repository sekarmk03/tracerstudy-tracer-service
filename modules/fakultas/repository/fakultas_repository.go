package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"tracerstudy-tracer-service/modules/fakultas/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type FakultasRepository struct {
	db *gorm.DB
}

func NewFakultasRepository(db *gorm.DB) *FakultasRepository {
	return &FakultasRepository{
		db: db,
	}
}

type FakultasRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.Fakultas, error)
	FindFakultasByKode(ctx context.Context, kodeFakultas string) (*entity.Fakultas, error)
	Create(ctx context.Context, req *entity.Fakultas) (*entity.Fakultas, error)
	Update(ctx context.Context, fakultas *entity.Fakultas, updatedFields map[string]interface{}) (*entity.Fakultas, error)
	Delete(ctx context.Context, kodeFakultas string) error
}

func (f *FakultasRepository) FindAll(ctx context.Context, req any) ([]*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "FakultasRepository - FindAll")
	defer span.End()

	var fakultas []*entity.Fakultas
	if err := f.db.Debug().WithContext(ctxSpan).Find(&fakultas).Error; err != nil {
		log.Println("ERROR: [FakultasRepository - FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return fakultas, nil
}

func (f *FakultasRepository) FindFakultasByKode(ctx context.Context, kodeFakultas string) (*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "FakultasRepository - FindFakultasByKode")
	defer span.End()

	var fakultas entity.Fakultas
	if err := f.db.Debug().WithContext(ctxSpan).Where("kode = ?", kodeFakultas).First(&fakultas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [FakultasRepository - FindFakultasByKode] Record not found for kode", kodeFakultas)
			return nil, status.Errorf(codes.NotFound, "record not found for kode %s", kodeFakultas)
		}
		log.Println("ERROR: [FakultasRepository - FindFakultasByKode] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &fakultas, nil
}

func (f *FakultasRepository) Create(ctx context.Context, req *entity.Fakultas) (*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "FakultasRepository - Create")
	defer span.End()

	if err := f.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		log.Println("ERROR: [FakultasRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (f *FakultasRepository) Update(ctx context.Context, fakultas *entity.Fakultas, updatedFields map[string]interface{}) (*entity.Fakultas, error) {
	ctxSpan, span := trace.StartSpan(ctx, "FakultasRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()
	if err := f.db.Debug().WithContext(ctxSpan).Model(&fakultas).Where("kode", fakultas.Kode).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidValue) {
			log.Println("ERROR: [FakultasRepository - Update] Invalid value:", err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid value: %v", err)
		}
		log.Println("ERROR: [FakultasRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return fakultas, nil
}

func (f *FakultasRepository) Delete(ctx context.Context, kodeFakultas string) error {
	ctxSpan, span := trace.StartSpan(ctx, "FakultasRepository - Delete")
	defer span.End()

	if err := f.db.Debug().WithContext(ctxSpan).Where("kode = ?", kodeFakultas).Delete(&entity.Fakultas{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [FakultasRepository - Delete] Record not found for kode", kodeFakultas)
			return status.Errorf(codes.NotFound, "record not found for kode %s", kodeFakultas)
		}
		log.Println("ERROR: [FakultasRepository - Delete] Internal server error:", err)
		return status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return nil
}