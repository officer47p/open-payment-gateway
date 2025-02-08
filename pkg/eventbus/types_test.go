package eventbus

import "testing"

func TestNewTransactionNotificationToJSON(t *testing.T) {
	notification := NewTransactionNotification{
		BlockNumber: 12345,
		BlockHash:   "0x123abc",
		Network:     "Mainnet",
		Currency:    "ETH",
		TxHash:      "0x456def",
		TxType:      "Transfer",
		Value:       "10.5",
		From:        "0xabcdef123",
		To:          "0x789ghi456",
	}

	expectedJSON := `{"block_number":12345,"block_hash":"0x123abc","network":"Mainnet","currency":"ETH","tx_hash":"0x456def","tx_type":"Transfer","value":"10.5","from":"0xabcdef123","to":"0x789ghi456"}`

	jsonStr, err := notification.ToJSON()
	if err != nil {
		t.Errorf("Error converting NewTransactionNotification to JSON: %v", err)
	}

	if jsonStr != expectedJSON {
		t.Errorf("Expected JSON: %s\nActual JSON: %s", expectedJSON, jsonStr)
	}
}
