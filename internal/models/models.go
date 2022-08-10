package models

import "time"

// ProductSpecification
type ProductSpecification struct {
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
}

// BillingInfo
type BillingInfo struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}

// Transaction
type Transaction struct {
	CompanyUID   string               `json:"company_uid"`
	Product      ProductSpecification `json:"product"` // Business product specification
	TxAmount     int                  `json:"tx_amount"`
	TxDate       time.Time            `json:"tx_date"`
	TxCardNumber string               `json:"tx_card_number"`
	TxCardCv     string               `json:"tx_card_cv"`
	BillingInfo  BillingInfo          `json:"billing_info"`
}

// CompanyEnvoice
type CompanyEnvoice struct {
	ID           string      `json:"_id"`
	Date         string      `json:"date"`
	EnvoiceUUID  string      `json:"envoice_uuid"`
	MessageState string      `json:"message_state"`
	TxAccepted   bool        `json:"tx_accepted"`
	Transaction  Transaction `json:"transaction"`
}

// TransactionStatus
type TransactionStatus struct {
	Proceed        bool   `json:"proceed"`
	TxAmountIntent int    `json:"tx_amount_intent"`
	TxStatusCode   int    `json:"tx_status_code"`
	TxMessage      string `json:"tx_message"`
}

// TransactioData
type TransactionData struct {
	TxAccepted   bool        `json:"tx_accepted"`
	MessageState string      `json:"message_state"`
	Date         time.Time   `json:"date"`
	Transaction  Transaction `json:"transaction"`
}

// EnvoiceInfo record created
type EnvoiceInfo struct {
	Recived       bool   `json:"recived"`        // request was recived
	RecordCreated bool   `json:"record_created"` // new record was created no metter accepted or declined
	EnvoUUID      string `json:"envo_uuid"`
}
