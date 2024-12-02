package entity

import (
	"gorm.io/gorm"
	"time"
)

type (
	User struct {
		// gorm.Model
		ID        uint `gorm:"primarykey;autoIncrement" json:"id"`
		Email    string `gorm:"unique" json:"email" validate:"required,email"`
		Nama     string `gorm:"type:varchar(255)" json:"name" validate:"required"`
		KataSandi string `gorm:"type:varchar(255)" json:"kata_sandi" validate:"required"`
		NoTelp   string `gorm:"unique;type:varchar(255);column:no_telp" json:"no_telp" validate:"required"`
		TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir" validate:"required"`
		JenisKelamin string `gorm:"type:varchar(255)" json:"jenis_kelamin" validate:"required"`
		Tentang    string `gorm:"type:text" json:"tentang" validate:"required"`
		Pekerjaan  string `gorm:"type:varchar(255)" json:"pekerjaan" validate:"required"`
		IdProvinsi string `gorm:"type:varchar(255)" json:"id_provinsi" validate:"required"`
		IdKota     string `gorm:"type:varchar(255)" json:"id_kota" validate:"required"`
		IsAdmin    bool `gorm:"type:bool" json:"is_admin"`
		CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"-"`
		Alamat    []Alamat
	}

	FilterUser struct {
		Limit, Offset int
		Title         string
	}
)
