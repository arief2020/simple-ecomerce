package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TrxUseCase interface {
	CreateTrx(ctx context.Context, trxDto dto.TransactionRequest, userId uint ) (int, *helper.ErrorStruct)
	// GetAllTransctionByUserID(ctx context.Context, userId uint) ([]entity.Trx, *helper.ErrorStruct)
	GetAllTransaction(ctx context.Context, req dto.AllTransactionReq, userId uint) (*dto.AllTransactionResponse, *helper.ErrorStruct)
	GetTransactionByID(ctx context.Context, trxId uint, userId uint) (*dto.TransactionResponse, *helper.ErrorStruct)
}

type TrxUseCaseImpl struct {
	trxRepo repository.TrxRepository
	userRepo repository.UsersRepository
	productRepo repository.ProductRepository
}

func NewTrxUseCase(trxRepo repository.TrxRepository, userRepo repository.UsersRepository, productRepo repository.ProductRepository) TrxUseCase {
	return &TrxUseCaseImpl{
		trxRepo: trxRepo,
		userRepo: userRepo,
		productRepo: productRepo,
	}
}



func (t *TrxUseCaseImpl) CreateTrx(ctx context.Context, trxDto dto.TransactionRequest, userId uint) (int, *helper.ErrorStruct) {
	_, err := t.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return 0, &helper.ErrorStruct{
			Err: err,
			Code: fiber.StatusBadRequest,
		}
	}

	dataDetailsTrx := []entity.DetailTrx{}
	trxTotal := 0
	
	for _, detailTrx := range trxDto.DetailTrx {
		resRepoProduct, errRepoProduct := t.productRepo.GetProductByID(ctx, uint(detailTrx.ProductID))
		if errRepoProduct != nil {
			return 0, &helper.ErrorStruct{
				Err: errRepoProduct,
				Code: fiber.StatusBadRequest,
			}
		}
		intStock := utils.StringToUint(resRepoProduct.Stok)
		if int(intStock) < detailTrx.Kuantitas {
			return 0, &helper.ErrorStruct{
				Err: errRepoProduct,
				Code: fiber.StatusBadRequest,
			}
		}

		dataLogProduct := entity.LogProduct{
			ProductID: resRepoProduct.ID,
			NamaProduk: resRepoProduct.NamaProduk,
			Slug: resRepoProduct.Slug,
			HargaReseller: resRepoProduct.HargaReseller,
			HargaKonsumen: resRepoProduct.HargaKonsumen,
			Stok: resRepoProduct.Stok,
			Deskripsi: resRepoProduct.Deskripsi,
			TokoID: resRepoProduct.Toko.ID,
			CategoryID: resRepoProduct.Category.ID,
		}

		resRepoLogProduct, errRepoLogProduct := t.productRepo.CreateLogProduct(ctx, dataLogProduct)
		if errRepoLogProduct != nil {
			return 0, &helper.ErrorStruct{
				Err: errRepoLogProduct,
				Code: fiber.StatusBadRequest,
			}
		}

		uintHargaKonsumen := utils.StringToUint(resRepoProduct.HargaKonsumen)

		detailTrxTotal := int(uintHargaKonsumen) * detailTrx.Kuantitas

		trxTotal += detailTrxTotal

		dataDetailTrx := &entity.DetailTrx{
			LogProductId: resRepoLogProduct.ID, 
			TokoID: resRepoProduct.Toko.ID,
			Kuantitas: detailTrx.Kuantitas,
			HargaTotal: detailTrxTotal,
		}

		dataDetailsTrx = append(dataDetailsTrx, *dataDetailTrx)

	}

	dataTrx := &entity.Trx{
		UserID:   userId,
		AlamatID: uint(trxDto.AlamatKirim),
		HargaTotal: trxTotal,
		KodeInvoice: fmt.Sprintf("INV-%d", time.Now().UnixNano()),
		MethodBayar: trxDto.MethodBayar,
	}
	

	resRepoTrx, errRepoTrx := t.trxRepo.CreateTrx(ctx, *dataTrx)
	if errRepoTrx != nil {
		return 0, &helper.ErrorStruct{
			Err: errRepoTrx,
			Code: fiber.StatusBadRequest,
		}
	}

	for _, detailTrx := range dataDetailsTrx {
		detailTrx.TrxID = resRepoTrx.ID

		_, errRepoDetailTrx := t.trxRepo.CreateDetailTrx(ctx, detailTrx)
		if errRepoDetailTrx != nil {
			return 0, &helper.ErrorStruct{
				Err: errRepoDetailTrx,
				Code: fiber.StatusBadRequest,
			}
		}
	}
	return int(resRepoTrx.ID), nil
}

