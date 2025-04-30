package request

type RequestGetShopQRCode struct {
	WeComCorpId string `json:"wecom_corp_id"`
	WeComUserId string `json:"wecom_user_id"`
}
