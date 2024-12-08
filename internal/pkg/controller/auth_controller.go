package controller

import (
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authUsc usecase.AuthsUseCase
}

func NewAuthController(authUsc usecase.AuthsUseCase) AuthController {
	return &AuthControllerImpl{
		authUsc: authUsc,
	}
}

func (uc *AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	c := ctx.Context()

	data := new(dto.Login)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, fmt.Sprint("Error parse request body : ", err.Error()))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := uc.authUsc.Login(c, *data)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, fmt.Sprint("Error Login user : ", err.Err.Error()))
		return helper.BuildResponse(ctx, false, "Failed to Login user", err.Err.Error(), nil, fiber.StatusUnauthorized)
	}

	
	return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, res, fiber.StatusCreated)
}

func (uc *AuthControllerImpl) Register(ctx *fiber.Ctx) error {

    c := ctx.Context()

    data := new(dto.CreateUser)
    if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, fmt.Sprint("Error parse request body : ", err.Error()))
        return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
    }

    _, err := uc.authUsc.CreateUsers(c, *data)
    if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, fmt.Sprint("Error Register user : ", err.Err.Error()))
        return helper.BuildResponse(ctx, false, "Failed to POST data", err.Err.Error(), nil, fiber.StatusBadRequest)
    }

    return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, "Register Succeed", fiber.StatusCreated)

}
