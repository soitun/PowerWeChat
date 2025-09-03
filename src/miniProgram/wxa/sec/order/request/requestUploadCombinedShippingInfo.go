package request

type RequestUploadCombinedShippingInfo struct {
	OrderKey   OrderKey   `json:"order_key"`
	SubOrders  []SubOrder `json:"sub_orders"`
	UploadTime string     `json:"upload_time"`
	Payer      struct {
		OpenID string `json:"openid"`
	}
}

type SubOrder struct {
	OrderKey       OrderKey       `json:"order_key"`
	LogisticsType  int8           `json:"logistics_type"`
	DeliveryMode   int8           `json:"delivery_mode"`
	IsAllDelivered bool           `json:"is_all_delivered"`
	ShippingList   []ShippingList `json:"shipping_list"`
}
