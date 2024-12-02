package handler

import (
	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/server/http/middleware"
)

func TrxRoute(r fiber.Router, trxUsc usecase.TrxUseCase) {
	controller := controller.NewTrxController(trxUsc)

	trxAPI := r.Group("/trx")
	trxAPI.Get("", middleware.MiddlewareAuth, controller.GetAllTransctionByUserID)
	trxAPI.Get("/:id_trx", middleware.MiddlewareAuth, controller.GetTransactionByID)
	trxAPI.Post("", middleware.MiddlewareAuth, controller.CreateTransction)
}
