package usecase

import (
	"context"
	"errors"
	"mime/multipart"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product dto.ProductCreateReq, photos []*multipart.FileHeader, userId uint) (int, *helper.ErrorStruct)
	GetAllProduct(ctx context.Context) ([]dto.ProductResp, *helper.ErrorStruct)
	GetProductByID(ctx context.Context, id uint) (*dto.ProductResp, *helper.ErrorStruct)
	UpdateProductByID(ctx context.Context, id uint, product dto.ProductUpdateReq) (string, *helper.ErrorStruct)
	DeleteProductByID(ctx context.Context, id uint) (string, *helper.ErrorStruct)
}

type ProductUseCaseImpl struct {
	productRepo repository.ProductRepository
	tokoRepo repository.TokoRepository
	userRepo repository.UsersRepository
}

func NewProductUseCase(productRepo repository.ProductRepository, tokoRepo repository.TokoRepository, userRepo repository.UsersRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		productRepo: productRepo,
		tokoRepo: tokoRepo,
		userRepo: userRepo,
	}
}

func (u *ProductUseCaseImpl) CreateProduct(ctx context.Context, productReq dto.ProductCreateReq,photos []*multipart.FileHeader, userId uint) (int, *helper.ErrorStruct) {
	_, err := u.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
	}

	

	dataToko, err := u.tokoRepo.GetTokoByUserId(ctx, userId)
	if err != nil {
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
	}

	slug := utils.CreateSlug(productReq.NamaProduk)
	dataReq := entity.Product{
		NamaProduk: productReq.NamaProduk,
		CategoryID:  productReq.CategoryID,
		HargaReseller: productReq.HargaReseller,
		HargaKonsumen: productReq.HargaKonsumen,
		Stok: productReq.Stok,
		Deskripsi: productReq.Deskripsi,
		TokoID: dataToko.ID,
		Slug: slug,
	}

	pathUploadedPhotos := []string{}
	for _, photo := range photos {
		uploadedPhoto, err := helper.UploadFile(photo, "uploads")
		if err != nil {
			return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
		}
		pathUploadedPhotos = append(pathUploadedPhotos, uploadedPhoto)
	}

	// dataPhoto := []entity.FotoProduct{}
	// for _, photo := range pathUploadedPhotos {
	// 	dataPhoto = append(dataPhoto, entity.FotoProduct{
	// 		UrlFoto: photo,
	// 	})
	// }

	resCreateRepo, errRepo := u.productRepo.CreateProduct(ctx, dataReq)
	if errRepo != nil {
		return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	for _, photo := range pathUploadedPhotos {
		data := entity.FotoProduct{
			UrlFoto: photo,
			ProductID: resCreateRepo.ID,
		}
		_, errRepoPhoto := u.productRepo.CreatePhotoProduct(ctx, data)
		if errRepoPhoto != nil {
			return 0, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepoPhoto.Error())}
		}
	}


	return int(resCreateRepo.ID), nil
}

func (u *ProductUseCaseImpl) GetAllProduct(ctx context.Context) ([]dto.ProductResp, *helper.ErrorStruct) {
	resRepo, errRepo := u.productRepo.GetAllProduct(ctx)
	if errRepo != nil {
		return nil, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	data := []dto.ProductResp{}
	for _, product := range resRepo {
		data = append(data, dto.ProductResp{
			ID:             product.ID,
			NamaProduk:     product.NamaProduk,
			Slug:           product.Slug,
			HargaReseller:  product.HargaReseller,
			HargaKonsumen:  product.HargaKonsumen,
			Stok:           product.Stok,
			Deskripsi:      product.Deskripsi,
		})
	}

	return data, nil
}

func (u *ProductUseCaseImpl) GetProductByID(ctx context.Context, id uint) (*dto.ProductResp, *helper.ErrorStruct) {
	resRepo, errRepo := u.productRepo.GetProductByID(ctx, id)
	if errRepo != nil {
		return nil, &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	data := dto.ProductResp{
		ID:             resRepo.ID,
		NamaProduk:     resRepo.NamaProduk,
		Slug:           resRepo.Slug,
		HargaReseller:  resRepo.HargaReseller,
		HargaKonsumen:  resRepo.HargaKonsumen,
		Stok:           resRepo.Stok,
		Deskripsi:      resRepo.Deskripsi,
	}

	return &data, nil
}

func (u *ProductUseCaseImpl) UpdateProductByID(ctx context.Context, id uint, productReq dto.ProductUpdateReq) (string, *helper.ErrorStruct) {
	_, errRepo := u.productRepo.GetProductByID(ctx, id)
	if errRepo != nil {
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	slug := utils.CreateSlug(productReq.NamaProduk)

	data := entity.Product{
		NamaProduk: productReq.NamaProduk,
		CategoryID: productReq.CategoryID,
		Slug: slug,
		HargaReseller: productReq.HargaReseller,
		HargaKonsumen: productReq.HargaKonsumen,
		Stok: productReq.Stok,
		Deskripsi: productReq.Deskripsi,
	}

	resUpdateRepo, errRepo := u.productRepo.UpdateProductByID(ctx, id, data)
	if errRepo != nil {
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	return resUpdateRepo, nil
}

func (u *ProductUseCaseImpl) DeleteProductByID(ctx context.Context, productId uint) (string, *helper.ErrorStruct) {
	_, errRepo := u.productRepo.GetProductByID(ctx, productId)
	if errRepo != nil {
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(errRepo.Error())}
	}

	res, err := u.productRepo.DeleteProductByID(ctx, productId)
	if err != nil {
		return "", &helper.ErrorStruct{Code: fiber.StatusBadRequest, Err: errors.New(err.Error())}
	}

	return res, nil
}