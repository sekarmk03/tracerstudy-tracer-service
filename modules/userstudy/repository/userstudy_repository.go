package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"tracerstudy-tracer-service/modules/userstudy/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserStudyRepository struct {
	db *gorm.DB
}

func NewUserStudyRepository(db *gorm.DB) *UserStudyRepository {
	return &UserStudyRepository{
		db: db,
	}
}

type UserStudyRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.UserStudy, error)
	FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error)
	Update(ctx context.Context, userStudy *entity.UserStudy, updatedFields map[string]interface{}) (*entity.UserStudy, error)
	Create(ctx context.Context, req *entity.UserStudy) (*entity.UserStudy, error)
	FindUserStudyRekap(ctx context.Context, req any) ([]*entity.UserStudyRekap, error)
}

func (r *UserStudyRepository) FindAll(ctx context.Context, req any) ([]*entity.UserStudy, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindAll")
	defer span.End()

	var userStudy []*entity.UserStudy
	if err := r.db.Debug().WithContext(ctxSpan).Order("created_at desc").Find(&userStudy).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudy, nil
}

func (r *UserStudyRepository) FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindByNim")
	defer span.End()

	var userStudy *entity.UserStudy
	if err := r.db.Debug().
		WithContext(ctxSpan).
		Where("nim_lulusan = ?", nim).
		Where("LOWER(email_responden) LIKE LOWER(?) OR LOWER(hp_responden) LIKE LOWER(?)", "%"+emailResponden+"%", "%"+hpResponden+"%").
		First(&userStudy).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [UserStudyRepository - FindByNim] Record not found for nim", nim)
			return nil, status.Errorf(codes.NotFound, "record not found for nim %s", nim)
		}
		log.Println("ERROR: [UserStudyRepository - FindByNim] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudy, nil
}

func (r *UserStudyRepository) Update(ctx context.Context, userStudy *entity.UserStudy, updatedFields map[string]interface{}) (*entity.UserStudy, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - Update")
	defer span.End()

	updatedFields["updated_at"] = time.Now()
	if err := r.db.Debug().WithContext(ctxSpan).Model(&userStudy).Updates(updatedFields).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudy, nil
}

func (r *UserStudyRepository) Create(ctx context.Context, req *entity.UserStudy) (*entity.UserStudy, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - Create")
	defer span.End()

	if err := r.db.Debug().WithContext(ctxSpan).Create(req).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (r *UserStudyRepository) FindUserStudyRekap(ctx context.Context, req any) ([]*entity.UserStudyRekap, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindUserStudyRekap")
	defer span.End()

	var userStudyRekap []*entity.UserStudyRekap
	query := `
		SELECT
		r.kode_prodi, p.nama AS nama_prodi, p.akronim_fakultas AS fakultas, p.jenjang,
		COUNT(DISTINCT r.nim) AS alumni_count,
		COUNT(DISTINCT pk.nim) AS pkts_count,
		COUNT(DISTINCT us.nim_lulusan) AS user_study_count,
		ROUND((COUNT(DISTINCT us.nim_lulusan) / COUNT(DISTINCT pk.nim)) * 100, 2) AS user_study_percentage
		FROM responden r
		JOIN ref_prodi p ON r.kode_prodi = p.kode
		LEFT JOIN pkts pk ON r.nim = pk.nim
		LEFT JOIN user_study us ON us.nim_lulusan = pk.nim
		GROUP BY r.kode_prodi;
	`
	if err := r.db.Debug().WithContext(ctxSpan).Raw(query).Scan(&userStudyRekap).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindUserStudyRekap] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudyRekap, nil
}
