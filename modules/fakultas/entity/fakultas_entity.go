package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	FakultasTableName = "ref_fakultas"
)

type Fakultas struct {
	Id        uint32         `json:"id"`
	Kode      string         `json:"kode"`
	Nama      string         `json:"nama"`
	Akronim   string         `json:"akronim"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewFakultas(id uint32, kode, nama, akronim string) *Fakultas {
	return &Fakultas{
		Id:        id,
		Kode:      kode,
		Nama:      nama,
		Akronim:   akronim,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (f *Fakultas) TableName() string {
	return FakultasTableName
}

func ConvertEntityToProto(f *Fakultas) *pb.Fakultas {
	return &pb.Fakultas{
		Id:        f.Id,
		Kode:      f.Kode,
		Nama:      f.Nama,
		Akronim:   f.Akronim,
		CreatedAt: timestamppb.New(f.CreatedAt),
		UpdatedAt: timestamppb.New(f.UpdatedAt),
	}
}
