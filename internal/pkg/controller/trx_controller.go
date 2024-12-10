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
	var req dto.TransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error parsing request body: "+err.Error())
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	userId := ctx.Locals("userid").(string)

	userIdUint := utils.StringToUint(userId)

	resUsc, errUsc := t.trxUsc.CreateTrx(ctx.Context(), req, uint(userIdUint))

	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Transaction")
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to CREATE data", nil, resUsc, fiber.StatusCreated)
}

func (t *TrxControllerImpl) GetAllTransctionByUserID(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid").(string)

	userIdUint := utils.StringToUint(userId)

	filter := new(dto.AllTransactionReq)
	if err := ctx.QueryParser(filter); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Query")
		return helper.BuildResponse(ctx, false, "Failed to parse query params", err.Error(), nil, fiber.StatusBadRequest)
	}

	resUsc, errUsc := t.trxUsc.GetAllTransaction(ctx.Context(), dto.AllTransactionReq{
		Limit:  filter.Limit,
		Page:   filter.Page,
		Search: filter.Search,
	}, uint(userIdUint))

	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Transaction")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, resUsc, fiber.StatusOK)
}

func (t *TrxControllerImpl) GetTransactionByID(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userid").(string)
	uintUserId := utils.StringToUint(userId)
	id, err := ctx.ParamsInt("id_trx")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Query")
		return helper.BuildResponse(ctx, false, "Failed to GET data 1", err.Error(), nil, fiber.StatusBadRequest)
	}

	resUsc, errUsc := t.trxUsc.GetTransactionByID(ctx.Context(), uint(id), uint(uintUserId))

	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Transaction By ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data 2", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data 3", nil, resUsc, fiber.StatusOK)
}
