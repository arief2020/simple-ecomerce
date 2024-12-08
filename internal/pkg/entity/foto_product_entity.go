package entity

import (
	// "time"

	"gorm.io/gorm"
)

// type FotoProduct struct {
// 	gorm.Model
// 	ID       uint    `gorm:"primarykey;autoIncrement" json:"id"`
// 	ProductID uint    `gorm:"type:uint;column:id_product" json:"id_product" validate:"required"`
// 	UrlFoto  string  `gorm:"type:varchar(255)" json:"url_foto" validate:"required"`

// 	// CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
// 	// UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
// 	// DeletedAt gorm.DeletedAt `gorm:"-" json:"deleted_at"`
// }


// type FotoProduct struct {
//     gorm.Model
//     UrlFoto  string  `gorm:"type:varchar(255)" json:"url_foto" validate:"required"`
//     ProductID uint   `gorm:"not null;column:id_product" json:"product_id"` // GORM akan membuat foreign key ini secara otomatis
// }


type FotoProduct struct {
    gorm.Model
    UrlFoto   string `gorm:"type:varchar(255)" json:"url_foto" validate:"required"`
    ProductID uint   `gorm:"not null;column:id_product" json:"product_id"` // Foreign key otomatis
}
