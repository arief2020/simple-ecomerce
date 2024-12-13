package dto

// InserAlamatReq represents the request data for inserting a new address.
// @Description Data that will be used to insert a new address
type InserAlamatReq struct {
	JudulAlamat  string `json:"judul_alamat" validate:"required" example:"Alamat Rumah"`
	NamaPenerima string `json:"nama_penerima" validate:"required" example:"Budi"`
	NoTelp       string `json:"no_telp" validate:"required" example:"08123456789"`
	DetailAlamat string `json:"detail_alamat" validate:"required" example:"Jl. Contoh No. 123"`
}

// UpdateAlamatReq represents the request data for updating an address.
// @Description Data that will be used to update an address
type UpdateAlamatReq struct {
	NamaPenerima string `json:"nama_penerima" validate:"required" example:"Budi"`
	NoTelp       string `json:"no_telp" validate:"required" example:"08123456789"`
	DetailAlamat string `json:"detail_alamat" validate:"required" example:"Jl. Contoh No. 123"`
}

// FiltersAlamat represents the query parameters for filtering addresses.
// @Description Query parameters that will be used to filter addresses
type FiltersAlamat struct {
	JudulAlamat string `query:"judul_alamat" example:"Alamat Rumah"`
}

// AlamatResp represents the response data for an address.
// @Description Data that will be returned in the response
type AlamatResp struct {
	Id           uint   `json:"id" example:999`
	JudulAlamat  string `json:"judul_alamat" validate:"required" example:"Alamat Rumah"`
	NamaPenerima string `json:"nama_penerima" validate:"required" example:"Budi"`
	NoTelp       string `json:"no_telp" validate:"required" example:"08123456789"`
	DetailAlamat string `json:"detail_alamat" validate:"required" example:"Jl. Contoh No. 123"`
}
