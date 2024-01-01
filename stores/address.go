package stores

type Address struct {
	Id         uint   `json:"id"`
	ExternalId string `json:"externalId"`
	InvoiceId  uint   `json:"invoiceId"`
	MerchantId uint   `json:"merchantId"`
	Network    string `json:"network"`
	Currency   string `json:"currency"`
	Bip32Path  string `json:"bip32Path"`
	Memo       string `json:"memo"`
	Address    string `json:"address"`
}
