package models

// BillingInfo
type BillingInfo struct {
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

// Transaction
type Transaction struct {
	TxAmount int `json:"tx_amount"`
	// TxCardIssuer string      `json:"tx_card_issuer"`
	TxCardNumber string      `json:"tx_card_number"`
	TxCardCv     string      `json:"tx_card_cv"`
	BillingInfo  BillingInfo `json:"billing_info"`
}

// TransactionStatus
type TransactionStatus struct {
	Proceed        bool   `json:"proceed"`
	TxAmountIntent int    `json:"tx_amount_intent"`
	TxStatusCode   int    `json:"tx_status_code"`
	TxMessage      string `json:"tx_message"`
}
