package dto

type ListProvResp struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProvinceFilter struct {
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}

// ProvinceResp represents the province data.
// @Description Data provinsi yang terkait dengan user
type ProvinceResp struct {
	Id   string `json:"id" example:"1"`
	Name string `json:"name" example:"Jawa Barat"`
}

// CityResp represents the city data.
// @Description Data kota yang terkait dengan user
type CityResp struct {
	Id         string `json:"id" example:"1"`
	ProvinceId string `json:"province_id" example:"1"`
	Name       string `json:"name" example:"Bandung"`
}
