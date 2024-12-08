package controller

import (
	"fmt"
	"strconv"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TokoController interface {
	GetMyToko(ctx *fiber.Ctx) error
	GetTokoByID(ctx *fiber.Ctx) error
	GetAllToko(ctx *fiber.Ctx) error
	UpdateMyToko(ctx *fiber.Ctx) error
}

type TokoControllerImpl struct {
	tokoUsc usecase.TokoUseCase
}

func NewTokoController(tokoUsc usecase.TokoUseCase) TokoController {
	return &TokoControllerImpl{tokoUsc: tokoUsc}
}


func (c *TokoControllerImpl) GetMyToko(ctx *fiber.Ctx) error {

	userId := ctx.Locals("userid").(string)
    fmt.Println(userId)
	
	id, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
        helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	toko, errorMyToko := c.tokoUsc.GetMyToko(ctx.Context(), uint(id))
	if errorMyToko != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Toko")
        return helper.BuildResponse(ctx, false, "Failed to GET data", errorMyToko.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, toko, fiber.StatusOK)
}

func (c *TokoControllerImpl) GetTokoByID(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id_toko")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Toko ID")
        return helper.BuildResponse(ctx, false, "Invalid Toko ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	toko, errUsc := c.tokoUsc.GetTokoByID(ctx.Context(), uint(id))
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Toko By ID")
        return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, toko, fiber.StatusOK)
}

func (c *TokoControllerImpl) GetAllToko(ctx *fiber.Ctx) error {

    filter := new(dto.TokoFilter)
    if err := ctx.QueryParser(filter); err != nil {
        helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Query")
        return helper.BuildResponse(ctx, false, "Failed to GET data", "Failed to parse request query", nil, fiber.StatusBadRequest)
    }

    toko, err := c.tokoUsc.GetAllToko(ctx.Context(), dto.TokoFilter{
        Nama: filter.Nama,
        Limit:    filter.Limit,
        Page:     filter.Page,
    })
	if err != nil {
        helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Toko")
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, toko, fiber.StatusOK)
}

func (c *TokoControllerImpl) UpdateMyToko(ctx *fiber.Ctx) error {
    idToko, err := ctx.ParamsInt("id_toko")
    if err != nil {
        helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Params Toko ID")
        return helper.BuildResponse(ctx, false, "Failed to GET data", err.Error(), nil, fiber.StatusBadRequest)
    }

    userId := ctx.Locals("userid").(string)
    userIdUint, err := strconv.ParseUint(userId, 10, 32)
    if err != nil {
        helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
        return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
    }

    input := &dto.UpdateProfileTokoReq{
        NamaToko: ctx.FormValue("nama_toko"),
    }

    file, _ := ctx.FormFile("photo")

    res, errRes := c.tokoUsc.UpdateMyToko(ctx.Context(), uint(userIdUint), uint(idToko), input, file)
    if errRes != nil {
        helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update My Toko")
        return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errRes.Err.Error(), nil, errRes.Code)
    }

    return helper.BuildResponse(ctx, true, "Succeed to UPDATE data", nil, res, fiber.StatusOK)
}