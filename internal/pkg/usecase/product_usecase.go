package usecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product dto.ProductCreateReq, photos []*multipart.FileHeader, userId uint) (int, *helper.ErrorStruct)
	GetAllProduct(ctx context.Context, params dto.AllProductFilter) (*dto.AllProductResp, *helper.ErrorStruct)
	GetProductByID(ctx context.Context, id uint) (*dto.ProductResp, *helper.ErrorStruct)
	UpdateProductByID(ctx context.Context, id uint, product dto.ProductUpdateReq) (string, *helper.ErrorStruct)
	DeleteProductByID(ctx context.Context, id uint) (string, *helper.ErrorStruct)
}

type ProductUseCaseImpl struct {
	productRepo  repository.ProductRepository
	tokoRepo     repository.TokoRepository
	userRepo     repository.UsersRepository
	categoryRepo repository.CategoryRepository
}

func NewProductUseCase(productRepo repository.ProductRepository, tokoRepo repository.TokoRepository, userRepo repository.UsersRepository, categoryRepo repository.CategoryRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		productRepo:  productRepo,
		tokoRepo:     tokoRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

func (u *ProductUseCaseImpl) CreateProduct(ctx context.Context, productReq dto.ProductCreateReq, photos []*multipart.FileHeader, userId uint) (int, *helper.ErrorStruct) {
	_, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get User By ID")
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
	}

	_, errCategoryRepo := u.categoryRepo.GetCategoryByID(ctx, productReq.CategoryID)
	if errCategoryRepo != nil {
		if errors.Is(errCategoryRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Category Not Found")
			return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New("category not found")}

		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Category By ID")
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errCategoryRepo.Error())}

	}

	dataToko, err := u.tokoRepo.GetTokoByUserId(ctx, userId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By User ID")
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
	}

	slug := utils.CreateSlug(productReq.NamaProduk)
	dataReq := entity.Product{
		NamaProduk:    productReq.NamaProduk,
		CategoryID:    productReq.CategoryID,
		HargaReseller: productReq.HargaReseller,
		HargaKonsumen: productReq.HargaKonsumen,
		Stok:          productReq.Stok,
		Deskripsi:     productReq.Deskripsi,
		TokoID:        dataToko.ID,
		Slug:          slug,
	}

	pathUploadedPhotos := []string{}
	for _, photo := range photos {
		uploadedPhoto, err := helper.UploadFile(photo, "uploads")
		if err != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Upload File")
			return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
		}
		pathUploadedPhotos = append(pathUploadedPhotos, uploadedPhoto)
	}

	resCreateRepo, errRepo := u.productRepo.CreateProduct(ctx, dataReq)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Product")
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelInfo, fmt.Sprintf("Product ID : %v", resCreateRepo.ID))

	for _, photo := range pathUploadedPhotos {
		data := entity.FotoProduct{
			UrlFoto:   photo,
			ProductID: resCreateRepo.ID,
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelInfo, fmt.Sprintf("Photo Product : %v", data))

		_, errRepoPhoto := u.productRepo.CreatePhotoProduct(ctx, data)
		if errRepoPhoto != nil {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Photo Product")
			return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepoPhoto.Error())}
		}
	}

	return int(resCreateRepo.ID), nil
}

func (u *ProductUseCaseImpl) GetAllProduct(ctx context.Context, params dto.AllProductFilter) (*dto.AllProductResp, *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}

	resRepo, errRepo := u.productRepo.GetAllProduct(ctx, dto.AllProductFilter{
		Limit:      params.Limit,
		Page:       params.Page,
		NamaProduk: params.NamaProduk,
		CategoryID: params.CategoryID,
		TokoID:     params.TokoID,
		MaxHarga:   params.MaxHarga,
		MinHarga:   params.MinHarga,
	})
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Product")
		return nil, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	data := []dto.ProductResp{}
	for i := range resRepo {
		data = append(data, mapToProductResp(resRepo[i]))
	}

	resp := &dto.AllProductResp{
		Data:  data,
		Page:  params.Page/params.Limit + 1,
		Limit: params.Limit,
	}

	return resp, nil
}

