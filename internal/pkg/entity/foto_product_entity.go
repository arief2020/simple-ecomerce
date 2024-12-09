package entity

import (
	// "time"

	"gorm.io/gorm"
)

type FotoProduct struct {
	gorm.Model
	UrlFoto   string `gorm:"type:varchar(255)" json:"url_foto" validate:"required"`
	ProductID uint   `gorm:"not null;column:id_product" json:"product_id"`
}
