package dto

// CategoryResp represents the response data for getting a category.
// @Description Data that will be returned in the response
type CategoryResp struct {
	ID           uint   `json:"id" example:"999"`
	NamaCategory string `json:"nama_category" example:"Makanan"`
}

// CategoryReq represents the request data for creating a new category.
// @Description Data that will be used to create or update a category
type CategoryReq struct {
	NamaCategory string `json:"nama_category" validate:"required" example:"Makanan"`
}
