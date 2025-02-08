package eventbus

import "encoding/json"

type NewTransactionNotification struct {
	BlockNumber int64  `json:"block_number"`
	BlockHash   string `json:"block_hash"`
	Network     string `json:"network"`
	Currency    string `json:"currency"`
	TxHash      string `json:"tx_hash"`
	TxType      string `json:"tx_type"`
	Value       string `json:"value"`
	From        string `json:"from"`
	To          string `json:"to"`
}

func (n NewTransactionNotification) ToJSON() (string, error) {
	b, err := json.Marshal(n)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