func (u *ProductUseCaseImpl) GetProductByID(ctx context.Context, id uint) (*dto.ProductResp, *helper.ErrorStruct) {
	resRepo, errRepo := u.productRepo.GetProductByID(ctx, id)
	if errRepo != nil {
		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Product")
			return nil, &helper.ErrorStruct{Code: fiber.StatusNotFound, Err: errors.New("no data product")}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Product By Id")
		return nil, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelInfo, fmt.Sprintf("Product ID : %v", resRepo.ID))

	dataResp := mapToProductResp(resRepo)

	data := dto.ProductResp{
		ID:            dataResp.ID,
		NamaProduk:    dataResp.NamaProduk,
		Slug:          dataResp.Slug,
		HargaReseller: dataResp.HargaReseller,
		HargaKonsumen: dataResp.HargaKonsumen,
		Stok:          dataResp.Stok,
		Deskripsi:     dataResp.Deskripsi,
		Toko:          dataResp.Toko,
		Category:      dataResp.Category,
		Photos:        dataResp.Photos,
	}

	return &data, nil
}

func mapToProductResp(product entity.Product) dto.ProductResp {
	fmt.Println(product)
	photos := []dto.PhotoProductResp{}
	for _, photo := range product.FotoProduct {
		photos = append(photos, dto.PhotoProductResp{
			Id:        photo.ID,
			ProductID: photo.ProductID,
			Url:       photo.UrlFoto,
		})
	}

	return dto.ProductResp{
		ID:            product.ID,
		NamaProduk:    product.NamaProduk,
		Slug:          product.Slug,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		Toko: dto.TokoResp{
			ID:       product.Toko.ID,
			NamaToko: product.Toko.NamaToko,
			UrlFoto:  product.Toko.UrlFoto,
		},
		Category: dto.CategoryResp{
			ID:           product.Category.ID,
			NamaCategory: product.Category.NamaCategory,
		},
		Photos: photos,
	}
}

func (u *ProductUseCaseImpl) UpdateProductByID(ctx context.Context, id uint, productReq dto.ProductUpdateReq) (string, *helper.ErrorStruct) {
	// _, errRepo := u.productRepo.GetProductByID(ctx, id)
	// if errRepo != nil {
	// 	helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Product By Id")
	// 	return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	// }

	userId := ctx.Value("userid").(string)
	uintUser := utils.StringToUint(userId)

	resToko, errToko := u.tokoRepo.GetTokoByUserId(ctx, uint(uintUser))
	if errToko != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By Id")
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errToko.Error())}
	}

	_, errRepo := u.productRepo.GetMyProductById(ctx, uint(uintUser), resToko.ID, id)
	if errRepo != nil {

		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Product")
			return "", &helper.ErrorStruct{Code: fiber.StatusNotFound, Err: errors.New("no data product")}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Product By Id")
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	slug := utils.CreateSlug(productReq.NamaProduk)

	data := entity.Product{
		NamaProduk:    productReq.NamaProduk,
		CategoryID:    productReq.CategoryID,
		Slug:          slug,
		HargaReseller: productReq.HargaReseller,
		HargaKonsumen: productReq.HargaKonsumen,
		Stok:          productReq.Stok,
		Deskripsi:     productReq.Deskripsi,
	}

	resUpdateRepo, errRepo := u.productRepo.UpdateProductByID(ctx, id, data)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update Product By Id")
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	return resUpdateRepo, nil
}

func (u *ProductUseCaseImpl) DeleteProductByID(ctx context.Context, productId uint) (string, *helper.ErrorStruct) {

	// _, errRepo := u.productRepo.GetProductByID(ctx, productId)
	// if errRepo != nil {
	// 	helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Product By Id")
	// 	return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	// }

	userId := ctx.Value("userid").(string)
	uintUser := utils.StringToUint(userId)

	resToko, errToko := u.tokoRepo.GetTokoByUserId(ctx, uint(uintUser))
	if errToko != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By Id")
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errToko.Error())}
	}

	_, errRepo := u.productRepo.GetMyProductById(ctx, uint(uintUser), resToko.ID, productId)
	if errRepo != nil {

		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Product")
			return "", &helper.ErrorStruct{Code: fiber.StatusNotFound, Err: errors.New("no data product")}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Product By Id")
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	res, err := u.productRepo.DeleteProductByID(ctx, productId)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete Product By Id")
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
	}

	return res, nil
}
