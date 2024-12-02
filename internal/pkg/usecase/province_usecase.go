package usecase

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProvinceCityUseCase interface {
	// ListProvincies(ctx context.Context, params dto.ListProvResp) (res dto.ListProvResp, err *helper.ErrorStruct)
	GetAllProvinces(ctx context.Context, filter *dto.ProvinceFilter) (res []*dto.ProvinceResp, err *helper.ErrorStruct)
	GetAllCitiesByProvinceID(ctx context.Context, provinceid string) (res []*dto.CityResp, err *helper.ErrorStruct)
	GetProvinceByID(ctx context.Context, provinceid string) (res *dto.ProvinceResp, err *helper.ErrorStruct)
	GetCityByID(ctx context.Context, cityid string) (res *dto.CityResp, err *helper.ErrorStruct)
}

// type provCityUseCaseImpl struct {
// 	provinceCityRepository userrepository.UsersRepository
// }

type ProvinceCityUseCaseImpl struct {
	provinceCityRepository repository.ProvinceCityRepository
}

func NewProvinceCityUseCase(provinceCityRepository repository.ProvinceCityRepository) ProvinceCityUseCase {
	return &ProvinceCityUseCaseImpl{
		provinceCityRepository: provinceCityRepository,
	}
}

func (alc *ProvinceCityUseCaseImpl) GetAllProvinces(ctx context.Context, filter *dto.ProvinceFilter) (res []*dto.ProvinceResp, err *helper.ErrorStruct) {
	if filter.Limit < 1 {
		filter.Limit = 10
	}

	if filter.Page < 1 {
		filter.Page = 1
	}

	res, errRepo := alc.provinceCityRepository.GetAllProvinces(ctx, filter.Limit, (filter.Page-1)*filter.Limit, filter.Search)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}

func (alc *ProvinceCityUseCaseImpl) GetAllCitiesByProvinceID(ctx context.Context, provinceid string) (res []*dto.CityResp, err *helper.ErrorStruct) {
	res, errRepo := alc.provinceCityRepository.GetAllCitiesByProvinceID(ctx, provinceid)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return res, nil
}

func (alc *ProvinceCityUseCaseImpl) GetProvinceByID(ctx context.Context, provinceid string) (res *dto.ProvinceResp, err *helper.ErrorStruct) {
	res, errRepo := alc.provinceCityRepository.GetProvinceByID(ctx, provinceid)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return res, nil
}

func (alc *ProvinceCityUseCaseImpl) GetCityByID(ctx context.Context, cityid string) (res *dto.CityResp, err *helper.ErrorStruct) {
	res, errRepo := alc.provinceCityRepository.GetCityByID(ctx, cityid)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return res, nil
}


