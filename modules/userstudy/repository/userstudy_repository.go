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
	FindAll(ctx context.Context, limit, offset int) ([]*entity.UserStudy, int64, error)
	FindByNim(ctx context.Context, nim, emailResponden, hpResponden string) (*entity.UserStudy, error)
	Update(ctx context.Context, userStudy *entity.UserStudy, updatedFields map[string]interface{}) (*entity.UserStudy, error)
	Create(ctx context.Context, req *entity.UserStudy) (*entity.UserStudy, error)
	FindUserStudyRekap(ctx context.Context, limit, offset int) ([]*entity.UserStudyRekap, int64, error)
	FindUserStudyRekapByProdi(ctx context.Context, limit, offset int, kodeProdi string) ([]*entity.UserStudyRekapByProdi, int64, error)
}

func (r *UserStudyRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.UserStudy, int64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindAll")
	defer span.End()

	var userStudy []*entity.UserStudy
	var totalRecords int64

	if limit == 0 && offset == 0 {
		if err := r.db.Debug().WithContext(ctxSpan).Order("created_at desc").Find(&userStudy).Error; err != nil {
			log.Println("ERROR: [UserStudyRepository - FindAll] Internal server error:", err)
			return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
		}
	} else {
		if err := r.db.Debug().WithContext(ctxSpan).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&userStudy).Error; err != nil {
			log.Println("ERROR: [UserStudyRepository - FindAll] Internal server error:", err)
			return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
		}
	}

	if err := r.db.Debug().WithContext(ctxSpan).Model(&entity.UserStudy{}).Count(&totalRecords).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindAll] Internal server error:", err)
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudy, totalRecords, nil
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

func (r *UserStudyRepository) FindUserStudyRekap(ctx context.Context, limit, offset int) ([]*entity.UserStudyRekap, int64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindUserStudyRekap")
	defer span.End()

	var userStudyRekap []*entity.UserStudyRekap
	var totalRecords int64

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
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	countQuery := `
		SELECT COUNT(*)
		FROM (
			SELECT r.kode_prodi
			FROM responden r
			JOIN ref_prodi p ON r.kode_prodi = p.kode
			LEFT JOIN pkts pk ON r.nim = pk.nim
			LEFT JOIN user_study us ON us.nim_lulusan = pk.nim
			GROUP BY r.kode_prodi
		) AS count_table;
	`
	if err := r.db.Debug().WithContext(ctxSpan).Raw(countQuery).Scan(&totalRecords).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindUserStudyRekap] Internal server error:", err)
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudyRekap, totalRecords, nil
}

func (r *UserStudyRepository) FindUserStudyRekapByProdi(ctx context.Context, limit, offset int, kodeProdi string) ([]*entity.UserStudyRekapByProdi, int64, error) {
	ctxSpan, span := trace.StartSpan(ctx, "UserStudyRepository - FindUserStudyRekapByProdi")
	defer span.End()

	var userStudyRekapByProdi []*entity.UserStudyRekapByProdi
	var totalRecords int64

	countQuery := `
		SELECT COUNT(*)
		FROM responden r
		LEFT JOIN pkts pk ON r.nim = pk.nim
		LEFT JOIN user_study us ON r.nim = us.nim_lulusan
		WHERE pk.f8 IN (1, 3)
		AND r.kode_prodi = ?;
	`
	if err := r.db.Debug().WithContext(ctxSpan).Raw(countQuery, kodeProdi).Scan(&totalRecords).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindUserStudyRekapByProdi] Internal server error:", err)
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	query := `
		SELECT 
		r.nim,
		r.nama AS nama,
		r.tahun_sidang AS tahun_lulus,
		pk.nama_atasan AS pkts_nama_atasan,
		pk.email_atasan AS pkts_email_atasan,
		pk.f5b AS pkts_instansi,
		CASE 
			WHEN us.nim_lulusan IS NULL THEN 'Belum' 
			ELSE 'Sudah'
		END AS userstudy_status,
		us.nama_responden AS userstudy_nama,
		us.email_responden AS userstudy_email,
		us.nama_instansi AS userstudy_instansi,
		us.jabatan AS userstudy_jabatan,
		us.lama_mengenal_lulusan AS lama_mengenal_lulusan,
		CASE
			WHEN avg_grades.avg_grade IS NULL THEN NULL
			WHEN avg_grades.avg_grade >= 3.5 THEN 'sangat baik'
			WHEN avg_grades.avg_grade >= 2.5 THEN 'baik'
			WHEN avg_grades.avg_grade >= 1.5 THEN 'cukup'
			ELSE 'kurang'
		END AS average_grade
		FROM 
			responden r
		LEFT JOIN
			pkts pk ON r.nim = pk.nim
		LEFT JOIN 
			user_study us ON r.nim = us.nim_lulusan
		LEFT JOIN (
			SELECT 
				nim_lulusan,
				(CASE etika
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE keahlian_bid_ilmu
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE bahasa_inggris
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE penggunaan_ti
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE komunikasi
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE kerjasama_tim
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE pengembangan_diri
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END +
				CASE kesiapan_terjun_masy
					WHEN 'kurang' THEN 1
					WHEN 'cukup' THEN 2
					WHEN 'baik' THEN 3
					WHEN 'sangat baik' THEN 4
				END) / 8.0 AS avg_grade
			FROM 
				user_study
		) avg_grades ON r.nim = avg_grades.nim_lulusan
		WHERE pk.f8 IN (1, 3)
		AND r.kode_prodi = ?
		ORDER BY userstudy_status DESC;
	`
	if err := r.db.Debug().WithContext(ctxSpan).Raw(query, kodeProdi).Scan(&userStudyRekapByProdi).Error; err != nil {
		log.Println("ERROR: [UserStudyRepository - FindUserStudyRekapByProdi] Internal server error:", err)
		return nil, 0, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return userStudyRekapByProdi, totalRecords, nil
}
