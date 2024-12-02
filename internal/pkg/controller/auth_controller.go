package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// @TODO IMRPOVE FORMAT RESPONSE
	res, err := uc.authUsc.Login(c, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, "Failed to Login user", err.Err.Error(), nil, fiber.StatusUnauthorized)
	}

	// @TODO IMRPOVE FORMAT RESPONSE
	// return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
	// 	"data": res,
	// })
	return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, res, fiber.StatusCreated)
}

func (uc *AuthControllerImpl) Register(ctx *fiber.Ctx) error {
	// c := ctx.Context()

	// data := new(authmodel.CreateUser)
	// if err := ctx.BodyParser(data); err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": err.Error(),
	// 	})
	// }

	// _, err := uc.authUsc.CreateUsers(c, *data)
	// if err != nil {
	// 	return helper.BuildResponse(ctx, false, "Failed to POST data", err.Err.Error(), nil, fiber.StatusBadRequest)
	// }

	// return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, "Register Succeed", fiber.StatusCreated)

    c := ctx.Context()

    data := new(dto.CreateUser)
    if err := ctx.BodyParser(data); err != nil {
        return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
    }

    _, err := uc.authUsc.CreateUsers(c, *data)
    if err != nil {
        return helper.BuildResponse(ctx, false, "Failed to POST data", err.Err.Error(), nil, fiber.StatusBadRequest)
    }

    return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, "Register Succeed", fiber.StatusCreated)

}
