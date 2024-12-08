package handler

import (
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/server/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(r fiber.Router, productUsc usecase.ProductUseCase) {
	controller := controller.NewProductController(productUsc)

	productApi := r.Group("/product")
	productApi.Post("", middleware.MiddlewareAuth, controller.CreateProduct)
	productApi.Get("", controller.GetAllProduct)
	productApi.Get("/:id_product", controller.GetProductByID)
	productApi.Put("/:id_product", middleware.MiddlewareAuth, controller.UpdateProductByID)
	productApi.Delete("/:id_product", middleware.MiddlewareAuth, controller.DeleteProductByID)

}
