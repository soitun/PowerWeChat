package request

type RequestGetOrderList struct {
	PayTimeRange *PayTimeRange `json:"pay_time_range,omitempty"`
	OrderState   int8          `json:"order_state,omitempty"`
	OpenID       string        `json:"openid,omitempty"`
	LastIndex    string        `json:"last_index,omitempty"`
	PageSize     int           `json:"page_size,omitempty"`
}

type PayTimeRange struct {
	BeginTime int `json:"begin_time,omitempty"`
	EndTime   int `json:"end_time,omitempty"`
}
