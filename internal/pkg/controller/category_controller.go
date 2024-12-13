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

// @Summary Get All Category
// @Description Endpoint for get all category
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} helper.Response{data=[]dto.CategoryResp} "Succeed to get all category"
// @Router /category [get]
func (c *CategoryControllerImpl) GetAllCategory(ctx *fiber.Ctx) error {
	categories, err := c.categoryUsc.GetAllCategory(ctx.Context())
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Category")
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, categories, fiber.StatusOK)
}


// @Summary Get Category By ID
// @Description Endpoint for get category by id (admin only)
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Success 200 {object} helper.Response{data=dto.CategoryResp} "Succeed to get category by id"
// @Router /category/{id} [get]
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

// @Summary Create Category
// @Description Endpoint for create category (admin only)
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param create-category body dto.CategoryReq true "Create Category"
// @Success 201 {object} helper.Response{data=dto.CategoryResp} "Succeed to create category"
// @Router /category [post]
func (c *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	data := new(dto.CategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Body")
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	category, errUsc := c.categoryUsc.CreateCategory(ctx.Context(), *data)
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Category")
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errUsc.Err.Error(), nil, errUsc.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to CREATE data", nil, category, fiber.StatusCreated)
}

// @Summary Update Category By ID
// @Description Endpoint for update category by id (admin only)
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Param update-category body dto.CategoryReq true "Update Category"
// @Success 200 {object} helper.Response{data=dto.CategoryResp} "Succeed to update category by id"
// @Router /category/{id} [put]
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
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errUsc.Err.Error(), nil, errUsc.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to UPDATE data", nil, "", fiber.StatusOK)
}


// @Summary Delete Category By ID
// @Description Endpoint for delete category by id (admin only)
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Category ID"
// @Success 200 {object} helper.Response{data=dto.CategoryResp} "Succeed to delete category by id"
// @Router /category/{id} [delete]	
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
