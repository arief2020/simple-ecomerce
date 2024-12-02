package handler

import (
	"github.com/gofiber/fiber/v2"

	// "tugas_akhir_example/internal/infrastructure/container"
	// "tugas_akhir_example/internal/pkg/controller"
	// "tugas_akhir_example/internal/pkg/repository"
	// "tugas_akhir_example/internal/pkg/usecase"
	ProvinceCityController "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/server/http/middleware"
)

// func AuthRoute(r fiber.Router, UserUsc usecase.UsersUseCase) {
// 	controller := authcontroller.NewAuthController(UserUsc)

// 	booksAPI := r.Group("/auth")
// 	booksAPI.Post("/register", controller.Register)
// 	booksAPI.Post("/login", controller.Login)
// }

func ProvinceCityRoute(r fiber.Router, ProvinceCityUsc usecase.ProvinceCityUseCase) {
	// repo := repository.NewProvinceCityRepository()
	// usecase := usecase.NewProvinceCityUseCase(repo)
	controller := ProvinceCityController.NewProvinceCityController(ProvinceCityUsc)

	provinceCityAPI := r.Group("/provcity")
	provinceCityAPI.Get("listprovincies", middleware.MiddlewareGetHeader, controller.GetAllProvinces)
	provinceCityAPI.Get("listcities/:prov_id", controller.GetAllCitiesByProvinceID)
	provinceCityAPI.Get("detailprovince/:prov_id", controller.GetProvinceByID)
	provinceCityAPI.Get("detailcity/:city_id", controller.GetCityByID)
}