package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID            uint   `gorm:"primarykey;autoIncrement" json:"id"`
	NamaProduk    string `gorm:"type:varchar(255)" json:"nama_produk" validate:"required"`
	Slug          string `gorm:"type:varchar(255)" json:"slug" validate:"required"`
	HargaReseller string `gorm:"type:varchar(255)" json:"harga_reseller" validate:"required"`
	HargaKonsumen string `gorm:"type:varchar(255)" json:"harga_konsumen" validate:"required"`
	Stok          string `gorm:"type:varchar(255)" json:"stok" validate:"required"`
	Deskripsi     string `gorm:"type:text" json:"deskripsi" validate:"required"`
	TokoID        uint   `gorm:"type:uint;column:id_toko" json:"id_toko" validate:"required"`
	CategoryID    uint   `gorm:"type:uint;column:id_category" json:"id_category" validate:"required"`

	FotoProduct []FotoProduct
	Category    Category `gorm:"foreignKey:CategoryID"`
	Toko        Toko
}
