package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseFundBalance struct {
	response.ResponsePayment
	AvailableAmount int `json:"available_amount"`
	PendingAmount   int `json:"pending_amount"`
}
