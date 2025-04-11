package request

type CodePayPayer struct {
	AuthCode string `json:"auth_code,omitempty"` // 付款码支付授权码
}

type CodePayAmount struct {
	Total    int    `json:"total,omitempty"`    // 总金额
	Currency string `json:"currency,omitempty"` // 货币类型
}

type CodePayStoreInfo struct {
	ID    string `json:"id,omitempty"`     // 门店编号
	OutID string `json:"out_id,omitempty"` // 商家自定义编码
}

type CodePaySceneInfo struct {
	DeviceID  string            `json:"device_id,omitempty"`  // 商户端设备号
	StoreInfo *CodePayStoreInfo `json:"store_info,omitempty"` // 商户门店信息
	DeviceIP  string            `json:"device_ip,omitempty"`  // 用户终端IP
}

type CodePayGoodsDetail struct {
	MerchantGoodsID string `json:"merchant_goods_id,omitempty"` // 商品编码
	WxpayGoodsID    string `json:"wxpay_goods_id,omitempty"`    // 微信支付商品编码
	GoodsName       string `json:"goods_name,omitempty"`        // 商品名称
	Quantity        int    `json:"quantity,omitempty"`          // 商品数量
	UnitPrice       int    `json:"unit_price,omitempty"`        // 商品单价
}

type CodePayDetail struct {
	CostPrice   int                   `json:"cost_price,omitempty"`   // 订单原价
	InvoiceID   string                `json:"invoice_id,omitempty"`   // 商品小票ID
	GoodsDetail []*CodePayGoodsDetail `json:"goods_detail,omitempty"` // 单品列表
}

type CodePaySettleInfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"` // 是否指定分账
}

type RequestCodePay struct {
	PrepayBase
	Description   string             `json:"description"`              // 商品描述
	OutTradeNo    string             `json:"out_trade_no"`             // 商户订单号
	Attach        string             `json:"attach"`                   // 附加数据
	GoodsTag      string             `json:"goods_tag,omitempty"`      // 订单优惠标记
	SupportFapiao bool               `json:"support_fapiao,omitempty"` // 电子发票入口开放标识
	Payer         *CodePayPayer      `json:"payer,omitempty"`          // 支付者
	Amount        *CodePayAmount     `json:"amount,omitempty"`         // 订单金额
	SceneInfo     *CodePaySceneInfo  `json:"scene_info,omitempty"`     // 场景信息
	Detail        *CodePayDetail     `json:"detail,omitempty"`         // 优惠功能
	SettleInfo    *CodePaySettleInfo `json:"settle_info,omitempty"`    // 结算信息
}
