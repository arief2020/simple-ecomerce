package middleware

import (
	"encoding/json"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type simpleHeader struct {
	Name      string `reqHeader:"name"`
	Pass      string `reqHeader:"pass"`
	ID        string `reqHeader:"id"`
	Timestamp string `reqHeader:"timestamp"`
}

type authHeader struct {
	Token string `reqHeader:"token"`
}

func MiddlewareGetHeader(ctx *fiber.Ctx) error {
	header := new(simpleHeader)
	if err := ctx.ReqHeaderParser(header); err != nil {
		return helper.BuildResponse(ctx, false, "Failed to parse request header", err.Error(), nil, fiber.StatusBadRequest)
	}

	ress, err := json.MarshalIndent(header, "", " ")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error marshal header")
	}
	helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelInfo, string(ress))

	return ctx.Next()
}

func MiddlewareAuth(ctx *fiber.Ctx) error {
	header := new(authHeader)
	if err := ctx.ReqHeaderParser(header); err != nil {
		return helper.BuildResponse(ctx, false, "Failed to parse request header", err.Error(), nil, fiber.StatusBadRequest)
	}

	ress, err := json.MarshalIndent(header, "", " ")
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error marshal header")
	}
	helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelInfo, string(ress))

	data, err := utils.DecodeToken(header.Token)
	if err != nil {
		return helper.BuildResponse(ctx, false, "Failed to decode token", err.Error(), nil, fiber.StatusUnauthorized)
	}
	fmt.Printf("User Data: %+v\n", data)

	ctx.Locals("userid", data["id"])
	ctx.Locals("useremail", data["email"])
	ctx.Locals("isAdmin", data["is_admin"])
	return ctx.Next()
}

func MiddlewareAdmin(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("isAdmin").(bool)
	if !isAdmin {
		return helper.BuildResponse(ctx, false, "Failed to POST data", "Unauthorized", nil, fiber.StatusUnauthorized)
	}
	return ctx.Next()
}
