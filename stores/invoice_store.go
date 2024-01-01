package stores

import "open-payment-gateway/db"

type InvoiceStore struct{}

type NewInvoiceStore struct {
	DBConnection db.DBConnection
}
