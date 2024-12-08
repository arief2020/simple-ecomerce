package dto

type CategoryResp struct {
	ID           uint   `json:"id"`
	NamaCategory string `json:"nama_category"`
}

type CategoryReq struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}
