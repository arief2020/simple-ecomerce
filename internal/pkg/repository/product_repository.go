package repository

import (
	"context"
	// "fmt"

	// "fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, data entity.Product) (entity.Product, error)
	CreatePhotoProduct(ctx context.Context, data entity.FotoProduct) (entity.FotoProduct, error)
	CreateLogProduct(ctx context.Context, data entity.LogProduct) (entity.LogProduct, error)

	GetAllProduct(ctx context.Context, params dto.AllProductFilter) (res []entity.Product, err error)

	GetProductByID(ctx context.Context, id uint) (res entity.Product, err error)

	UpdateProductByID(ctx context.Context, id uint, data entity.Product) (string, error)
	DeleteProductByID(ctx context.Context, id uint) (string, error)

	GetMyProductById(ctx context.Context, userId uint, tokoId uint, productId uint) (res entity.Product, err error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) CreateProduct(ctx context.Context, data entity.Product) (entity.Product, error) {
	product := entity.Product{
		NamaProduk:    data.NamaProduk,
		Deskripsi:     data.Deskripsi,
		Slug:          data.Slug,
		TokoID:        data.TokoID,
		CategoryID:    data.CategoryID,
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
	}

	if err := r.db.WithContext(ctx).Create(&product).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Product")
		return product, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) CreatePhotoProduct(ctx context.Context, data entity.FotoProduct) (entity.FotoProduct, error) {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Foto Product")
		return data, err
	}
	return data, nil
}

func (r *ProductRepositoryImpl) GetAllProduct(ctx context.Context, params dto.AllProductFilter) (res []entity.Product, err error) {

	// fmt.Printf("params: %+v\n", params)

	query := r.db.Debug().WithContext(ctx).
		Preload("FotoProduct").
		Preload("Toko").
		Preload("Category").
		Model(&entity.Product{})

	// Filter berdasarkan nama produk
	if params.NamaProduk != "" {
		query = query.Where("nama_produk LIKE ?", "%"+params.NamaProduk+"%")
	}

	// Filter berdasarkan kategori
	if params.CategoryID != 0 {
		query = query.Where("id_category = ?", params.CategoryID)
	}

	// Filter berdasarkan toko
	if params.TokoID != 0 {
		query = query.Where("id_toko = ?", params.TokoID)
	}

	if params.MinHarga != 0 {
		query = query.Where("harga_konsumen >= ?", params.MinHarga)
	}
	if params.MaxHarga != 0 {
		query = query.Where("harga_konsumen <= ?", params.MaxHarga)
	}

	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}
	if params.Page > 0 {
		query = query.Offset(params.Page)
	}

	if err := query.Find(&res).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Product")
		return nil, err
	}
	return res, nil
}
func (r *ProductRepositoryImpl) GetProductByID(ctx context.Context, id uint) (res entity.Product, err error) {
	// var product entity.Product
	if err := r.db.Debug().WithContext(ctx).
		Preload("FotoProduct").
		Preload("Toko").
		Preload("Category").
		Where("id = ?", id).
		First(&res).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Product By ID")
		return res, err
	}
	return res, nil
}

func (r *ProductRepositoryImpl) UpdateProductByID(ctx context.Context, id uint, data entity.Product) (string, error) {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Updates(&data).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update Product")
		return "", err
	}
	return "Product updated successfully", nil
}

func (r *ProductRepositoryImpl) DeleteProductByID(ctx context.Context, id uint) (string, error) {
	var product entity.Product
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&product).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete Product")
		return "", err
	}
	return "Product deleted successfully", nil
}

func (r *ProductRepositoryImpl) CreateLogProduct(ctx context.Context, data entity.LogProduct) (entity.LogProduct, error) {
	if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Log Product")
		return data, err
	}
	return data, nil
}

func (r *ProductRepositoryImpl) GetMyProductById(ctx context.Context, userId uint, tokoId uint, productId uint) (res entity.Product, err error) {
	if err := r.db.WithContext(ctx).Where("id = ?", productId).Where("id_toko = ?", tokoId).First(&res).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Product By ID")
		return res, err
	}
	return res, nil
}
