package entity

import (
	"time"
	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	PKTSTableName = "pkts"
)

type PKTS struct {
	Id                  uint64         `json:"id"`
	Nim                 string         `json:"nim"`
	KodeProdi           string         `json:"kode_prodi"`
	TahunSidang         string         `json:"tahun_sidang"`
	F8                  uint32         `json:"f8"`
	F504                uint32         `json:"f504"`
	F502                uint32         `json:"f502"`
	F506                uint32         `json:"f506"`
	F505                uint64         `json:"f505"`
	F5a1                string         `json:"f5a1"`
	F5a2                string         `json:"f5a2"`
	F1101               uint32         `json:"f1101"`
	F1102               string         `json:"f1102"`
	F5b                 string         `json:"f5b"`
	F5c                 uint32         `json:"f5c"`
	F5d                 uint32         `json:"f5d"`
	F18a                uint32         `json:"f18a"`
	F18b                string         `json:"f18b"`
	F18c                string         `json:"f18c"`
	F18d                string         `json:"f18d"`
	F1201               uint32         `json:"f1201"`
	F1202               string         `json:"f1202"`
	F14                 uint32         `json:"f14"`
	F15                 uint32         `json:"f15"`
	F1761               uint32         `json:"f1761"`
	F1762               uint32         `json:"f1762"`
	F1763               uint32         `json:"f1763"`
	F1764               uint32         `json:"f1764"`
	F1765               uint32         `json:"f1765"`
	F1766               uint32         `json:"f1766"`
	F1767               uint32         `json:"f1767"`
	F1768               uint32         `json:"f1768"`
	F1769               uint32         `json:"f1769"`
	F1770               uint32         `json:"f1770"`
	F1771               uint32         `json:"f1771"`
	F1772               uint32         `json:"f1772"`
	F1773               uint32         `json:"f1773"`
	F1774               uint32         `json:"f1774"`
	F21                 uint32         `json:"f21"`
	F22                 uint32         `json:"f22"`
	F23                 uint32         `json:"f23"`
	F24                 uint32         `json:"f24"`
	F25                 uint32         `json:"f25"`
	F26                 uint32         `json:"f26"`
	F27                 uint32         `json:"f27"`
	F301                uint32         `json:"f301"`
	F302                uint32         `json:"f302"`
	F303                uint32         `json:"f303"`
	F401                uint32         `json:"f401"`
	F402                uint32         `json:"f402"`
	F403                uint32         `json:"f403"`
	F404                uint32         `json:"f404"`
	F405                uint32         `json:"f405"`
	F406                uint32         `json:"f406"`
	F407                uint32         `json:"f407"`
	F408                uint32         `json:"f408"`
	F409                uint32         `json:"f409"`
	F410                uint32         `json:"f410"`
	F411                uint32         `json:"f411"`
	F412                uint32         `json:"f412"`
	F413                uint32         `json:"f413"`
	F414                uint32         `json:"f414"`
	F415                uint32         `json:"f415"`
	F416                string         `json:"f416"`
	F6                  uint32         `json:"f6"`
	F7                  uint32         `json:"f7"`
	F7a                 uint32         `json:"f7a"`
	F1001               uint32         `json:"f1001"`
	F1002               string         `json:"f1002"`
	F1601               uint32         `json:"f1601"`
	F1602               uint32         `json:"f1602"`
	F1603               uint32         `json:"f1603"`
	F1604               uint32         `json:"f1604"`
	F1605               uint32         `json:"f1605"`
	F1606               uint32         `json:"f1606"`
	F1607               uint32         `json:"f1607"`
	F1608               uint32         `json:"f1608"`
	F1609               uint32         `json:"f1609"`
	F1610               uint32         `json:"f1610"`
	F1611               uint32         `json:"f1611"`
	F1612               uint32         `json:"f1612"`
	F1613               uint32         `json:"f1613"`
	F1614               string         `json:"f1614"`
	NamaAtasan          string         `json:"nama_atasan"`
	HpAtasan            string         `json:"hp_atasan"`
	EmailAtasan         string         `json:"email_atasan"`
	TinggalSelamaKuliah string         `json:"tinggal_selama_kuliah"`
	Code                string         `json:"code"`
	MailSent            time.Time      `json:"mail_sent"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (p *PKTS) TableName() string {
	return PKTSTableName
}

func ConvertEntityToProto(p *PKTS) *pb.PKTS {
	return &pb.PKTS{
		Id:                  p.Id,
		Nim:                 p.Nim,
		KodeProdi:           p.KodeProdi,
		TahunSidang:         p.TahunSidang,
		F8:                  p.F8,
		F504:                p.F504,
		F502:                p.F502,
		F506:                p.F506,
		F505:                p.F505,
		F5A1:                p.F5a1,
		F5A2:                p.F5a2,
		F1101:               p.F1101,
		F1102:               p.F1102,
		F5B:                 p.F5b,
		F5C:                 p.F5c,
		F5D:                 p.F5d,
		F18A:                p.F18a,
		F18B:                p.F18b,
		F18C:                p.F18c,
		F18D:                p.F18d,
		F1201:               p.F1201,
		F1202:               p.F1202,
		F14:                 p.F14,
		F15:                 p.F15,
		F1761:               p.F1761,
		F1762:               p.F1762,
		F1763:               p.F1763,
		F1764:               p.F1764,
		F1765:               p.F1765,
		F1766:               p.F1766,
		F1767:               p.F1767,
		F1768:               p.F1768,
		F1769:               p.F1769,
		F1770:               p.F1770,
		F1771:               p.F1771,
		F1772:               p.F1772,
		F1773:               p.F1773,
		F1774:               p.F1774,
		F21:                 p.F21,
		F22:                 p.F22,
		F23:                 p.F23,
		F24:                 p.F24,
		F25:                 p.F25,
		F26:                 p.F26,
		F27:                 p.F27,
		F301:                p.F301,
		F302:                p.F302,
		F303:                p.F303,
		F401:                p.F401,
		F402:                p.F402,
		F403:                p.F403,
		F404:                p.F404,
		F405:                p.F405,
		F406:                p.F406,
		F407:                p.F407,
		F408:                p.F408,
		F409:                p.F409,
		F410:                p.F410,
		F411:                p.F411,
		F412:                p.F412,
		F413:                p.F413,
		F414:                p.F414,
		F415:                p.F415,
		F416:                p.F416,
		F6:                  p.F6,
		F7:                  p.F7,
		F7A:                 p.F7a,
		F1001:               p.F1001,
		F1002:               p.F1002,
		F1601:               p.F1601,
		F1602:               p.F1602,
		F1603:               p.F1603,
		F1604:               p.F1604,
		F1605:               p.F1605,
		F1606:               p.F1606,
		F1607:               p.F1607,
		F1608:               p.F1608,
		F1609:               p.F1609,
		F1610:               p.F1610,
		F1611:               p.F1611,
		F1612:               p.F1612,
		F1613:               p.F1613,
		F1614:               p.F1614,
		NamaAtasan:          p.NamaAtasan,
		HpAtasan:            p.HpAtasan,
		EmailAtasan:         p.EmailAtasan,
		TinggalSelamaKuliah: p.TinggalSelamaKuliah,
		Code:                p.Code,
		MailSent:            timestamppb.New(p.MailSent),
		CreatedAt:           timestamppb.New(p.CreatedAt),
		UpdatedAt:           timestamppb.New(p.UpdatedAt),
	}
}
