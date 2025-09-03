package request

type RequestOpSpecialOrder struct {
	OrderID string `json:"order_id"`
	Types   int8   `json:"type"`
	DelayTo int    `json:"delay_to,omitempty"`
}
