package dto

// AllProductFilter represents the query parameters for filtering products.
// @Description Query parameters that will be used to filter products
type AllProductFilter struct {
	NamaProduk string `query:"nama_produk" example:"Produk A"`
	CategoryID uint   `query:"category_id" example:"999"`
	TokoID     uint   `query:"toko_id" example:"999"`
	MaxHarga   int    `query:"max_harga" example:"9999"`
	MinHarga   int    `query:"min_harga" example:"9999"`
	Limit      int    `query:"limit" example:"10"`
	Page       int    `query:"page" example:"1"`
}

// AllProductResp represents the response data for getting all products.
// @Description Data that will be returned when getting all products
type AllProductResp struct {
	Data  []ProductResp `json:"data"`
	Page  int           `json:"page" example:"1"`
	Limit int           `json:"limit" example:"10"`
}

// ProductCreateReq represents the request data for creating a new product.
// @Description Data that will be used to create a new product
type ProductCreateReq struct {
	NamaProduk    string `form:"nama_produk" validate:"required" example:"Produk A"`
	CategoryID    uint   `form:"category_id" validate:"required" example:"999"`
	HargaReseller string `form:"harga_reseller" validate:"required" example:"9999"`
	HargaKonsumen string `form:"harga_konsumen" validate:"required" example:"99999"`
	Stok          string `form:"stok" validate:"required" example:"9999"`
	Deskripsi     string `form:"deskripsi" validate:"required" example:"Produk A"`
}

// ProductUpdateReq represents the request data for updating a product.
// @Description Data that will be used to update a product
type ProductUpdateReq struct {
	NamaProduk    string `json:"nama_produk" validate:"required" example:"Produk A"`
	CategoryID    uint   `json:"id_category" validate:"required" example:"999"`
	HargaReseller string `json:"harga_reseller" validate:"required" example:"9999"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required" example:"99999"`
	Stok          string `json:"stok" validate:"required" example:"9999"`
	Deskripsi     string `json:"deskripsi" validate:"required" example:"Produk A"`
}

// PhotoProductResp represents the response data for getting product photos.
// @Description Data that will be returned when getting product photos
type PhotoProductResp struct {
	Id        uint   `json:"id" example:"999"`
	ProductID uint   `json:"product_id" example:"999"`
	Url       string `json:"url" example:"https://example.com/image.jpg"`
}

// ProductResp represents the response data for getting a product.
// @Description Data that will be returned when getting a product
type ProductResp struct {
	ID            uint               `json:"id" example:"999"`
	NamaProduk    string             `json:"nama_produk" validate:"required" example:"Produk A"`
	Slug          string             `json:"slug" validate:"required" example:"produk-a"`
	HargaReseller string             `json:"harga_reseller" validate:"required" example:"9999"`
	HargaKonsumen string             `json:"harga_konsumen" validate:"required" example:"99999"`
	Stok          string             `json:"stok" validate:"required" example:"9999"`
	Deskripsi     string             `json:"deskripsi" validate:"required" example:"Produk A"`
	Toko          TokoResp           `json:"toko"`
	Category      CategoryResp       `json:"category"`
	Photos        []PhotoProductResp `json:"photos"`
}
