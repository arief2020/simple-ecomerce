package dto

type TokoFilter struct {
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Nama  string `query:"nama"`
}

type UpdateProfileTokoReq struct {
	NamaToko string `json:"nama_toko" validate:"required"`
}

type MyTokoResp struct {
	ID       uint    `json:"id"`
	NamaToko *string `json:"nama_toko"`
	UrlFoto  *string `json:"url_foto"`
	UserID   uint    `json:"id_user"`
}

type TokoResp struct {
	ID       uint    `json:"id"`
	NamaToko *string `json:"nama_toko"`
	UrlFoto  *string `json:"url_foto"`
}

type AllTokoResp struct {
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Data  []TokoResp `json:"data"`
}
