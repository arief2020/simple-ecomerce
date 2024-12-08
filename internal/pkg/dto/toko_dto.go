package dto

type CreateTokoReq struct {
	UserID uint `json:"id_user" validate:"required"`
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

type TokoFilter struct {
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
	Nama  string `query:"nama"`
}

type AllTokoResp struct {
	Page  int        `query:"page"`
	Limit int        `query:"limit"`
	Data  []TokoResp `json:"data"`
}
