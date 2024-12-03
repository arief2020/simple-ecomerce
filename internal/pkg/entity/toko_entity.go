package entity

import (
	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	ID        uint `gorm:"primarykey;autoIncrement" json:"id"`
	NamaToko  *string `gorm:"type:varchar(255);null" json:"nama_toko" validate:"required"`
	UrlFoto   *string `gorm:"type:varchar(255);null" json:"url_foto" validate:"required"`
	UserID    uint `gorm:"type:uint;column:id_user;unique" json:"id_user" validate:"required"`
	User      User 
	
	// CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"-" json:"deleted_at"`
	Product   []Product
}