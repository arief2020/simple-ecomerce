package controller

import (
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(ctx *fiber.Ctx) error
	GetAllProduct(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error
	UpdateProductByID(ctx *fiber.Ctx) error
	DeleteProductByID(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	productUsc usecase.ProductUseCase
}

func NewProductController(productUsc usecase.ProductUseCase) ProductController {
	return &ProductControllerImpl{productUsc: productUsc}
}

func (c *ProductControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	data := new(dto.ProductCreateReq)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userId := ctx.Locals("userid").(string)
	userIdInt := utils.StringToUint(userId)

	dataReq := &dto.ProductCreateReq{
		NamaProduk:    data.NamaProduk,
		CategoryID:    data.CategoryID,
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
	}

	files, err := ctx.MultipartForm()
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelPanic, "Failed to parse multipart form")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse multipart form",
		})
	}

	photos := files.File["photos"]

	resUsc, errUsc := c.productUsc.CreateProduct(ctx.Context(), *dataReq, photos, uint(userIdInt))
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Product")
		return helper.BuildResponse(ctx, false, "Failed to CREATE data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to CREATE data", nil, resUsc, fiber.StatusOK)
}

func (c *ProductControllerImpl) GetAllProduct(ctx *fiber.Ctx) error {
	filter := new(dto.AllProductFilter)
	if err := ctx.QueryParser(filter); err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Parse Request Query")
		return helper.BuildResponse(ctx, false, "Failed to parse query params", err.Error(), nil, fiber.StatusBadRequest)
	}
	resUsc, errUsc := c.productUsc.GetAllProduct(ctx.Context(), dto.AllProductFilter{
		NamaProduk: filter.NamaProduk,
		CategoryID: filter.CategoryID,
		TokoID:     filter.TokoID,
		MinHarga:   filter.MinHarga,
		MaxHarga:   filter.MaxHarga,
		Limit:      filter.Limit,
		Page:       filter.Page,
	})
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get All Product")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, resUsc, fiber.StatusOK)
}

func (c *ProductControllerImpl) GetProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id_product")
	productId := utils.StringToUint(id)
	resUsc, errUsc := c.productUsc.GetProductByID(ctx.Context(), uint(productId))
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Product By ID")
		return helper.BuildResponse(ctx, false, "Failed to GET data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to GET data", nil, resUsc, fiber.StatusOK)
}

func (c *ProductControllerImpl) DeleteProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id_product")
	productId := utils.StringToUint(id)
	_, errUsc := c.productUsc.DeleteProductByID(ctx.Context(), uint(productId))
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete Product By ID")
		return helper.BuildResponse(ctx, false, "Failed to DELETE data", errUsc.Err.Error(), nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to DELETE data", nil, "", fiber.StatusOK)
}

func (c *ProductControllerImpl) UpdateProductByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id_product")
	productId := utils.StringToUint(id)

	categoryId := ctx.FormValue("category_id")
	formatCategory := utils.StringToUint(categoryId)

	data := dto.ProductUpdateReq{
		NamaProduk:    ctx.FormValue("nama_produk"),
		CategoryID:    uint(formatCategory),
		HargaReseller: ctx.FormValue("harga_reseller"),
		HargaKonsumen: ctx.FormValue("harga_konsumen"),
		Stok:          ctx.FormValue("stok"),
		Deskripsi:     ctx.FormValue("deskripsi"),
	}

	_, errUsc := c.productUsc.UpdateProductByID(ctx.Context(), uint(productId), data)
	if errUsc != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update Product By ID")
		return helper.BuildResponse(ctx, false, "Failed to UPDATE data", errUsc, nil, fiber.StatusBadRequest)
	}

	return helper.BuildResponse(ctx, true, "Succeed to UPDATE data", nil, "", fiber.StatusOK)
}
