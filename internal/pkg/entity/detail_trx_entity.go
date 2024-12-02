package entity

import (
	"gorm.io/gorm"
)

type DetailTrx struct {
	gorm.Model
	ID       uint    `gorm:"primarykey;autoIncrement" json:"id"`
	TrxID    uint    `gorm:"type:uint;column:id_trx" json:"id_trx" validate:"required"`
	LogProductId uint    `gorm:"type:uint;column:id_log_product" json:"id_log_product" validate:"required"`
	Kuantitas int     `gorm:"type:int" json:"kuantitas" validate:"required"`
	HargaTotal int     `gorm:"type:int" json:"harga_total" validate:"required"`
	Trx      Trx
	LogProduct LogProduct
}