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

func (r *UserStudy) TableName() string {
	return UserStudyTableName
}

func ConvertEntityToProto(r *UserStudy) *pb.UserStudy {
	return &pb.UserStudy{
		Id:                                r.Id,
		NamaResponden:                     r.NamaResponden,
		EmailResponden:                    r.EmailResponden,
		HpResponden:                       r.HpResponden,
		NamaInstansi:                      r.NamaInstansi,
		Jabatan:                           r.Jabatan,
		AlamatInstansi:                    r.AlamatInstansi,
		NimLulusan:                        r.NimLulusan,
		NamaLulusan:                       r.NamaLulusan,
		ProdiLulusan:                      r.ProdiLulusan,
		TahunLulusan:                      r.TahunLulusan,
		LamaMengenalLulusan:               r.LamaMengenalLulusan,
		Etika:                             r.Etika,
		KeahlianBidIlmu:                   r.KeahlianBidIlmu,
		BahasaInggris:                     r.BahasaInggris,
		PenggunaanTi:                      r.PenggunaanTi,
		Komunikasi:                        r.Komunikasi,
		KerjasamaTim:                      r.KerjasamaTim,
		PengembanganDiri:                  r.PengembanganDiri,
		KesiapanTerjunMasy:                r.KesiapanTerjunMasy,
		KeunggulanLulusan:                 r.KeunggulanLulusan,
		KelemahanLulusan:                  r.KelemahanLulusan,
		SaranPeningkatanKompetensiLulusan: r.SaranPeningkatanKompetensiLulusan,
		SaranPerbaikanKurikulum:           r.SaranPerbaikanKurikulum,
		CreatedAt:                         timestamppb.New(r.CreatedAt),
		UpdatedAt:                         timestamppb.New(r.UpdatedAt),
	}
}

func ConvertProtoToEntity(r *pb.UserStudy) *UserStudy {
	return &UserStudy{
		Id:                                r.GetId(),
		NamaResponden:                     r.GetNamaResponden(),
		EmailResponden:                    r.GetEmailResponden(),
		HpResponden:                       r.GetHpResponden(),
		NamaInstansi:                      r.GetNamaInstansi(),
		Jabatan:                           r.GetJabatan(),
		AlamatInstansi:                    r.GetAlamatInstansi(),
		NimLulusan:                        r.GetNimLulusan(),
		NamaLulusan:                       r.GetNamaLulusan(),
		ProdiLulusan:                      r.GetProdiLulusan(),
		TahunLulusan:                      r.GetTahunLulusan(),
		LamaMengenalLulusan:               r.GetLamaMengenalLulusan(),
		Etika:                             r.GetEtika(),
		KeahlianBidIlmu:                   r.GetKeahlianBidIlmu(),
		BahasaInggris:                     r.GetBahasaInggris(),
		PenggunaanTi:                      r.GetPenggunaanTi(),
		Komunikasi:                        r.GetKomunikasi(),
		KerjasamaTim:                      r.GetKerjasamaTim(),
		PengembanganDiri:                  r.GetPengembanganDiri(),
		KesiapanTerjunMasy:                r.GetKesiapanTerjunMasy(),
		KeunggulanLulusan:                 r.GetKeunggulanLulusan(),
		KelemahanLulusan:                  r.GetKelemahanLulusan(),
		SaranPeningkatanKompetensiLulusan: r.GetSaranPeningkatanKompetensiLulusan(),
		SaranPerbaikanKurikulum:           r.GetSaranPerbaikanKurikulum(),
		CreatedAt:                         r.GetCreatedAt().AsTime(),
		UpdatedAt:                         r.GetUpdatedAt().AsTime(),
	}
}
