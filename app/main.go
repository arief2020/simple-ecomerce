package main

import (
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/infrastructure/container"
	"tugas_akhir_example/internal/utils"

	rest "tugas_akhir_example/internal/server/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"

	// swagger "github.com/gofiber/contrib/swagger"
	_ "tugas_akhir_example/docs"

	"github.com/gofiber/swagger"
)

// @title Simple Ecomerce API
// @version 1.0
// @description This is a sample server
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	// defer mysql.CloseDatabaseConnection(containerConf.Mysqldb)

	app := fiber.New()
	app.Use(logger.New())

	// swaggerCfg := swagger.Config{
	// 	BasePath: "/docs", // swagger ui base path
	// 	FilePath: "./docs/swagger.json",
	// }

	// app.Use(swagger.New(swaggerCfg))
	app.Get("/swagger/*", swagger.HandlerDefault) // default
	rest.HTTPRouteInit(app, containerConf)

	port := fmt.Sprintf("%s:%d", containerConf.Apps.Host, containerConf.Apps.HttpPort)
	if err := app.Listen(port); err != nil {
		// helper.Logger(helper.LoggerLevelFatal, "error", err)
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelFatal, "error")
	}
}
