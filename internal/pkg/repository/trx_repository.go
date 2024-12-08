package repository

import (
	"context"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/utils"

	"gorm.io/gorm"
)

type TrxRepository interface {
	CreateTrx(ctx context.Context, trx entity.Trx) (entity.Trx, error)
	CreateDetailTrx(ctx context.Context, detailTrx entity.DetailTrx) (entity.DetailTrx, error)
	GetAllTransaction(ctx context.Context, req dto.AllTransactionReq, userId uint) ([]entity.Trx, int64, error)

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
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, "Error Create Trx")
		return trx, err
	}
	return trx, nil
}

func (r *TrxRepositoryImpl) CreateDetailTrx(ctx context.Context, detailTrx entity.DetailTrx) (entity.DetailTrx, error) {
	if err := r.db.Create(&detailTrx).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, "Error Create Detail Trx")
		return detailTrx, err
	}
	return detailTrx, nil
}

func (r *TrxRepositoryImpl) GetTrxByID(ctx context.Context, trxId uint, userId uint) (entity.Trx, error) {
	var trx entity.Trx

	if err := r.db.Debug().
		Preload("Alamat").
		Preload("DetailTrx.LogProduct.Toko").
		Preload("DetailTrx.LogProduct.Category").
		Preload("DetailTrx.LogProduct.Product.FotoProduct").
		Where("id = ? AND id_user = ?", trxId, userId).
		First(&trx).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, "Error Get Trx By ID")
		return trx, err
	}

	return trx, nil
}

func (r *TrxRepositoryImpl) GetAllTransaction(ctx context.Context, req dto.AllTransactionReq, userId uint) ([]entity.Trx, int64, error) {
	var transactions []entity.Trx
	var total int64

	query := r.db.Debug().
		Preload("Alamat").
		Preload("DetailTrx.LogProduct.Toko").
		Preload("DetailTrx.LogProduct.Category").
		Preload("DetailTrx.LogProduct.Product.FotoProduct").
		Where("id_user = ?", userId)

	if req.Search != "" {
		query = query.Where("kode_invoice LIKE ?", "%"+req.Search+"%")
	}

	if err := query.Model(&entity.Trx{}).Count(&total).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, "Error Count Transaction")
		return nil, 0, err
	}

	offset := (req.Page - 1) * req.Limit
	if err := query.Limit(req.Limit).Offset(offset).Find(&transactions).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, "Error Get All Transaction")
		return nil, 0, err
	}

	return transactions, total, nil
}
