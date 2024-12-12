package dto

type CreateUser struct {
	Email        string `json:"email" validate:"required,email"`
	KataSandi    string `json:"kata_sandi" validate:"required"`
	Name         string `json:"name" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Tentang      string `json:"tentang" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}

type Login struct {
	NoTelp    string `json:"no_telp" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type UpdateUser struct {
	Email        string `json:"email" validate:"required,email"`
	Nama         string `json:"nama" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Tentang      string `json:"tentang" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	IdProvinsi   string `json:"id_provinsi" validate:"required"`
	IdKota       string `json:"id_kota" validate:"required"`
}
// LoginRes represents the response data for a successful login.
// @Description Data yang dikembalikan setelah user berhasil login
type LoginRes struct {
	Nama         string        `json:"nama" example:"John Doe"`
	NoTelp       string        `json:"no_telp" example:"1234567890"`
	TanggalLahir string        `json:"tanggal_lahir" example:"1990-01-01"`
	Tentang      string        `json:"tentang" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit."`
	Pekerjaan    string        `json:"pekerjaan" example:"Software Engineer"`
	Email        string        `json:"email" example:"L2DQK@example.com"`
	IdProvinsi   *ProvinceResp `json:"id_provinsi"`
	IdKota       *CityResp     `json:"id_kota"`
	Token        string        `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

type UserResp struct {
	Nama         string        `json:"nama"`
	NoTelp       string        `json:"no_telp"`
	TanggalLahir string        `json:"tanggal_lahir"`
	Tentang      string        `json:"tentang"`
	Pekerjaan    string        `json:"pekerjaan"`
	Email        string        `json:"email"`
	IdKota       *CityResp     `json:"id_kota"`
	IdProvinsi   *ProvinceResp `json:"id_provinsi"`
	IsAdmin      bool          `json:"is_admin"`
}
