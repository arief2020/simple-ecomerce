package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, data entity.Product) (entity.Product, error)
    CreatePhotoProduct(ctx context.Context, data entity.FotoProduct) (entity.FotoProduct, error)
    GetAllProduct(ctx context.Context) ([]entity.Product, error)
    GetProductByID(ctx context.Context, id uint) (entity.Product, error)
    UpdateProductByID(ctx context.Context, id uint, data entity.Product) (string, error)
    DeleteProductByID(ctx context.Context, id uint) (string, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

// func (r *ProductRepositoryImpl) CreateProduct(ctx context.Context, data entity.Product, photo []entity.FotoProduct) (entity.Product, error) {

// 	product := entity.Product{
// 		NamaProduk: data.NamaProduk,
// 		Deskripsi:  data.Deskripsi,
// 		Slug:       data.Slug,
// 		TokoID:     data.TokoID,
// 		CategoryID: data.CategoryID,
// 		HargaReseller: data.HargaReseller,
// 		HargaKonsumen: data.HargaKonsumen,
// 		Stok: data.Stok,
// 		FotoProduct: photo,
// 	}


// 	if err := r.db.WithContext(ctx).Create(&product).Error; err != nil {
// 		return data, err
// 	}
// 	return data, nil
// }


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
        // FotoProduct:   photo, // Ini akan di-embed di dalam Product
    }

//     for i := range product.FotoProduct {
//     product.FotoProduct[i].ProductID = product.ID // Pastikan ID produk diatur dengan benar
// }


    if err := r.db.WithContext(ctx).Create(&product).Error; err != nil {
        return product, err
    }

    //  // Set ProductID untuk setiap FotoProduct
    // for i := range photos {
    //     photos[i].ProductID = product.ID
    // }

    // // Simpan FotoProduct
    // if err := r.db.WithContext(ctx).Create(&photos).Error; err != nil {
    //     return product, err
    // }

    return product, nil
}


func (r *ProductRepositoryImpl) CreatePhotoProduct(ctx context.Context, data entity.FotoProduct) (entity.FotoProduct, error) {
    if err := r.db.WithContext(ctx).Create(&data).Error; err != nil {
        return data, err
    }
    return data, nil
}

func (r *ProductRepositoryImpl) GetAllProduct(ctx context.Context) ([]entity.Product, error) {
    var products []entity.Product
    if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (r *ProductRepositoryImpl) GetProductByID(ctx context.Context, id uint) (entity.Product, error) {
    var product entity.Product
    if err := r.db.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
        return product, err
    }
    return product, nil
}

func (r *ProductRepositoryImpl) UpdateProductByID(ctx context.Context, id uint, data entity.Product) (string, error) {
    if err := r.db.WithContext(ctx).Where("id = ?", id).Updates(&data).Error; err != nil {
        return "", err
    }
    return "Product updated successfully", nil
}

func (r *ProductRepositoryImpl) DeleteProductByID(ctx context.Context, id uint) (string, error) {
    var product entity.Product
    if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&product).Error; err != nil {
        return "", err
    }
    return "Product deleted successfully", nil
}