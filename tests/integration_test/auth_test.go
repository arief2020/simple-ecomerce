package tests

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	app := setupApp()
	t.Run("Success - Register", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/api/v1/auth/register", strings.NewReader(`{
    "name": "joko",
    "kata_sandi": "123456",
    "no_telp": "089612312457",
    "tanggal_lahir": "02/01/2002",
    "tentang": "joko account",
    "jenis_kelamin": "Laki-Laki",
    "pekerjaan": "developer",
    "email":"testjoko@mail.com",
    "id_provinsi":"11",
    "id_kota":"1101"
}`))
		req.Header.Set("Content-Type", "application/json") // Pastikan header JSON
		resp, _ := app.Test(req, -1)                       // Jangan cache response

		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	})
}
