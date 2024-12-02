package dto

type CreateTokoReq struct {
	UserID uint `json:"id_user" validate:"required"`
}

type UpdateProfileTokoReq struct {
	NamaToko  string `json:"nama_toko" validate:"required"`
}

type TokoResp struct {
	ID        uint   `json:"id"`
	NamaToko  *string `json:"nama_toko"`
	UrlFoto   *string `json:"url_foto"`
	UserID    uint   `json:"id_user"`
}