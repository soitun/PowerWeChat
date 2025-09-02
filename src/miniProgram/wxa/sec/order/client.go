package wxaSecOrder

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response4 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"net/http"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取小程序二维码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (comp *Client) CreateQRCode(ctx context.Context, path string, width int64) (*http.Response, error) {

	var header = &response4.ResponseHeaderMedia{}

	if width <= 0 {
		width = 430
	}

	data := &object.HashMap{
		"form_params": &object.HashMap{
			"path":  path,
			"width": width,
		},
	}

	rs, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/wxaapp/createwxaqrcode", http.MethodPost, data, &header, nil)

	return rs, err

}
