package tests

import (
	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	rest "tugas_akhir_example/internal/server/http"
)

func setupApp() *fiber.App {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()

	app := fiber.New()
	app.Use(logger.New())

	rest.HTTPRouteInit(app, containerConf)

	return app
}
