package handler

import (
	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/server/http/middleware"
)

func TokoRoute(r fiber.Router, tokoUsc usecase.TokoUseCase) {
	controller := controller.NewTokoController(tokoUsc)

	tokoAPI := r.Group("/toko")
	tokoAPI.Get("", controller.GetAllToko)
	tokoAPI.Get("my", middleware.MiddlewareAuth, controller.GetMyToko)
	tokoAPI.Put("/:id_toko", middleware.MiddlewareAuth, controller.UpdateMyToko)
	tokoAPI.Get("/:id_toko",middleware.MiddlewareAuth ,controller.GetTokoByID)
}
