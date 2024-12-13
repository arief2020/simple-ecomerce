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

// @Summary Create Transaction
// @Description Endpoint for create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param create-transaction body dto.TransactionRequest true "Create Transaction"
// @Success 201 {object} helper.Response{data=int} "Succeed to create transaction"
// @Router /trx [post]
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

// @Summary Get All Transaction
// @Description Endpoint for get all transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} helper.Response{data=[]dto.AllTransactionResponse} "Succeed to get all transaction"
// @Router /trx [get]
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

// @Summary Get Transaction By ID
// @Description Endpoint for get transaction by id
// @Tags Transaction
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id_trx path int true "Transaction ID"
// @Success 200 {object} helper.Response{data=dto.TransactionResponse} "Succeed to get transaction by id"
// @Router /trx/{id_trx} [get]
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
