package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"tracerstudy-tracer-service/modules/provinsi/entity"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	KabKotaTableName = "ref_kabkota"
)

type KabKota struct {
	Id         uint32          `json:"id"`
	IdWil      string          `json:"id_wil"`
	Nama       string          `json:"nama"`
	IdIndukWil string          `json:"id_induk_wil"`
	Provinsi   entity.Provinsi `gorm:"foreignKey:IdIndukWil;references:IdWil" json:"provinsi"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt  `gorm:"index" json:"deleted_at"`
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
		Provinsi:   entity.ConvertEntityToProto(&k.Provinsi),
		CreatedAt:  timestamppb.New(k.CreatedAt),
		UpdatedAt:  timestamppb.New(k.UpdatedAt),
	}
}
