package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type TokoRepository interface {
	GetAllToko(ctx context.Context) ([]entity.Toko, error)
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
	query := "SELECT * FROM tokos WHERE id_user = ? LIMIT 1"
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&toko).Error; err != nil {
		return toko, err
	}
	return toko, nil
}

func (r *TokoRepositoryImpl) CreateToko(ctx context.Context, toko entity.Toko) (entity.Toko, error) {
	if err := r.db.WithContext(ctx).Create(&toko).Error; err != nil {
		return toko, err
	}
	return toko, nil
}

// func (r *TokoRepositoryImpl) GetAllToko(ctx context.Context) ([]*entity.Toko, error) {
// 	var toko []*entity.Toko
// 	if err := r.db.WithContext(ctx).Find(&toko).Error; err != nil {
// 		return toko, err
// 	}
// 	return toko, nil
// }

func (r *TokoRepositoryImpl) GetAllToko(ctx context.Context) ([]entity.Toko, error) {
	var tokos []entity.Toko
	query := "SELECT * FROM tokos"
	if err := r.db.WithContext(ctx).Raw(query).Scan(&tokos).Error; err != nil {
		return tokos, err
	}
	return tokos, nil
}


func (r *TokoRepositoryImpl) UpdateToko(ctx context.Context, id uint, nama_toko string, url_foto string) error {
	query := "UPDATE tokos SET nama_toko = ?, url_foto = ? WHERE id_user = ?"
	if err := r.db.Exec(query, nama_toko, url_foto, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *TokoRepositoryImpl) GetTokoByUserId(ctx context.Context, id uint) (entity.Toko, error) {
	var toko entity.Toko
	query := "SELECT * FROM tokos WHERE id_user = ? LIMIT 1"
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&toko).Error; err != nil {
		return toko, err
	}
	return toko, nil
}