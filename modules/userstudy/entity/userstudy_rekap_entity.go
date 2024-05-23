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

type UserStudyRekapByProdi struct {
	Nim                 string `json:"nim"`
	Nama                string `json:"nama"`
	TahunLulus          string `json:"tahun_lulus"`
	PktsNamaAtasan      string `json:"pkts_nama_atasan"`
	PktsEmailAtasan     string `json:"pkts_email_atasan"`
	PktsInstansi        string `json:"pkts_instansi"`
	UserstudyStatus     string `json:"userstudy_status"`
	UserstudyNama       string `json:"userstudy_nama"`
	UserstudyEmail      string `json:"userstudy_email"`
	UserstudyInstansi   string `json:"userstudy_instansi"`
	UserstudyJabatan    string `json:"userstudy_jabatan"`
	LamaMengenalLulusan string `json:"lama_mengenal_lulusan"`
	AverageGrade        string `json:"average_grade"`
}

func ConvertUserStudyRekapByProdiEntityToProto(usr *UserStudyRekapByProdi) *pb.UserStudyRekapByProdi {
	return &pb.UserStudyRekapByProdi{
		Nim:                 usr.Nim,
		Nama:                usr.Nama,
		TahunLulus:          usr.TahunLulus,
		PktsNamaAtasan:      usr.PktsNamaAtasan,
		PktsEmailAtasan:     usr.PktsEmailAtasan,
		PktsInstansi:        usr.PktsInstansi,
		UserstudyStatus:     usr.UserstudyStatus,
		UserstudyNama:       usr.UserstudyNama,
		UserstudyEmail:      usr.UserstudyEmail,
		UserstudyInstansi:   usr.UserstudyInstansi,
		UserstudyJabatan:    usr.UserstudyJabatan,
		LamaMengenalLulusan: usr.LamaMengenalLulusan,
		AverageGrade:        usr.AverageGrade,
	}
}
