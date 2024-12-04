package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TrxController interface {
	CreateTransction(ctx *fiber.Ctx) error
	GetAllTransctionByUserID(ctx *fiber.Ctx) error
	GetTransactionByID(ctx *fiber.Ctx) error
}

type TrxControllerImpl struct {
	trxUsc usecase.TrxUseCase
}

func NewTrxController(trxUsc usecase.TrxUseCase) TrxController {
	return &TrxControllerImpl{trxUsc: trxUsc}
}


func (t *TrxControllerImpl) CreateTransction(ctx *fiber.Ctx) error {
	// Parse JSON body ke DTO
	var req dto.TransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userId := ctx.Locals("userid").(string)

	userIdUint := utils.StringToUint(userId)

	resUsc, errUsc := t.trxUsc.CreateTrx(ctx.Context(), req, uint(userIdUint))

	if errUsc != nil {
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errUsc, nil, fiber.StatusBadRequest)
	}

	// Kembalikan request sebagai respons
	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       resUsc,
	})
}


func (t *TrxControllerImpl) GetAllTransctionByUserID(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid").(string)

	userIdUint := utils.StringToUint(userId)

	filter := new(dto.AllTransactionReq)
	if err := ctx.QueryParser(filter); err != nil {
		return helper.BuildResponse(ctx, false, "Failed to parse query params", err.Error(), nil, fiber.StatusBadRequest)
	}

	resUsc, errUsc := t.trxUsc.GetAllTransaction(ctx.Context(), dto.AllTransactionReq{
		Limit: filter.Limit,
		Page: filter.Page,
		Search: filter.Search,
	}, uint(userIdUint))

	if errUsc != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc, nil, fiber.StatusBadRequest)
	}

	// Kembalikan request sebagai respons
	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, resUsc, fiber.StatusOK)
}


func (t *TrxControllerImpl) GetTransactionByID(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid").(string)
	uintUserId := utils.StringToUint(userId)
	id, err := ctx.ParamsInt("id_trx")
	if err != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data 1", err.Error(), nil, fiber.StatusBadRequest)
	}

	resUsc, errUsc := t.trxUsc.GetTransactionByID(ctx.Context(), uint(id), uint(uintUserId))

	if errUsc != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data 2", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	// Kembalikan request sebagai respons
	return helper.BuildResponse(ctx, true, "Succeed to GET data 3", nil, resUsc, fiber.StatusOK)
}