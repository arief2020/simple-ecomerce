package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type TrxRepository interface {
	CreateTrx(ctx context.Context, trx entity.Trx) (entity.Trx, error)
	CreateDetailTrx(ctx context.Context, detailTrx entity.DetailTrx) (entity.DetailTrx, error)
	GetAllTrxByUserID(ctx context.Context, userID uint) ([]entity.Trx, error)
	GetTrxByID(ctx context.Context, trxId uint, userId uint) (entity.Trx, error)
}

type TrxRepositoryImpl struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &TrxRepositoryImpl{db: db}
}

func (r *TrxRepositoryImpl) CreateTrx(ctx context.Context, trx entity.Trx) (entity.Trx, error) {
	if err := r.db.Create(&trx).Error; err != nil {
		return trx, err
	}
	return trx, nil
}

func (r *TrxRepositoryImpl) CreateDetailTrx(ctx context.Context, detailTrx entity.DetailTrx) (entity.DetailTrx, error) {
	if err := r.db.Create(&detailTrx).Error; err != nil {
		return detailTrx, err
	}
	return detailTrx, nil
}

func (r *TrxRepositoryImpl) GetAllTrxByUserID(ctx context.Context, userID uint) ([]entity.Trx, error) {
	var trx []entity.Trx
	if err := r.db.Where("id_user = ?", userID).Find(&trx).Error; err != nil {
		return trx, err
	}
	return trx, nil
}

func (r *TrxRepositoryImpl) GetTrxByID(ctx context.Context, trxId uint, userId uint) (entity.Trx, error) {
	var trx entity.Trx
	if err := r.db.Where("id = ? AND id_user = ?", trxId, userId).Find(&trx).Error; err != nil {
		return trx, err
	}
	return trx, nil
}