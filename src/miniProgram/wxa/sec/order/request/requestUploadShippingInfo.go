package request

type RequestUploadShippingInfo struct {
	OrderKey       OrderKey       `json:"order_key"`
	LogisticsType  int8           `json:"logistics_type"`
	DeliveryMode   int8           `json:"delivery_mode"`
	IsAllDelivered bool           `json:"is_all_delivered,omitempty"`
	ShippingList   []ShippingList `json:"shipping_list"`
	UploadTime     string         `json:"upload_time"`
	Payer          struct {
		OpenID string `json:"openid"`
	} `json:"payer"`
}

type OrderKey struct {
	OrderNumberType int8   `json:"order_number_type"`
	TransactionID   string `json:"transaction_id"`
	MchID           string `json:"mchid,omitempty"`
	OutTradeNo      string `json:"out_trade_no,omitempty"`
}

type ShippingList struct {
	TrackingNo     string `json:"tracking_no"`
	ExpressCompany string `json:"express_company"`
	ItemDesc       string `json:"item_desc"`
	Contact        struct {
		ConsignorContact string `json:"consignor_contact,omitempty"`
		ReceiverContact  string `json:"receiver_contact,omitempty"`
	} `json:"contact"`
}
