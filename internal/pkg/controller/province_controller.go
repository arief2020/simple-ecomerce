package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// type ProvinceCityController interface {
// 	ListProvincies(ctx *fiber.Ctx) error
// }

// type ProvinceCityControllerImpl struct {
// 	provCityUsc authUsc.UsersUseCase
// }

type ProvinceCityController interface {
	GetAllProvinces(ctx *fiber.Ctx) error
	GetAllCitiesByProvinceID(ctx *fiber.Ctx) error
	GetProvinceByID(ctx *fiber.Ctx) error
	GetCityByID(ctx *fiber.Ctx) error
}

type ProvinceCityControllerImpl struct {
	provincecityusecase usecase.ProvinceCityUseCase
}

func NewProvinceCityController(provincecityusecase usecase.ProvinceCityUseCase) ProvinceCityController {
	return &ProvinceCityControllerImpl{
		provincecityusecase: provincecityusecase,
	}
}


// @Summary Get All Provinces
// @Description Endpoint for get all provinces
// @Tags Province City
// @Accept json	
// @Produce json
// @Param filter query dto.ProvinceFilter true "Province Filter"
// @Success 200 {object} helper.Response{data=[]dto.ProvinceResp} "Succeed to get all provinces"
// @Router /provcity/listprovincies [get]
func (uc *ProvinceCityControllerImpl) GetAllProvinces(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := &dto.ProvinceFilter{}
	err := ctx.QueryParser(filter)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, err.Error())
		return helper.BuildResponse(ctx, false, "Failed to parse query params", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, customErr := uc.provincecityusecase.GetAllProvinces(c, filter)
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.BuildResponse(ctx, false, "Failed to GET data", customErr.Err, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

// @Summary Get All Cities By Province ID
// @Description Endpoint for get all cities by province id
// @Tags Province City
// @Accept json	
// @Produce json
// @Param prov_id path int true "Province ID"
// @Success 200 {object} helper.Response{data=[]dto.CityResp} "Succeed to get all cities by province id"
// @Router /provcity/listcities/{prov_id} [get]
func (uc *ProvinceCityControllerImpl) GetAllCitiesByProvinceID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	provinceid := ctx.Params("prov_id")
	if provinceid == "" {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Bad request")
		return helper.BuildResponse(ctx, false, "Failed to GET data", "Bad request", nil, fiber.StatusBadRequest)
	}

	res, customErr := uc.provincecityusecase.GetAllCitiesByProvinceID(c, provinceid)
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.BuildResponse(ctx, false, "Failed to GET data", customErr.Err, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

// @Summary Get Province By ID
// @Description Endpoint for get province by id
// @Tags Province City
// @Accept json	
// @Produce json
// @Param prov_id path int true "Province ID"
// @Success 200 {object} helper.Response{data=dto.ProvinceResp} "Succeed to get province by id"
// @Router /provcity/detailprovince/{prov_id} [get]
func (uc *ProvinceCityControllerImpl) GetProvinceByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	provinceid := ctx.Params("prov_id")
	if provinceid == "" {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Bad request")
		return helper.BuildResponse(ctx, false, "Failed to GET data", "Bad request", nil, fiber.StatusBadRequest)
	}

	res, customErr := uc.provincecityusecase.GetProvinceByID(c, provinceid)
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.BuildResponse(ctx, false, "Failed to GET data", customErr.Err, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

// @Summary Get City By ID
// @Description Endpoint for get city by id
// @Tags Province City
// @Accept json	
// @Produce json
// @Param city_id path int true "City ID"
// @Success 200 {object} helper.Response{data=dto.CityResp} "Succeed to get city by id"
// @Router /provcity/detailcity/{city_id} [get]
func (uc *ProvinceCityControllerImpl) GetCityByID(ctx *fiber.Ctx) error {
	c := ctx.Context()

	cityId := ctx.Params("city_id")
	if cityId == "" {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Bad request")
		return helper.BuildResponse(ctx, false, "Failed to GET data", "Bad request", nil, fiber.StatusBadRequest)
	}

	res, customErr := uc.provincecityusecase.GetCityByID(c, cityId)
	if customErr != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, customErr.Err.Error())
		return helper.BuildResponse(ctx, false, "Failed to GET data", customErr.Err, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}
