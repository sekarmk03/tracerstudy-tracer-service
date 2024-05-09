package repository

import (
	"context"
	"errors"
	"log"
	"time"
	"tracerstudy-tracer-service/modules/pkts/entity"

	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

type PKTSRepository struct {
	db *gorm.DB
}

func NewPKTSRepository(db *gorm.DB) *PKTSRepository {
	return &PKTSRepository{
		db: db,
	}
}

type PKTSRepositoryUseCase interface {
	FindAll(ctx context.Context, req any) ([]*entity.PKTS, error)
	FindByNim(ctx context.Context, nim string) (*entity.PKTS, error)
	Create(ctx context.Context, req *entity.PKTS) (*entity.PKTS, error)
	Update(ctx context.Context, pkts *entity.PKTS, updatedFields map[string]interface{}) (*entity.PKTS, error)
	FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error)
	FindAllReport(ctx context.Context, tahunSidang string) ([]*entity.PKTSReport, error)
	FindPKTSRekap(ctx context.Context, kodeprodi string) ([]*entity.PKTSRekap, error)
}

func (p *PKTSRepository) FindAll(ctx context.Context, req any) ([]*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindAll")
	defer span.End()

	var pkts []*entity.PKTS
	if err := p.db.Debug().WithContext(ctxSpan).Order("created_at desc").Find(&pkts).Error; err != nil {
		log.Println("ERROR: [PKTSRepository - FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return pkts, nil
}

func (p *PKTSRepository) FindByNim(ctx context.Context, nim string) (*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindByNim")
	defer span.End()

	var pkts entity.PKTS
	if err := p.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&pkts).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("WARNING: [PKTSRepository - FindByNim] Record not found for nim", nim)
			return nil, status.Errorf(codes.NotFound, "record not found for nim %s", nim)
		}
		log.Println("ERROR: [PKTSRepository - FindByNim] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return &pkts, nil
}

func (p *PKTSRepository) Create(ctx context.Context, req *entity.PKTS) (*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - Create")
	defer span.End()

	if err := p.db.Debug().WithContext(ctxSpan).Create(&req).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Println("ERROR: [PKTSRepository - Create] Duplicated key:", err)
			return nil, status.Errorf(codes.AlreadyExists, "duplicated key %v", err)
		}
		log.Println("ERROR: [PKTSRepository - Create] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return req, nil
}

func (p *PKTSRepository) Update(ctx context.Context, pkts *entity.PKTS, updatedFields map[string]interface{}) (*entity.PKTS, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - Update")
	defer span.End()

	// var pkts *entity.PKTS
	// if err := p.db.Debug().WithContext(ctxSpan).Where("nim = ?", nim).First(&pkts).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		log.Println("ERROR: [PKTSRepository - Update] Record not found for nim:", nim)
	// 		return nil, status.Errorf(codes.NotFound, "record not found for nim %s", nim)
	// 	}
	// 	log.Println("ERROR: [PKTSRepository - Update] Internal server error:", err)
	// 	return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	// }

	updatedFields["updated_at"] = time.Now()
	if err := p.db.Debug().WithContext(ctxSpan).Model(&pkts).Updates(updatedFields).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidValue) {
			log.Println("ERROR: [PKTSRepository - Update] Invalid value:", err)
			return nil, status.Errorf(codes.InvalidArgument, "invalid value %v", err)
		}
		log.Println("ERROR: [PKTSRepository - Update] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return pkts, nil
}

func (p *PKTSRepository) FindByAtasan(ctx context.Context, namaA, hpA, emailA string) ([]*string, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindByAtasan")
	defer span.End()

	var nims []*string
	if err := p.db.Debug().WithContext(ctxSpan).
		Table(entity.PKTSTableName).
		Select("nim").
		Where("hp_atasan LIKE ?", "%"+hpA+"%").
		Or("LOWER(email_atasan) LIKE LOWER(?)", "%"+emailA+"%").
		Or("LOWER(nama_atasan) LIKE LOWER(?)", "%"+namaA+"%").
		Pluck("nim", &nims).
		Error; err != nil {
		log.Println("ERROR: [PKTSRepository - FindByAtasan] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	return nims, nil
}

func (p *PKTSRepository) FindAllReport(ctx context.Context, tahunSidang string) ([]*entity.PKTSReport, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindAll")
	defer span.End()

	var pkts []*entity.PKTSReport
	query := `
        SELECT pk.*, r.nama, r.jenis_kelamin, r.hp, r.email, r.tahun_sidang, r.nik, r.npwp, pk.kode_prodi, p.nama AS nama_prodi, p.kode_dikti, p.jenjang
        FROM pkts AS pk
        JOIN responden AS r ON pk.nim = r.nim
        JOIN ref_prodi AS p ON pk.kode_prodi = p.kode
		WHERE pk.tahun_sidang = ?;
    `
	if err := p.db.Debug().WithContext(ctxSpan).Raw(query, tahunSidang).Scan(&pkts).Error; err != nil {
		log.Println("ERROR: [PKTSRepository - FindAll] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return pkts, nil
}

func (p *PKTSRepository) FindPKTSRekap(ctx context.Context, kodeprodi string) ([]*entity.PKTSRekap, error) {
	ctxSpan, span := trace.StartSpan(ctx, "PKTSRepository - FindPKTSRekap")
	defer span.End()

	var pkts []*entity.PKTSRekap
	query := `
		SELECT r.nim, r.nama, pk.f8, r.email, r.hp, r.tanggal_sidang, p.nama AS prov_kerja, pk.f5_05 AS penghasilan, pk.created_at AS input_pkts, pk.updated_at AS update_pkts
		FROM pkts AS pk
		JOIN responden AS r ON pk.nim = r.nim
		JOIN ref_provinsi AS p ON pk.f5a1 = p.id_wil
		WHERE pk.kode_prodi = ?
		LIMIT 10;
	`
	if err := p.db.Debug().WithContext(ctxSpan).Raw(query, kodeprodi).Scan(&pkts).Error; err != nil {
		log.Println("ERROR: [PKTSRepository - FindPKTSRekap] Internal server error:", err)
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	codeToStatus := map[int]string{
		1: "Bekerja",
		2: "Belum Memungkinkan Bekerja",
		3: "Wiraswasta",
		4: "Melanjutkan Pendidikan",
		5: "Tidak Kerja tetapi sedang mencari kerja",
	}

	for _, p := range pkts {
		status, ok := codeToStatus[int(p.F8)]
		if !ok {
			status = ""
		}
		p.Status = status
	}

	return pkts, nil
}
