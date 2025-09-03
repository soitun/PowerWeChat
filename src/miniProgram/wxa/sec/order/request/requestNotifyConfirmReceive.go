package request

type RequestNotifyConfirmReceive struct {
	TransactionID   string `json:"transaction_id,omitempty"`
	MerchantID      string `json:"merchant_id,omitempty"`
	SubMerchantID   string `json:"sub_merchant_id,omitempty"`
	MerchantTradeNo string `json:"merchant_trade_no,omitempty"`
	ReceivedTime    int    `json:"received_time,omitempty"`
}
