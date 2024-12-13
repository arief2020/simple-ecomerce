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

type UserController interface {
	GetMyProfile(ctx *fiber.Ctx) error
	UpdateMyProfile(ctx *fiber.Ctx) error
	GetMyAlamat(ctx *fiber.Ctx) error
	CreateMyNewAlamat(ctx *fiber.Ctx) error
	GetMyAlamatById(ctx *fiber.Ctx) error
	UpdateMyAlamatById(ctx *fiber.Ctx) error
	DeleteMyAlamatById(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	userUsc usecase.UserUseCase
}

func NewUserController(userUsc usecase.UserUseCase) UserController {
	return &UserControllerImpl{
		userUsc: userUsc,
	}
}

// @Summary Get My Profile
// @Description Endpoint for get my profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} helper.Response{data=dto.UserResp} "Succeed to get my profile"
// @Router /user [get]
func (uc *UserControllerImpl) GetMyProfile(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	user, errStruct := uc.userUsc.GetMyProfile(ctx.Context(), uint(id))
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Profile")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, user, fiber.StatusOK)
}

// @Summary Update My Profile
// @Description Endpoint for get my profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param update-profile body dto.UpdateUser true "Update My Profile"
// @Success 200 {object} helper.Response{data=string} "Succeed to update my profile"
// @Router /user [put]
func (uc *UserControllerImpl) UpdateMyProfile(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.UpdateUser)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Body")
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.UpdateMyProfile(ctx.Context(), uint(id), *data)
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update My Profile")
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, res, nil, "", fiber.StatusOK)
}

// @Summary Get My Alamat
// @Description Endpoint for get all my address
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param judul_alamat query string false "Judul Alamat"
// @Success 200 {object} helper.Response{data=[]dto.AlamatResp} "Succeed to get all my address"
// @Router /user/alamat [get]
func (uc *UserControllerImpl) GetMyAlamat(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	queryJudulAlamat := ctx.Query("judul_alamat")

	var filter dto.FiltersAlamat
	if queryJudulAlamat != "" {
		filter.JudulAlamat = queryJudulAlamat
	}

	fmt.Printf("Filter: %+v\n", filter)

	res, errStruct := uc.userUsc.GetMyAlamat(ctx.Context(), uint(id), dto.FiltersAlamat{
		JudulAlamat: filter.JudulAlamat,
	})
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Alamat")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

// @Summary Create My New Alamat
// @Description Endpoint for create new address
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param update-profile body dto.InserAlamatReq true "Success to create new address"
// @Success 200 {object} helper.Response{data=int} "Succeed to create new address"
// @Router /user/alamat [post]
func (uc *UserControllerImpl) CreateMyNewAlamat(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.InserAlamatReq)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Body")
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.CreateMyNewAlamat(ctx.Context(), uint(id), *data)
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create My New Alamat")
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, res, fiber.StatusCreated)
}

// @Summary Get My Alamat By ID
// @Description Endpoint for get address by id
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID of the address"
// @Success 200 {object} helper.Response{data=dto.AlamatResp} "Succeed to get address by id"
// @Router /user/alamat/{id} [get]
func (uc *UserControllerImpl) GetMyAlamatById(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	idAlamat, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Alamat ID")
		return helper.BuildResponse(ctx, false, "Invalid Alamat ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.GetMyAlamatById(ctx.Context(), uint(id), uint(idAlamat))
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Alamat By ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errStruct.Err.Error(), nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

// @Summary Get My Alamat By ID
// @Description Endpoint for get address by id
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID of the address"
// @Param update-alamat body dto.UpdateAlamatReq true "Success to update address"
// @Success 200 {object} helper.Response{data=dto.AlamatResp} "Succeed to update address by id"
// @Router /user/alamat/{id} [put]
func (uc *UserControllerImpl) UpdateMyAlamatById(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	idAlamat, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Alamat ID")
		return helper.BuildResponse(ctx, false, "Invalid Alamat ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.UpdateAlamatReq)
	if err := ctx.BodyParser(data); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Body")
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.UpdateMyAlamatById(ctx.Context(), uint(id), uint(idAlamat), *data)
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update My Alamat By ID")
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errStruct.Err.Error(), nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, res, nil, "", fiber.StatusOK)
}

// @Summary Delete My Alamat By ID
// @Description Endpoint for delete address by id
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID of the address"
// @Success 200 {object} helper.Response{data=string} "Succeed to update address by id"
// @Router /user/alamat/{id} [delete]
func (uc *UserControllerImpl) DeleteMyAlamatById(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse User ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	idAlamat, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Alamat ID")
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.DeleteMyAlamatById(ctx.Context(), uint(id), uint(idAlamat))
	if errStruct != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete My Alamat By ID")
		return helper.BuildResponse(ctx, false, "Failed to DELETE data", errStruct.Err.Error(), nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, res, nil, "", fiber.StatusOK)
}
