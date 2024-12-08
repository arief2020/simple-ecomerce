package controller

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/gofiber/fiber/v2"
)

type TokoController interface {
	GetMyToko(ctx *fiber.Ctx) error
	GetTokoByID(ctx *fiber.Ctx) error
	GetAllToko(ctx *fiber.Ctx) error
	UpdateMyToko(ctx *fiber.Ctx) error
	UpdateMyToko2(ctx *fiber.Ctx) error
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
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{"Invalid user ID"},
		})
	}

	toko, err := c.tokoUsc.GetMyToko(ctx.Context(), uint(id))
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{err.Error()},
		})
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       toko,
	})
}

func (c *TokoControllerImpl) GetTokoByID(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id_toko")
	if err != nil {
		return helper.ResponseWithJSON(&helper.JSONRespArgs{
			Ctx:        ctx,
			StatusCode: fiber.StatusBadRequest,
			Errors:     []string{"Invalid toko ID"},
		})
	}

	toko, errUsc := c.tokoUsc.GetTokoByID(ctx.Context(), uint(id))
	if errUsc != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       toko,
	})
}

func (c *TokoControllerImpl) GetAllToko(ctx *fiber.Ctx) error {

    filter := new(dto.TokoFilter)
    if err := ctx.QueryParser(filter); err != nil {
        return helper.BuildResponse(ctx, false, "Failed to GET data", "Failed to parse request query", nil, fiber.StatusBadRequest)
    }

    toko, err := c.tokoUsc.GetAllToko(ctx.Context(), dto.TokoFilter{
        Nama: filter.Nama,
        Limit:    filter.Limit,
        Page:     filter.Page,
    })
	if err != nil {
		return helper.BuildResponse(ctx, false, "Failed to GET data", err.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.ResponseWithJSON(&helper.JSONRespArgs{
		Ctx:        ctx,
		StatusCode: fiber.StatusOK,
		Data:       toko,
	})
}


	

func ensureUploadsFolderExists() error {
    // Path folder uploads
    uploadsFolder := "uploads"

    // Cek apakah folder sudah ada
    if _, err := os.Stat(uploadsFolder); os.IsNotExist(err) {
        // Buat folder jika belum ada
        if err := os.Mkdir(uploadsFolder, os.ModePerm); err != nil {
            return err
        }
    }
    return nil
}

func (c *TokoControllerImpl) UpdateMyToko2(ctx *fiber.Ctx) error {
    file, err := ctx.FormFile("file")
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "File is required",
        })
    }

    // Pastikan folder uploads ada
    if err := ensureUploadsFolderExists(); err != nil {
        log.Printf("Error ensuring uploads folder exists: %v", err)
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to prepare uploads folder",
        })
    }

    // Mendapatkan path file
    savePath := filepath.Join("uploads", file.Filename)

    // Simpan file
    if err := ctx.SaveFile(file, savePath); err != nil {
        log.Printf("Error saving file: %v", err)
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to save file",
        })
    }

    // Respon sukses
    return ctx.JSON(fiber.Map{
        "message": "File uploaded successfully",
        "file": fiber.Map{
            "name": file.Filename,
            "path": savePath,
        },
    })
}

func (c *TokoControllerImpl) UpdateMyToko(ctx *fiber.Ctx) error {
    idToko, err := ctx.ParamsInt("id_toko")
    if err != nil {
        return helper.ResponseWithJSON(&helper.JSONRespArgs{
            Ctx:        ctx,
            StatusCode: fiber.StatusBadRequest,
            Errors:     []string{"Invalid toko ID"},
        })
    }

    userId := ctx.Locals("userid").(string)
    userIdUint, err := strconv.ParseUint(userId, 10, 32)
    if err != nil {
        return helper.ResponseWithJSON(&helper.JSONRespArgs{
            Ctx:        ctx,
            StatusCode: fiber.StatusBadRequest,
            Errors:     []string{"Invalid user ID"},
        })
    }

    input := &dto.UpdateProfileTokoReq{
        NamaToko: ctx.FormValue("nama_toko"),
    }

    // Ambil file dari request
    file, _ := ctx.FormFile("photo") // File opsional

    res, errRes := c.tokoUsc.UpdateMyToko(ctx.Context(), uint(userIdUint), uint(idToko), input, file)
    if errRes != nil {
        return helper.ResponseWithJSON(&helper.JSONRespArgs{
            Ctx:        ctx,
            StatusCode: errRes.Code,
            Errors:     []string{errRes.Err.Error()},
        })
    }

    return helper.BuildResponse(ctx, true, "Succeed to UPDATE data", nil, res, fiber.StatusOK)
}