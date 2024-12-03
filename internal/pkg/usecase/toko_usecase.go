package usecase

import (
	"context"
	"errors"
	"mime/multipart"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

type TokoUseCase interface {
	GetMyToko(ctx context.Context, id uint) ( *dto.MyTokoResp, error)
	UpdateMyToko(ctx context.Context, userId uint, idToko uint, params *dto.UpdateProfileTokoReq, file *multipart.FileHeader) (string, *helper.ErrorStruct)

	GetAllToko(ctx context.Context, params dto.TokoFilter) (*dto.AllTokoResp, *helper.ErrorStruct)
	GetTokoByID(ctx context.Context, id uint) (*dto.MyTokoResp, *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	tokoRepository repository.TokoRepository
}

func NewTokoUseCase(tokoRepository repository.TokoRepository) TokoUseCase {
	return &TokoUseCaseImpl{
		tokoRepository: tokoRepository,
	}
}

func (t *TokoUseCaseImpl) GetMyToko(ctx context.Context, id uint) (*dto.MyTokoResp, error) {
	toko, err := t.tokoRepository.GetTokoById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.MyTokoResp{
		ID:        toko.ID,
		NamaToko:  toko.NamaToko,
		UrlFoto:   toko.UrlFoto,
		UserID:    toko.UserID,
	}, nil
}


func (t *TokoUseCaseImpl) GetTokoByID(ctx context.Context, id uint) (*dto.MyTokoResp, *helper.ErrorStruct) {
	toko, err := t.tokoRepository.GetTokoById(ctx, id)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:     errors.New("Toko tidak ditemukan"),
		}
	}
	return &dto.MyTokoResp{
		ID:        toko.ID,
		NamaToko:  toko.NamaToko,
		UrlFoto:   toko.UrlFoto,
		UserID:    toko.UserID,
	}, nil
}

func (t *TokoUseCaseImpl) GetAllToko(ctx context.Context, params dto.TokoFilter) (*dto.AllTokoResp, *helper.ErrorStruct) {
	if params.Limit < 1 {
		params.Limit = 10
	}

	if params.Page < 1 {
		params.Page = 0
	} else {
		params.Page = (params.Page - 1) * params.Limit
	}
	
	toko, err := t.tokoRepository.GetAllToko(ctx, dto.TokoFilter{
		Limit: params.Limit,
		Page:  params.Page,
		Nama: params.Nama,
	})
	if err != nil {
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:     errors.New("toko tidak ditemukan"),
		}
	}

	allTokoResp := &dto.AllTokoResp{
		Page:  params.Page/params.Limit + 1,
		Limit: params.Limit,
	}

	var tokoResp []dto.TokoResp



	for _, toko := range toko {
		tokoResp = append(tokoResp, dto.TokoResp{
			ID:        toko.ID,
			NamaToko:  toko.NamaToko,
			UrlFoto:   toko.UrlFoto,
		})
	}
	allTokoResp.Data = tokoResp
	return allTokoResp, nil
}

func (t *TokoUseCaseImpl) UpdateMyToko(ctx context.Context, userId uint, idToko uint, data *dto.UpdateProfileTokoReq, file *multipart.FileHeader) (string, *helper.ErrorStruct) {
	dataToko, err := t.tokoRepository.GetTokoById(ctx, idToko)
	if err != nil {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errors.New(err.Error()),
		}
	}

	if dataToko.UserID != userId {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("unauthorized"),
		}
	}

	// Upload file
	uploadsFolder := "uploads"
	savePath, err := helper.UploadFile(file, uploadsFolder)
	if err != nil {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  err,
		}
	}

	// Update database
	err = t.tokoRepository.UpdateToko(ctx, idToko, data.NamaToko, savePath)
	if err != nil {
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  err,
		}
	}

	return "Update toko succeed", nil
}