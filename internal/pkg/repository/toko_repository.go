package repository

import (
	"context"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/utils"

	"gorm.io/gorm"
)

type TokoRepository interface {
	GetAllToko(ctx context.Context, params dto.TokoFilter) (tokos []entity.Toko, err error)
	GetTokoById(ctx context.Context, id uint) (entity.Toko, error)
	GetTokoByUserId(ctx context.Context, id uint) (entity.Toko, error)
	CreateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error)
	UpdateToko(ctx context.Context, id uint, nama_toko string, url_foto string) error
}

type TokoRepositoryImpl struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{
		db: db,
	}
}

func (r *TokoRepositoryImpl) GetTokoById(ctx context.Context, id uint) (entity.Toko, error) {
	var toko entity.Toko
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&toko).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By ID")
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepositoryImpl) CreateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error) {
	if err := r.db.WithContext(ctx).Create(&toko).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Toko")
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepositoryImpl) GetAllToko(ctx context.Context, params dto.TokoFilter) (tokos []entity.Toko, err error) {
	db := r.db

	if params.Nama != "" {
		db = db.Where("nama_toko like ?", "%"+params.Nama+"%")
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Page).Find(&tokos).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Toko")
		return tokos, err
	}
	return tokos, nil
}

func (r *TokoRepositoryImpl) UpdateToko(ctx context.Context, id uint, nama_toko string, url_foto string) error {
	var toko entity.Toko
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&toko).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By ID")
		return err
	}
	toko.NamaToko = &nama_toko
	toko.UrlFoto = &url_foto
	if err := r.db.WithContext(ctx).Save(&toko).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update Toko")
		return err
	}
	return nil
}

func (r *TokoRepositoryImpl) GetTokoByUserId(ctx context.Context, id uint) (entity.Toko, error) {
	var toko entity.Toko
	if err := r.db.WithContext(ctx).Where("id_user = ?", id).First(&toko).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By User ID")
		return toko, err
	}
	return toko, nil
}
