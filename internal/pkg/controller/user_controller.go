package controller

import (
	"fmt"
	"strconv"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

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

func (uc *UserControllerImpl) GetMyProfile(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	user, errStruct := uc.userUsc.GetMyProfile(ctx.Context(), uint(id))
	if errStruct != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, user, fiber.StatusOK)
}

func (uc *UserControllerImpl) UpdateMyProfile(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.UpdateUser)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.UpdateMyProfile(ctx.Context(), uint(id), *data)
	if errStruct != nil {
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, res, nil, "", fiber.StatusOK)
}

func (uc *UserControllerImpl) GetMyAlamat(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	// filter := new(dto.FiltersAlamat)
	// if err := ctx.QueryParser(filter); err != nil {
	// 	return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	// }

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
		return helper.BuildResponse(ctx, false, "Failed to GET data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

func (uc *UserControllerImpl) CreateMyNewAlamat(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.InserAlamatReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.CreateMyNewAlamat(ctx.Context(), uint(id), *data)
	if errStruct != nil {
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errStruct.Err, nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to POST data", nil, res, fiber.StatusCreated)
}

func (uc *UserControllerImpl) GetMyAlamatById(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	idAlamat, err := strconv.ParseUint(ctx.Params("id"), 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.GetMyAlamatById(ctx.Context(), uint(id), uint(idAlamat))
	if errStruct != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data", errStruct.Err.Error(), nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, res, fiber.StatusOK)
}

func (uc *UserControllerImpl) UpdateMyAlamatById(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	idAlamat, err := strconv.ParseUint(ctx.Params("id"), 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	data := new(dto.UpdateAlamatReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, "Failed to parse request body", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.UpdateMyAlamatById(ctx.Context(), uint(id), uint(idAlamat), *data)
	if errStruct != nil {
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errStruct.Err.Error(), nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, res, nil, "", fiber.StatusOK)
}

func (uc *UserControllerImpl) DeleteMyAlamatById(ctx *fiber.Ctx) error {
	idStr := ctx.Locals("userid").(string) // Ambil ID sebagai string
	id, err := strconv.ParseUint(idStr, 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	idAlamat, err := strconv.ParseUint(ctx.Params("id"), 10, 32) // Konversi string ke uint
	if err != nil {
		return helper.BuildResponse(ctx, false, "Invalid user ID", err.Error(), nil, fiber.StatusBadRequest)
	}

	res, errStruct := uc.userUsc.DeleteMyAlamatById(ctx.Context(), uint(id), uint(idAlamat))
	if errStruct != nil {
		return helper.BuildResponse(ctx, false, "Failed to DELETE data", errStruct.Err.Error(), nil, errStruct.Code)
	}

	return helper.BuildResponse(ctx, true, res, nil, "", fiber.StatusOK)
}