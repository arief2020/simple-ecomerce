package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"

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

func (r *TokoRepositoryImpl) GetAllToko(ctx context.Context, params dto.TokoFilter) (tokos []entity.Toko, err error) {
	db := r.db

	// filter := map[string][]any{
	// 	"nama_toko like ?": {fmt.Sprint("%" + params.Nama + "%")},
	// }

	if params.Nama != "" {
		db = db.Where("nama_toko like ?", "%"+params.Nama+"%")
	}

	if err := db.Debug().WithContext(ctx).Limit(params.Limit).Offset(params.Page).Find(&tokos).Error; err != nil {
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
	// var toko entity.Toko
	// query := "SELECT * FROM tokos WHERE id_user = ? LIMIT 1"
	// if err := r.db.WithContext(ctx).Raw(query, id).Scan(&toko).Error; err != nil {
	// 	return toko, err
	// }
	var toko entity.Toko
	if err := r.db.WithContext(ctx).Where("id_user = ?", id).First(&toko).Error; err != nil {
		return toko, err
	}
	return toko, nil
}