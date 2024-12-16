package tests

// import (
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAllCategory(t *testing.T) {
// 	app := setupApp()

// 	// resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/category", nil), -1)
// 	// if err != nil {
// 	// 	t.Fatalf("Failed to send request: %v", err)
// 	// }

// 	// assert.Equal(t, 200, resp.StatusCode)
// 	t.Run("Success - Get All Categories", func(t *testing.T) {
// 		req := httptest.NewRequest("GET", "/api/v1/categories", nil)
// 		resp, _ := app.Test(req)

// 		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
// 	})
// }
