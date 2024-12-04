package dto

type AllProductFilter struct {
	NamaProduk string `query:"nama_produk"`
	CategoryID uint   `query:"category_id"`
	TokoID     uint   `query:"toko_id"`
	MaxHarga   int `query:"max_harga"`
	MinHarga   int `query:"min_harga"`
	Limit      int    `query:"limit"`
	Page       int    `query:"page"`
}

type AllProductResp struct {
	Data []ProductResp `json:"data"`
	Page int `json:"page"`
	Limit int `json:"limit"`
}

type ProductCreateReq struct {
	NamaProduk   string `json:"nama_produk" validate:"required"`
	CategoryID   uint   `json:"id_category" validate:"required"`
	HargaReseller string `json:"harga_reseller" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
	Stok         string `json:"stok" validate:"required"`
	Deskripsi    string `json:"deskripsi" validate:"required"`
}

type ProductUpdateReq struct {
	NamaProduk   string `json:"nama_produk" validate:"required"`
	CategoryID   uint   `json:"id_category" validate:"required"`
	HargaReseller string `json:"harga_reseller" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
	Stok         string `json:"stok" validate:"required"`
	Deskripsi    string `json:"deskripsi" validate:"required"`
}

type PhotoProductResp struct {
	Id uint `json:"id"`
	ProductID uint `json:"product_id"`
	Url string `json:"url"`
}
type ProductResp struct {
	ID           uint   `json:"id"`
	NamaProduk   string `json:"nama_produk" validate:"required"`
	Slug         string `json:"slug" validate:"required"`
	HargaReseller string `json:"harga_reseller" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
	Stok         string `json:"stok" validate:"required"`
	Deskripsi    string `json:"deskripsi" validate:"required"`
	Toko         TokoResp 	`json:"toko"`
	Category     CategoryResp `json:"category"`
	Photos       []PhotoProductResp `json:"photos"`
}