// func (t *TrxUseCaseImpl) GetAllTransctionByUserID(ctx context.Context, userId uint) ([]entity.Trx, *helper.ErrorStruct) {
// 	resRepoTrx, errRepoTrx := t.trxRepo.GetAllTrxByUserID(ctx, userId)
// 	if errRepoTrx != nil {
// 		return nil, &helper.ErrorStruct{
// 			Err: errRepoTrx,
// 			Code: fiber.StatusBadRequest,
// 		}
// 	}
// 	return resRepoTrx, nil
// }

// func (t *TrxUseCaseImpl) GetTransactionByID(ctx context.Context, trxId uint, userId uint) (*dto.TransactionResponse, *helper.ErrorStruct) {
// 	resRepoTrx, errRepoTrx := t.trxRepo.GetTrxByID(ctx, trxId, userId)
// 	if errRepoTrx != nil {
// 		return nil, &helper.ErrorStruct{
// 			Err: errors.New("no data trx"),
// 			Code: fiber.StatusBadRequest,
// 		}
// 	}
// 	dataResp := &dto.TransactionResponse{
// 		ID: int(resRepoTrx.ID),
// 		HargaTotal: resRepoTrx.HargaTotal,
// 		KodeInvoice: resRepoTrx.KodeInvoice,
// 		MethodBayar: resRepoTrx.MethodBayar,
// 		AlamatKirim: dto.AlamatResp{
// 			Id: resRepoTrx.AlamatID,
// 			JudulAlamat: resRepoTrx.Alamat.JudulAlamat,
// 			NamaPenerima: resRepoTrx.Alamat.NamaPenerima,
// 			NoTelp: resRepoTrx.Alamat.NoTelp,
// 			DetailAlamat: resRepoTrx.Alamat.DetailAlamat,
// 		},
// 		DetailTrx: nil,
// 	}
// 	return dataResp, nil
// }

func (t *TrxUseCaseImpl) GetTransactionByID(ctx context.Context, trxId uint, userId uint) (*dto.TransactionResponse, *helper.ErrorStruct) {
	// Ambil data transaksi dari repository
	resRepoTrx, errRepoTrx := t.trxRepo.GetTrxByID(ctx, trxId, userId)
	if errRepoTrx != nil {
		if errors.Is(errRepoTrx, gorm.ErrRecordNotFound) {
			return nil, &helper.ErrorStruct{
				Err:  errors.New("transaction not found"),
				Code: fiber.StatusNotFound,
			}
		}
		return nil, &helper.ErrorStruct{
			Err:  errRepoTrx,
			Code: fiber.StatusInternalServerError,
		}
	}

	// Bangun detail transaksi
	var detailTrx []dto.DetailTrx
	for _, detail := range resRepoTrx.DetailTrx {
		detailTrx = append(detailTrx, dto.DetailTrx{
			Product: dto.TransactionProductResp{
				ID:             detail.LogProduct.ID,
				NamaProduk:     detail.LogProduct.NamaProduk,
				Slug:           detail.LogProduct.Slug,
				HargaReseller:  detail.LogProduct.HargaReseller,
				HargaKonsumen:  detail.LogProduct.HargaKonsumen,
				Deskripsi:      detail.LogProduct.Deskripsi,
				Toko:           dto.TokoResp{
					ID:       detail.LogProduct.Toko.ID,
					NamaToko: detail.LogProduct.Toko.NamaToko,
					UrlFoto:  detail.LogProduct.Toko.UrlFoto,
				},
				Category:       dto.CategoryResp{
					ID:           detail.LogProduct.Category.ID,
					NamaCategory: detail.LogProduct.Category.NamaCategory,
				},
				Photos:         mapPhotos(detail.LogProduct.Product.FotoProduct),
			},
			Toko: dto.TokoResp{
				ID:       detail.LogProduct.Toko.ID,
				NamaToko: detail.LogProduct.Toko.NamaToko,
				UrlFoto:  detail.LogProduct.Toko.UrlFoto,
			},
			Kuantitas:  detail.Kuantitas,
			HargaTotal: detail.HargaTotal,
		})
	}

	// Bangun response transaksi
	dataResp := &dto.TransactionResponse{
		ID:          int(resRepoTrx.ID),
		HargaTotal:  resRepoTrx.HargaTotal,
		KodeInvoice: resRepoTrx.KodeInvoice,
		MethodBayar: resRepoTrx.MethodBayar,
		AlamatKirim: dto.AlamatResp{
			Id:           resRepoTrx.AlamatID,
			JudulAlamat:  resRepoTrx.Alamat.JudulAlamat,
			NamaPenerima: resRepoTrx.Alamat.NamaPenerima,
			NoTelp:       resRepoTrx.Alamat.NoTelp,
			DetailAlamat: resRepoTrx.Alamat.DetailAlamat,
		},
		DetailTrx: detailTrx,
	}

	return dataResp, nil
}


