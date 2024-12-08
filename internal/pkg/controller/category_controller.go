package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetAllCategory(ctx *fiber.Ctx) error
	GetCategoryByID(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategoryByID(ctx *fiber.Ctx) error
	DeleteCategoryByID(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	categoryUsc usecase.CategoryUseCase
}

func NewCategoryController(categoryUsc usecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		categoryUsc: categoryUsc,
	}
}

func (c *CategoryControllerImpl) GetAllCategory(ctx *fiber.Ctx) error {
	categories, err := c.categoryUsc.GetAllCategory(ctx.Context())
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Category")
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Err, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, categories, fiber.StatusOK)
}

func (c *CategoryControllerImpl) GetCategoryByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Params Category ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Error(), nil, fiber.StatusBadRequest)
	}

	category, errUsc := c.categoryUsc.GetCategoryByID(ctx.Context(), uint(id))
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Category By ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, category, fiber.StatusOK)
}

func (c *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	data := new(dto.CategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Body")
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	category, errUsc := c.categoryUsc.CreateCategory(ctx.Context(), *data)
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Category")
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errUsc.Err, nil, errUsc.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to CREATE data", nil, category, fiber.StatusCreated)
}

func (c *CategoryControllerImpl) UpdateCategoryByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Params Category ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.CategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Body")
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	_, errUsc := c.categoryUsc.UpdateCategoryByID(ctx.Context(), uint(id), *data)
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update Category By ID")
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errUsc.Err, nil, errUsc.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to UPDATE data", nil, "", fiber.StatusOK)
}

func (c *CategoryControllerImpl) DeleteCategoryByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Params Category ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Error(), nil, fiber.StatusBadRequest)
	}

	_, errUsc := c.categoryUsc.DeleteCategoryByID(ctx.Context(), uint(id))
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete Category By ID")
		return helper.BuildResponse(ctx, false, "Failed to DELETE data", errUsc.Err.Error(), nil, errUsc.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to DELETE data", nil, "", fiber.StatusOK)
}