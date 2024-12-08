package dto

type TransactionDetailReq struct {
	ProductID int `json:"product_id"`
	Kuantitas int `json:"kuantitas"`
}

type TransactionRequest struct {
	MethodBayar string                 `json:"method_bayar"`
	AlamatKirim int                    `json:"alamat_kirim"`
	DetailTrx   []TransactionDetailReq `json:"detail_trx"`
}

type AllTransactionReq struct {
	Limit  int    `query:"limit"`
	Page   int    `query:"page"`
	Search string `query:"search"`
}

type AllTransactionResponse struct {
	Data  []TransactionResponse `json:"data"`
	Page  int                   `json:"page"`
	Limit int                   `json:"limit"`
}

type TransactionProductResp struct {
	ID            uint               `json:"id"`
	NamaProduk    string             `json:"nama_produk" validate:"required"`
	Slug          string             `json:"slug" validate:"required"`
	HargaReseller string             `json:"harga_reseller" validate:"required"`
	HargaKonsumen string             `json:"harga_konsumen" validate:"required"`
	Deskripsi     string             `json:"deskripsi" validate:"required"`
	Toko          TokoResp           `json:"toko"`
	Category      CategoryResp       `json:"category"`
	Photos        []PhotoProductResp `json:"photos"`
}
type DetailTrx struct {
	Product    TransactionProductResp `json:"product"`
	Toko       TokoResp               `json:"toko"`
	Kuantitas  int                    `json:"kuantitas"`
	HargaTotal int                    `json:"harga_total"`
}

type TransactionResponse struct {
	ID          int         `json:"id"`
	HargaTotal  int         `json:"harga_total"`
	KodeInvoice string      `json:"kode_invoice"`
	MethodBayar string      `json:"method_bayar"`
	AlamatKirim AlamatResp  `json:"alamat_kirim"`
	DetailTrx   []DetailTrx `json:"detail_trx"`
}
