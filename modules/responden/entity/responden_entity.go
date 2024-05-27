package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	RespondenTableName = "responden"
)

type Responden struct {
	Id            uint32         `json:"id"`
	Nim           string         `json:"nim"`
	Nama          string         `json:"nama"`
	StatusUpdate  uint32         `json:"status_update"`
	JalurMasuk    string         `json:"jalur_masuk"`
	TahunMasuk    string         `json:"tahun_masuk"`
	LamaStudi     uint32         `json:"lama_studi"`
	KodeFakultas  string         `json:"kode_fakultas"`
	KodeProdi     string         `json:"kode_prodi"`
	JenisKelamin  string         `json:"jenis_kelamin"`
	Email         string         `json:"email"`
	Hp            string         `json:"hp"`
	Ipk           string         `json:"ipk"`
	TanggalSidang string         `json:"tanggal_sidang"`
	TahunSidang   string         `json:"tahun_sidang"`
	TanggalWisuda string         `json:"tanggal_wisuda"`
	Nik           string         `json:"nik"`
	Npwp          string         `json:"npwp"`
	CreatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewResponden(id uint32, nim, nama string, statusUpdate uint32, jlrMasuk, thnMasuk string, lamaStudi uint32, kodeFak, kodeProdi, jenisKel, email, hp, ipk, tglSidang, thnSidang, tglWisuda, nik, npwp string) *Responden {
	return &Responden{
		Id:            id,
		Nim:           nim,
		Nama:          nama,
		StatusUpdate:  statusUpdate,
		JalurMasuk:    jlrMasuk,
		TahunMasuk:    thnMasuk,
		LamaStudi:     lamaStudi,
		KodeFakultas:  kodeFak,
		KodeProdi:     kodeProdi,
		JenisKelamin:  jenisKel,
		Email:         email,
		Hp:            hp,
		Ipk:           ipk,
		TanggalSidang: tglSidang,
		TahunSidang:   thnSidang,
		TanggalWisuda: tglWisuda,
		Nik:           nik,
		Npwp:          npwp,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

func (r *Responden) TableName() string {
	return RespondenTableName
}

func ConvertEntityToProto(r *Responden) *pb.Responden {
	return &pb.Responden{
		Id:            r.Id,
		Nim:           r.Nim,
		Nama:          r.Nama,
		StatusUpdate:  r.StatusUpdate,
		JalurMasuk:    r.JalurMasuk,
		TahunMasuk:    r.TahunMasuk,
		LamaStudi:     r.LamaStudi,
		KodeFakultas:  r.KodeFakultas,
		KodeProdi:     r.KodeProdi,
		JenisKelamin:  r.JenisKelamin,
		Email:         r.Email,
		Hp:            r.Hp,
		Ipk:           r.Ipk,
		TanggalSidang: r.TanggalSidang,
		TahunSidang:   r.TahunSidang,
		TanggalWisuda: r.TanggalWisuda,
		Nik:           r.Nik,
		Npwp:          r.Npwp,
		CreatedAt:     timestamppb.New(r.CreatedAt),
		UpdatedAt:     timestamppb.New(r.UpdatedAt),
	}
}
