package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"
)

const (
	KabKotaTableName = "ref_kabkota"
)

type KabKota struct {
	Id         uint32    `json:"id"`
	IdWil      string    `json:"id_wil"`
	Nama       string    `json:"nama"`
	IdIndukWil string    `json:"id_induk_wil"`
	CreatedAt  time.Time `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamptz;not_null" json:"updated_at"`
}

func NewKabKota(id uint32, idWil, nama, idIndukWil string) *KabKota {
	return &KabKota{
		Id:         id,
		IdWil:      idWil,
		Nama:       nama,
		IdIndukWil: idIndukWil,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (k *KabKota) TableName() string {
	return KabKotaTableName
}

func ConvertEntityToProto(k *KabKota) *pb.KabKota {
	return &pb.KabKota{
		Id:         k.Id,
		IdWil:      k.IdWil,
		Nama:       k.Nama,
		IdIndukWil: k.IdIndukWil,
		CreatedAt:  k.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  k.UpdatedAt.Format(time.RFC3339),
	}
}
