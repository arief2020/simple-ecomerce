package usecase

import (
	"context"
	"fmt"
	"time"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TrxUseCase interface {
	CreateTrx(ctx context.Context, trxDto dto.TransactionRequest, userId uint ) (int, *helper.ErrorStruct)
	GetAllTransctionByUserID(ctx context.Context, userId uint) ([]entity.Trx, *helper.ErrorStruct)
	GetTransactionByID(ctx context.Context, trxId uint, userId uint) (entity.Trx, *helper.ErrorStruct)
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

	// dataLogProducts := []entity.LogProduct{}
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
			TokoID: resRepoProduct.TokoID,
			CategoryID: resRepoProduct.CategoryID,
		}

		// ubah perintah menjadi create log product
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
			// TrxID: trxDto.ID, // pastikan transaksi id sudah terbuat
			LogProductId: resRepoLogProduct.ID, //ubah ke hasil create log product id
			TokoID: resRepoProduct.TokoID,
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

func (t *TrxUseCaseImpl) GetAllTransctionByUserID(ctx context.Context, userId uint) ([]entity.Trx, *helper.ErrorStruct) {
	resRepoTrx, errRepoTrx := t.trxRepo.GetAllTrxByUserID(ctx, userId)
	if errRepoTrx != nil {
		return nil, &helper.ErrorStruct{
			Err: errRepoTrx,
			Code: fiber.StatusBadRequest,
		}
	}
	return resRepoTrx, nil
}

func (t *TrxUseCaseImpl) GetTransactionByID(ctx context.Context, trxId uint, userId uint) (entity.Trx, *helper.ErrorStruct) {
	resRepoTrx, errRepoTrx := t.trxRepo.GetTrxByID(ctx, trxId, userId)
	if errRepoTrx != nil {
		return entity.Trx{}, &helper.ErrorStruct{
			Err: errRepoTrx,
			Code: fiber.StatusBadRequest,
		}
	}
	return resRepoTrx, nil
}