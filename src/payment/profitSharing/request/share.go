package request

import (
	"encoding/xml"
	"time"
)

type RequestShare struct {
	SubMchID        string      `json:"sub_mchid,omitempty"` // 服务商模式下必填
	SubAppID        string      `json:"sub_appid,omitempty"` // 分账接收方类型包含PERSONAL_SUB_OPENID时必填
	AppID           string      `json:"appid,omitempty"`
	TransactionID   string      `json:"transaction_id,omitempty"` // OutTradeNo 和 TransactionID 二选一
	OutOrderNO      string      `json:"out_order_no,omitempty"`
	Receivers       []*Receiver `json:"receivers,omitempty"`
	UnfreezeUnSplit bool        `json:"unfreeze_unsplit"`
}

type Receiver struct {
	Type        string    `json:"type"`
	Account     string    `json:"account"`
	Name        string    `json:"name,omitempty"`
	Amount      int64     `json:"amount,omitempty"`
	Description string    `json:"description,omitempty"`
	Result      string    `json:"result,omitempty"`
	FailReason  string    `json:"fail_reason,omitempty"`
	DetailId    string    `json:"detail_id,omitempty"`
	CreateTime  time.Time `json:"create_time,omitempty"`
	FinishTime  time.Time `json:"finish_time,omitempty"`
}

type RequestShareReturnV3 struct {
	SubMchID    string `json:"sub_mchid,omitempty"`    // 服务商模式下必填
	OrderID     string `json:"order_id,omitempty"`     // 微信分账单号，微信系统返回的唯一标识。微信分账单号和商户分账单号二选一填写
	OutOrderNO  string `json:"out_order_no,omitempty"` // 微信分账单号和商户分账单号二选一填写
	OutReturnNO string `json:"out_return_no"`          // 此回退单号是商户在自己后台生成的一个新的回退单号，在商户后台唯一
	ReturnMchID string `json:"return_mchid"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}

type RequestShareReturn struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`

	AppID string `xml:"appid"`
	MchID string `xml:"mch_id"`
	//NonceStr string `xml:"nonce_str"`
	//SignType          string `xml:"sign_type"`
	//Sign              string `xml:"sign"`
	OutOrderNo        string `xml:"out_order_no"`
	OutReturnNo       string `xml:"out_return_no"`
	ReturnAccountType string `xml:"return_account_type"`
	ReturnAccount     string `xml:"return_account"`
	ReturnAmount      string `xml:"return_amount"`
	Description       string `xml:"description"`
}
