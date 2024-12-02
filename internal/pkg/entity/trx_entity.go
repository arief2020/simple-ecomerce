package entity

import (
	"gorm.io/gorm"
)

type Trx struct {
	gorm.Model
	ID       uint    `gorm:"primarykey;autoIncrement" json:"id"`
	UserID   uint    `gorm:"type:uint;column:id_user" json:"id_user" validate:"required"`
	AlamatID uint    `gorm:"type:uint;column:alamat_pengiriman" json:"id_alamat" validate:"required"`
	HargaTotal int     `gorm:"type:int" json:"harga_total" validate:"required"`
	KodeInvoice string  `gorm:"type:varchar(255)" json:"kode_invoice" validate:"required"`
	MethodBayar string  `gorm:"type:varchar(255)" json:"method_bayar" validate:"required"`
	User      User
	Alamat    Alamat
	DetailTrx []DetailTrx
}