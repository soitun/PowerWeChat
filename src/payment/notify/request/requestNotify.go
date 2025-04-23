package request

import (
	"net/http"
	"time"
)

// Request 微信支付通知请求结构
type RequestNotify struct {
	ID           string             `json:"id"`
	CreateTime   *time.Time         `json:"create_time"`
	EventType    string             `json:"event_type"`
	ResourceType string             `json:"resource_type"`
	Resource     *EncryptedResource `json:"resource"`
	Summary      string             `json:"summary"`

	// 原始通知请求
	RawRequest *http.Request
}

// EncryptedResource 微信支付通知请求中的内容
type EncryptedResource struct {
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
	OriginalType   string `json:"original_type"`

	Plaintext string // Ciphertext 解密后内容
}

// TransferBill 转账通知
type TransferBillCallback struct {
	OutBillNo      string `json:"out_bill_no"`      // 商户转账单号
	TransferBillNo string `json:"transfer_bill_no"` // 微信转账单号
	State          string `json:"state"`            // 转账状态
	MchId          string `json:"mch_id"`           // 商户号
	TransferAmount int64  `json:"transfer_amount"`  // 转账金额
	OpenId         string `json:"openid"`           // 用户openID
	FailReason     string `json:"fail_reason"`      // 转账失败原因
	CreateTime     string `json:"create_time"`      // 创建时间
	UpdateTime     string `json:"update_time"`      // 更新时间
}
