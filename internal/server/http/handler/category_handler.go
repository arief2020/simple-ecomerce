package handler

import (
	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/server/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(r fiber.Router, categoryUsc usecase.CategoryUseCase) {
	controller := controller.NewCategoryController(categoryUsc)

	categoryAPI := r.Group("/category")
	categoryAPI.Get("", controller.GetAllCategory)
	categoryAPI.Get("/:id", middleware.MiddlewareAuth, middleware.MiddlewareAdmin, controller.GetCategoryByID)
	categoryAPI.Post("", middleware.MiddlewareAuth, middleware.MiddlewareAdmin, controller.CreateCategory)
	categoryAPI.Put("/:id", middleware.MiddlewareAuth, middleware.MiddlewareAdmin, controller.UpdateCategoryByID)
	categoryAPI.Delete("/:id", middleware.MiddlewareAuth, middleware.MiddlewareAdmin, controller.DeleteCategoryByID)
}
