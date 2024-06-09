package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"tracerstudy-tracer-service/modules/responden/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type RespondenRepository struct {
	db *gorm.DB
}

func NewRespondenRepository(db *gorm.DB) *RespondenRepository {
	return &RespondenRepository{
		db: db,
	}
}

type RespondenRepositoryUseCase interface {
	FindAll(ctx context.Context, limit, offset int) ([]*entity.Responden, int64, error)
	FindByNim(ctx context.Context, nim string) (*entity.Responden, error)
	Update(ctx context.Context, responden *entity.Responden, updatedFields map[string]interface{}) (*entity.Responden, error)
	Create(ctx context.Context, req *entity.Responden) (*entity.Responden, error)
	FindByNimList(ctx context.Context, nimList []string) ([]*entity.Responden, error)
}

func (r *RespondenRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.Responden, int64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindAll")
	defer span.End()

	var responden []*entity.Responden
	var totalRecords int64

	if err := r.db.Debug().WithContext(ctxSpan).
	Order("created_at desc").
	Limit(limit).
	Offset(offset).
	Find(&responden).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - FindAll] Internal server error:", err)
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	if err := r.db.Debug().WithContext(ctxSpan).Model(&entity.Responden{}).Count(&totalRecords).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - FindAll] Internal server error:", err)
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return responden, totalRecords, nil
}

func (r *RespondenRepository) FindByNim(ctx context.Context, nim string) (*entity.Responden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindByNim")
	defer span.End()

	var responden *entity.Responden
	if err := r.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&responden).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [RespondenRepository - FindByNim] Record not found for nim", nim)
			return nil, status.Errorf(codes.NotFound, "record not found for nim %s", nim)
		}
		log.Println("ERROR: [RespondenRepository - FindByNim] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return responden, nil
}

func (r *RespondenRepository) Update(ctx context.Context, responden *entity.Responden, updatedFields map[string]interface{}) (*entity.Responden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - Update")
	defer span.End()

	// var responden *entity.Responden
	// if err := r.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&responden).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		log.Println("WARNING: [RespondenRepository-Update] Record not found for nim", nim)
	// 		return nil, status.Errorf(codes.NotFound, "record not found for nim %s", nim)
	// 	}
	// 	log.Println("ERROR: [RespondenRepository-Update] Internal server error:", err)
	// 	return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	// }

	updatedFields["updated_at"] = time.Now()
	if err := r.db.Debug().WithContext(ctxSpan).Model(&responden).Updates(updatedFields).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return responden, nil
}

func (r *RespondenRepository) Create(ctx context.Context, req *entity.Responden) (*entity.Responden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - Create")
	defer span.End()

	if err := r.db.Debug().WithContext(ctxSpan).Select("nim").Create(&req).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("ERROR: [RespondenRepository - Create] Duplicated key:", err)
			return nil, status.Errorf(codes.AlreadyExists, "duplicated key %v", err)
		}
		log.Println("ERROR: [RespondenRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (r *RespondenRepository) FindByNimList(ctx context.Context, nimList []string) ([]*entity.Responden, error) {
	ctxSpan, span := trace.StartSpan(ctx, "RespondenRepository - FindByNimList")
	defer span.End()

	var responden []*entity.Responden
	if err := r.db.Debug().WithContext(ctxSpan).Where("nim IN (?)", nimList).Find(&responden).Error; err != nil {
		log.Println("ERROR: [RespondenRepository - FindByNimList] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return responden, nil
}
