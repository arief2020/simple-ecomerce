package http

import (
	route "tugas_akhir_example/internal/server/http/handler"

	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	route.AuthRoute(api, containerConf.AuthUsc)
	route.UserRoute(api, containerConf.UserUsc)
	route.ProvinceCityRoute(api, containerConf.ProvinceCityUsc)
	route.TokoRoute(api, containerConf.TokoUsc)
	route.CategoryRoute(api, containerConf.CategoryUsc)
	route.ProductRoute(api, containerConf.ProductUsc)
	route.TrxRoute(api, containerConf.TrxUsc)
}
