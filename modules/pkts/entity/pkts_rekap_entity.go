package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PKTSRekap struct {
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

func ConvertEntityPKTSRekapToProto(p *PKTSRekap) *pb.PKTSRekapByProdi {
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
