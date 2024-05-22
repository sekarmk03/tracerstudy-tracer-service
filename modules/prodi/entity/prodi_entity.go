package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	ProdiTableName = "ref_prodi"
)

// Prodi represents the entity for the Prodi model.
type Prodi struct {
	Id              uint32         `json:"id"`
	Kode            string         `json:"kode"`
	KodeDikti       string         `json:"kode_dikti"`
	KodeIntegrasi   string         `json:"kode_integrasi"`
	Nama            string         `json:"nama"`
	Jenjang         string         `json:"jenjang"`
	KodeFakultas    string         `json:"kode_fakultas"`
	NamaFakultas    string         `json:"nama_fakultas"`
	AkronimFakultas string         `json:"akronim_fakultas"`
	CreatedAt       time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

func NewProdi(
	kode, kodeDikti, kodeIntegrasi, nama, jenjang, kodeFakultas, namaFakultas, akronimFakultas string,
) *Prodi {
	return &Prodi{
		Kode:            kode,
		KodeDikti:       kodeDikti,
		KodeIntegrasi:   kodeIntegrasi,
		Nama:            nama,
		Jenjang:         jenjang,
		KodeFakultas:    kodeFakultas,
		NamaFakultas:    namaFakultas,
		AkronimFakultas: akronimFakultas,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

func (r *Prodi) TableName() string {
	return ProdiTableName
}

func ConvertEntityToProto(p *Prodi) *pb.Prodi {
	return &pb.Prodi{
		Kode:            p.Kode,
		KodeDikti:       p.KodeDikti,
		KodeIntegrasi:   p.KodeIntegrasi,
		Nama:            p.Nama,
		Jenjang:         p.Jenjang,
		KodeFakultas:    p.KodeFakultas,
		NamaFakultas:    p.NamaFakultas,
		AkronimFakultas: p.AkronimFakultas,
		CreatedAt:       timestamppb.New(p.CreatedAt),
		UpdatedAt:       timestamppb.New(p.CreatedAt),
	}
}
