package handler

import (
	"github.com/gofiber/fiber/v2"

	"tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/server/http/middleware"
)

func UserRoute(r fiber.Router, userUsc usecase.UserUseCase) {
	controller := controller.NewUserController(userUsc)

	userAPI := r.Group("/user")
	userAPI.Get("", middleware.MiddlewareAuth, controller.GetMyProfile)
	userAPI.Put("", middleware.MiddlewareAuth, controller.UpdateMyProfile)

	userAPI.Get("/alamat", middleware.MiddlewareAuth, controller.GetMyAlamat)
	userAPI.Post("/alamat", middleware.MiddlewareAuth, controller.CreateMyNewAlamat)
	userAPI.Get("alamat/:id", middleware.MiddlewareAuth, controller.GetMyAlamatById)
	userAPI.Put("alamat/:id", middleware.MiddlewareAuth, controller.UpdateMyAlamatById)
	userAPI.Delete("alamat/:id", middleware.MiddlewareAuth, controller.DeleteMyAlamatById)
}
