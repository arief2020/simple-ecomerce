package helper

// @TODO : make helper response
import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Error   []interface{} `json:"errors"`
	Data    interface{}   `json:"data"`
}

const (
	SUCCEED = "Success"
)

func BuildResponse(ctx *fiber.Ctx, status bool, message string, errors interface{}, data interface{}, code int) error {
	var errorArray []interface{}
	if errors != nil {
		errorArray = append(errorArray, errors)
	}

	return ctx.Status(code).JSON(&Response{
		Status:  status,
		Message: message,
		Error:   errorArray,
		Data:    data,
	})
}

// punya orang
type JSONResp struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type JSONRespArgs struct {
	Ctx        *fiber.Ctx
	StatusCode int
	Errors     []string
	Data       interface{}
}

// ResponseWithJSON responses to the request with json format data
func ResponseWithJSON(args *JSONRespArgs) error {
	hasAnError := args.Errors != nil
	messagePrefix := "Succeed"
	if hasAnError {
		messagePrefix = "Failed"
	}
	message := fmt.Sprintf("%s to %s data", messagePrefix, strings.ToUpper(args.Ctx.Method()))

	return args.Ctx.Status(args.StatusCode).JSON(&JSONResp{
		Status:  !hasAnError,
		Message: message,
		Errors:  args.Errors,
		Data:    args.Data,
	})
}
