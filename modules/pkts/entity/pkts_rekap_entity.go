package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PKTSRekapByProdi struct {
	Nim           string    `json:"nim"`
	Nama          string    `json:"nama"`
	F8            uint32    `json:"f8"`
	Status        string    `json:"status"`
	Email         string    `json:"email"`
	Hp            string    `json:"hp"`
	TanggalSidang string    `json:"tanggal_sidang"`
	PktsStatus    string    `json:"pkts_status"`
	ProvKerja     string    `json:"prov_kerja"`
	UmpPkts       string    `json:"ump_pkts"`
	Penghasilan   string    `json:"penghasilan"`
	InputPkts     time.Time `json:"input_pkts"`
	UpdatePkts    time.Time `json:"update_pkts"`
}

func ConvertEntityPKTSRekapByProdiToProto(p *PKTSRekapByProdi) *pb.PKTSRekapByProdi {
	return &pb.PKTSRekapByProdi{
		Nim:           p.Nim,
		Nama:          p.Nama,
		F8:            p.F8,
		Status:        p.Status,
		Email:         p.Email,
		Hp:            p.Hp,
		TanggalSidang: p.TanggalSidang,
		PktsStatus:    p.PktsStatus,
		ProvKerja:     p.ProvKerja,
		UmpPkts:       p.UmpPkts,
		Penghasilan:   p.Penghasilan,
		InputPkts:     timestamppb.New(p.InputPkts),
		UpdatePkts:    timestamppb.New(p.UpdatePkts),
	}
}

type PKTSRekapByYear struct {
	Fakultas                string  `json:"fakultas"`
	KodeProdi               string  `json:"kode_prodi"`
	NamaProdi               string  `json:"nama_prodi"`
	Jenjang                 string  `json:"jenjang"`
	AlumniCount             uint32  `json:"alumni_count"`
	PktsCount               uint32  `json:"pkts_count"`
	PktsPercentage          float64 `json:"pkts_percentage"`
	StatusLanjutstudiCount  uint32  `json:"status_lanjutstudi_count"`
	StatusHasincomeCount    uint32  `json:"status_hasincome_count"`
	HasincomeUmpCount       uint32  `json:"hasincome_ump_count"`
	HasincomeUmpPercentage  float64 `json:"hasincome_ump_percentage"`
	StatusLainnyaCount      uint32  `json:"status_lainnya_count"`
	StatusLainnyaPercentage float64 `json:"status_lainnya_percentage"`
}

func ConvertEntityPKTSRekapByYearToProto(p *PKTSRekapByYear) *pb.PKTSRekapByYear {
	return &pb.PKTSRekapByYear{
		Fakultas:                p.Fakultas,
		KodeProdi:               p.KodeProdi,
		NamaProdi:               p.NamaProdi,
		Jenjang:                 p.Jenjang,
		AlumniCount:             p.AlumniCount,
		PktsCount:               p.PktsCount,
		PktsPercentage:          p.PktsPercentage,
		StatusLanjutstudiCount:  p.StatusLanjutstudiCount,
		StatusHasincomeCount:    p.StatusHasincomeCount,
		HasincomeUmpCount:       p.HasincomeUmpCount,
		HasincomeUmpPercentage:  p.HasincomeUmpPercentage,
		StatusLainnyaCount:      p.StatusLainnyaCount,
		StatusLainnyaPercentage: p.StatusLainnyaPercentage,
	}
}
