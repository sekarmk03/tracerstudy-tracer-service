package entity

import (
	"time"

	"tracerstudy-tracer-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	UserStudyTableName = "user_study"
)

type UserStudy struct {
	Id                                uint64         `json:"id"`
	NamaResponden                     string         `json:"nama_responden"`
	EmailResponden                    string         `json:"email_responden"`
	HpResponden                       string         `json:"hp_responden"`
	NamaInstansi                      string         `json:"nama_instansi"`
	Jabatan                           string         `json:"jabatan"`
	AlamatInstansi                    string         `json:"alamat_instansi"`
	NimLulusan                        string         `json:"nim_lulusan"`
	NamaLulusan                       string         `json:"nama_lulusan"`
	ProdiLulusan                      string         `json:"prodi_lulusan"`
	TahunLulusan                      string         `json:"tahun_lulusan"`
	LamaMengenalLulusan               string         `json:"lama_mengenal_lulusan"`
	Etika                             string         `json:"etika"`
	KeahlianBidIlmu                   string         `json:"keahlian_bid_ilmu"`
	BahasaInggris                     string         `json:"bahasa_inggris"`
	PenggunaanTi                      string         `json:"penggunaan_ti"`
	Komunikasi                        string         `json:"komunikasi"`
	KerjasamaTim                      string         `json:"kerjasama_tim"`
	PengembanganDiri                  string         `json:"pengembangan_diri"`
	KesiapanTerjunMasy                string         `json:"kesiapan_terjun_masy"`
	KeunggulanLulusan                 string         `json:"keunggulan_lulusan"`
	KelemahanLulusan                  string         `json:"kelemahan_lulusan"`
	SaranPeningkatanKompetensiLulusan string         `json:"saran_peningkatan_kompetensi_lulusan"`
	SaranPerbaikanKurikulum           string         `json:"saran_peningkatan_kurikulum"`
	CreatedAt                         time.Time      `gorm:"type:timestamptz;not_null" json:"created_at"`
	UpdatedAt                         time.Time      `gorm:"type:timestamptz;not_null" json:"updated_at"`
	DeletedAt                         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func NewUserStudy(id uint64, namaResponden, emailResponden, hpResponden, namaInstansi, jabatan, alamatInstansi, nimLulusan, namaLulusan, prodiLulusan, tahunLulusan, lamaMengenalLulusan, etika, keahlianBidIlmu, bahasaInggris, penggunaanTi, komunikasi, kerjasamaTim, pengembanganDiri, kesiapanTerjunMasy, keunggulanLulusan, kelemahanLulusan, saranPeningkatanKompetensiLulusan, saranPerbaikanKurikulum, createdBy, updatedBy string) *UserStudy {
	return &UserStudy{
		Id:                                id,
		NamaResponden:                     namaResponden,
		EmailResponden:                    emailResponden,
		HpResponden:                       hpResponden,
		NamaInstansi:                      namaInstansi,
		Jabatan:                           jabatan,
		AlamatInstansi:                    alamatInstansi,
		NimLulusan:                        nimLulusan,
		NamaLulusan:                       namaLulusan,
		ProdiLulusan:                      prodiLulusan,
		TahunLulusan:                      tahunLulusan,
		LamaMengenalLulusan:               lamaMengenalLulusan,
		Etika:                             etika,
		KeahlianBidIlmu:                   keahlianBidIlmu,
		BahasaInggris:                     bahasaInggris,
		PenggunaanTi:                      penggunaanTi,
		Komunikasi:                        komunikasi,
		KerjasamaTim:                      kerjasamaTim,
		PengembanganDiri:                  pengembanganDiri,
		KesiapanTerjunMasy:                kesiapanTerjunMasy,
		KeunggulanLulusan:                 keunggulanLulusan,
		KelemahanLulusan:                  kelemahanLulusan,
		SaranPeningkatanKompetensiLulusan: saranPeningkatanKompetensiLulusan,
		SaranPerbaikanKurikulum:           saranPerbaikanKurikulum,
		CreatedAt:                         time.Now(),
		UpdatedAt:                         time.Now(),
	}
}

func (us *UserStudy) TableName() string {
	return UserStudyTableName
}

func ConvertEntityToProto(us *UserStudy) *pb.UserStudy {
	return &pb.UserStudy{
		Id:                                us.Id,
		NamaResponden:                     us.NamaResponden,
		EmailResponden:                    us.EmailResponden,
		HpResponden:                       us.HpResponden,
		NamaInstansi:                      us.NamaInstansi,
		Jabatan:                           us.Jabatan,
		AlamatInstansi:                    us.AlamatInstansi,
		NimLulusan:                        us.NimLulusan,
		NamaLulusan:                       us.NamaLulusan,
		ProdiLulusan:                      us.ProdiLulusan,
		TahunLulusan:                      us.TahunLulusan,
		LamaMengenalLulusan:               us.LamaMengenalLulusan,
		Etika:                             us.Etika,
		KeahlianBidIlmu:                   us.KeahlianBidIlmu,
		BahasaInggris:                     us.BahasaInggris,
		PenggunaanTi:                      us.PenggunaanTi,
		Komunikasi:                        us.Komunikasi,
		KerjasamaTim:                      us.KerjasamaTim,
		PengembanganDiri:                  us.PengembanganDiri,
		KesiapanTerjunMasy:                us.KesiapanTerjunMasy,
		KeunggulanLulusan:                 us.KeunggulanLulusan,
		KelemahanLulusan:                  us.KelemahanLulusan,
		SaranPeningkatanKompetensiLulusan: us.SaranPeningkatanKompetensiLulusan,
		SaranPerbaikanKurikulum:           us.SaranPerbaikanKurikulum,
		CreatedAt:                         timestamppb.New(us.CreatedAt),
		UpdatedAt:                         timestamppb.New(us.UpdatedAt),
	}
}

func ConvertProtoToEntity(us *pb.UserStudy) *UserStudy {
	return &UserStudy{
		Id:                                us.GetId(),
		NamaResponden:                     us.GetNamaResponden(),
		EmailResponden:                    us.GetEmailResponden(),
		HpResponden:                       us.GetHpResponden(),
		NamaInstansi:                      us.GetNamaInstansi(),
		Jabatan:                           us.GetJabatan(),
		AlamatInstansi:                    us.GetAlamatInstansi(),
		NimLulusan:                        us.GetNimLulusan(),
		NamaLulusan:                       us.GetNamaLulusan(),
		ProdiLulusan:                      us.GetProdiLulusan(),
		TahunLulusan:                      us.GetTahunLulusan(),
		LamaMengenalLulusan:               us.GetLamaMengenalLulusan(),
		Etika:                             us.GetEtika(),
		KeahlianBidIlmu:                   us.GetKeahlianBidIlmu(),
		BahasaInggris:                     us.GetBahasaInggris(),
		PenggunaanTi:                      us.GetPenggunaanTi(),
		Komunikasi:                        us.GetKomunikasi(),
		KerjasamaTim:                      us.GetKerjasamaTim(),
		PengembanganDiri:                  us.GetPengembanganDiri(),
		KesiapanTerjunMasy:                us.GetKesiapanTerjunMasy(),
		KeunggulanLulusan:                 us.GetKeunggulanLulusan(),
		KelemahanLulusan:                  us.GetKelemahanLulusan(),
		SaranPeningkatanKompetensiLulusan: us.GetSaranPeningkatanKompetensiLulusan(),
		SaranPerbaikanKurikulum:           us.GetSaranPerbaikanKurikulum(),
		CreatedAt:                         us.GetCreatedAt().AsTime(),
		UpdatedAt:                         us.GetUpdatedAt().AsTime(),
	}
}
