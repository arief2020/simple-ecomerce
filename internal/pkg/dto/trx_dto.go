package dto

type TransactionDetailReq struct {
	ProductID int `json:"product_id"`
	Kuantitas int `json:"kuantitas"`
}

type TransactionRequest struct {
	MethodBayar string             `json:"method_bayar"`
	AlamatKirim int                `json:"alamat_kirim"`
	DetailTrx   []TransactionDetailReq `json:"detail_trx"`
}