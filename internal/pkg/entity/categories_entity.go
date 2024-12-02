package entity

import (
	// "time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID       uint    `gorm:"primarykey;autoIncrement" json:"id"`
	NamaCategory string `gorm:"type:varchar(255)" json:"nama_category" validate:"required"`

	// CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"-" json:"deleted_at"`
	Product []Product
}