func (t *TrxUseCaseImpl) GetAllTransaction(ctx context.Context, req dto.AllTransactionReq, userId uint) (*dto.AllTransactionResponse, *helper.ErrorStruct) {
	// Default pagination
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Page < 1 {
		req.Page = 1
	}

	// Ambil data dari repository
	transactions, _, err := t.trxRepo.GetAllTransaction(ctx, req, userId)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusInternalServerError,
		}
	}

	// Bangun respons transaksi
	var transactionResponses []dto.TransactionResponse
	for _, trx := range transactions {
		// Detail transaksi
		var detailTrx []dto.DetailTrx
		for _, detail := range trx.DetailTrx {
			detailTrx = append(detailTrx, dto.DetailTrx{
				Product: dto.TransactionProductResp{
					ID:             detail.LogProduct.ID,
				NamaProduk:     detail.LogProduct.NamaProduk,
				Slug:           detail.LogProduct.Slug,
				HargaReseller:  detail.LogProduct.HargaReseller,
				HargaKonsumen:  detail.LogProduct.HargaKonsumen,
				Deskripsi:      detail.LogProduct.Deskripsi,
				Toko:           dto.TokoResp{
					ID:       detail.LogProduct.Toko.ID,
					NamaToko: detail.LogProduct.Toko.NamaToko,
					UrlFoto:  detail.LogProduct.Toko.UrlFoto,
				},
				Category:       dto.CategoryResp{
					ID:           detail.LogProduct.Category.ID,
					NamaCategory: detail.LogProduct.Category.NamaCategory,
				},
				Photos:         mapPhotos(detail.LogProduct.Product.FotoProduct),
				},
				Toko: dto.TokoResp{
				ID:       detail.LogProduct.Toko.ID,
				NamaToko: detail.LogProduct.Toko.NamaToko,
				UrlFoto:  detail.LogProduct.Toko.UrlFoto,
			},
				Kuantitas:  detail.Kuantitas,
				HargaTotal: detail.HargaTotal,
			})
		}

		// Respons transaksi
		transactionResponses = append(transactionResponses, dto.TransactionResponse{
			ID:          int(trx.ID),
			HargaTotal:  trx.HargaTotal,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
			AlamatKirim: dto.AlamatResp{
				Id:           trx.AlamatID,
				JudulAlamat:  trx.Alamat.JudulAlamat,
				NamaPenerima: trx.Alamat.NamaPenerima,
				NoTelp:       trx.Alamat.NoTelp,
				DetailAlamat: trx.Alamat.DetailAlamat,
			},
			DetailTrx: detailTrx,
		})
	}

	// Bangun respons utama
	response := &dto.AllTransactionResponse{
		Data:  transactionResponses,
		Page:  req.Page,
		Limit: req.Limit,
	}

	return response, nil
}

// Helper untuk memetakan foto produk
func mapPhotos(photos []entity.FotoProduct) []dto.PhotoProductResp {
	var photoResps []dto.PhotoProductResp
	for _, photo := range photos {
		photoResps = append(photoResps, dto.PhotoProductResp{
			Id:  photo.ID,
			ProductID: photo.ProductID,
			Url: photo.UrlFoto,
		})
	}
	return photoResps
}