package stores

type Invoice struct {
	Id          uint    `json:"id"`
	ExternalId  string  `json:"externalId"`
	Currency    string  `json:"currency"`
	TotalAmount float64 `json:"totalAmount"`
	PayedAmount float64 `json:"payedAmount"`
	Status      string  `json:"status"`
	MerchantId  uint    `json:"merchantId"`
}
