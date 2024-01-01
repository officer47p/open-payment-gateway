package stores

type Transaction struct {
	Id        uint `json:"id"`
	InvoiceId uint `json:"invoiceId"`
	//AmountFiat
	AmountCrypto         float64 `json:"amountCrypto"`
	Currency             string  `json:"currency"`
	Network              string  `json:"network"`
	NetworkTransactionId string  `json:"networkTransactionId"`
	AddressId            uint    `json:"addressId"`
	Status               string  `json:"status"`
}
