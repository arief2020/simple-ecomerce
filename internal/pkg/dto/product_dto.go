package dto

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

type ProductResp struct {
	ID           uint   `json:"id"`
	NamaProduk   string `json:"nama_produk" validate:"required"`
	Slug         string `json:"slug" validate:"required"`
	HargaReseller string `json:"harga_reseller" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
	Stok         string `json:"stok" validate:"required"`
	Deskripsi    string `json:"deskripsi" validate:"required"`
}