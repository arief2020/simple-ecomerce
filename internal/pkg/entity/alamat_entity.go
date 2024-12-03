package entity

import (
	"gorm.io/gorm"
)

type Alamat struct {
	gorm.Model
	ID           uint           `gorm:"primarykey;autoIncrement" json:"id"`
	JudulAlamat  string         `gorm:"type:varchar(255)" json:"judul_alamat" validate:"required"`
	NamaPenerima string         `gorm:"type:varchar(255)" json:"nama_penerima" validate:"required"`
	NoTelp       string         `gorm:"type:varchar(255)" json:"no_telp" validate:"required"`
	DetailAlamat string         `gorm:"type:varchar(255)" json:"detail_alamat" validate:"required"`
	UserID       uint           `gorm:"type:uint;column:id_user" json:"id_user" validate:"required"`
	
	// CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	// DeletedAt    gorm.DeletedAt `gorm:"-" json:"deleted_at"`
}