package handler

import (
	"github.com/gofiber/fiber/v2"

	authcontroller "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
)

func AuthRoute(r fiber.Router, AuthUsc usecase.AuthsUseCase) {
	controller := authcontroller.NewAuthController(AuthUsc)

	autAPI := r.Group("/auth")
	autAPI.Post("/register", controller.Register)
	autAPI.Post("/login", controller.Login)
}
