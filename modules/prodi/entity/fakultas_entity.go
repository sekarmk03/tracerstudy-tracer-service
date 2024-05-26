package entity

import (
	"tracerstudy-tracer-service/pb"
)

const (
	FakultasTableName = "ref_prodi"
)

func (r *Fakultas) TableName() string {
	return FakultasTableName
}

type Fakultas struct {
	Kode    string `json:"kode"`
	Nama    string `json:"nama"`
	Akronim string `json:"akronim"`
}

func ConvertEntityFakultasToProto(p *Fakultas) *pb.Fakultas {
	return &pb.Fakultas{
		Kode:    p.Kode,
		Nama:    p.Nama,
		Akronim: p.Akronim,
	}
}
