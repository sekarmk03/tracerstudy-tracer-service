package entity

import "tracerstudy-tracer-service/pb"

type UserStudyRekap struct {
	Fakultas            string  `json:"fakultas"`
	KodeProdi           string  `json:"kode_prodi"`
	NamaProdi           string  `json:"nama_prodi"`
	Jenjang             string  `json:"jenjang"`
	AlumniCount         uint32  `json:"alumni_count"`
	PktsCount           uint32  `json:"pkts_count"`
	UserStudyCount      uint32  `json:"user_study_count"`
	UserStudyPercentage float64 `json:"user_study_percentage"`
}

func ConvertUserStudyRekapEntityToProto(usr *UserStudyRekap) *pb.UserStudyRekap {
	return &pb.UserStudyRekap{
		Fakultas:            usr.Fakultas,
		KodeProdi:           usr.KodeProdi,
		NamaProdi:           usr.NamaProdi,
		Jenjang:             usr.Jenjang,
		AlumniCount:         usr.AlumniCount,
		PktsCount:           usr.PktsCount,
		UserStudyCount:      usr.UserStudyCount,
		UserStudyPercentage: usr.UserStudyPercentage,
	}
}
