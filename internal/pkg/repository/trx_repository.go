package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type TrxRepository interface {
	CreateTrx(ctx context.Context, trx entity.Trx) (entity.Trx, error)
	CreateDetailTrx(ctx context.Context, detailTrx entity.DetailTrx) (entity.DetailTrx, error)
	// GetAllTrxByUserID(ctx context.Context, userID uint) ([]entity.Trx, error)
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

// func (r *TrxRepositoryImpl) GetAllTrxByUserID(ctx context.Context, userID uint) ([]entity.Trx, error) {
// 	var trx []entity.Trx
// 	if err := r.db.Where("id_user = ?", userID).Find(&trx).Error; err != nil {
// 		return trx, err
// 	}
// 	return trx, nil
// }

// func (r *TrxRepositoryImpl) GetTrxByID(ctx context.Context, trxId uint, userId uint) (res dto.TransactionResponse, error) {
// 	var trx entity.Trx
// 	if err := r.db.Debug().Where("id = ? AND id_user = ?", trxId, userId).First(&trx).Error; err != nil {
// 		return trx, err
// 	}
// 	return trx, nil
// }


func (r *TrxRepositoryImpl) GetTrxByID(ctx context.Context, trxId uint, userId uint) (entity.Trx, error) {
	var trx entity.Trx

	// Gunakan preload untuk eager loading relasi
	if err := r.db.Debug().
		Preload("Alamat").
		Preload("DetailTrx.LogProduct.Toko").
		Preload("DetailTrx.LogProduct.Category").
		Preload("DetailTrx.LogProduct.Product.FotoProduct").
		Where("id = ? AND id_user = ?", trxId, userId).
		First(&trx).Error; err != nil {
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

	// Jika terdapat parameter pencarian
	if req.Search != "" {
		query = query.Where("kode_invoice LIKE ?", "%"+req.Search+"%")
	}

	// Hitung total transaksi
	if err := query.Model(&entity.Trx{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination
	offset := (req.Page - 1) * req.Limit
	if err := query.Limit(req.Limit).Offset(offset).Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}