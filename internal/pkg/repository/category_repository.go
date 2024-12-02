package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategory(ctx context.Context) ([]entity.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (entity.Category, error)
	CreateCategory(ctx context.Context, data entity.Category) (entity.Category, error)
	UpdateCategoryByID(ctx context.Context, id uint, data entity.Category) (entity.Category, error)
	DeleteCategoryByID(ctx context.Context, id uint) (string, error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (r *CategoryRepositoryImpl) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	var category []entity.Category
	if err := r.db.WithContext(ctx).Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) GetCategoryByID(ctx context.Context, id uint) (entity.Category, error) {
	var category entity.Category
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (r *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data entity.Category) (entity.Category, error) {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *CategoryRepositoryImpl) UpdateCategoryByID(ctx context.Context, id uint, data entity.Category) (entity.Category, error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Updates(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *CategoryRepositoryImpl) DeleteCategoryByID(ctx context.Context, id uint) (string, error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&entity.Category{}).Error; err != nil {
		return "", err
	}
	return "success", nil
}
