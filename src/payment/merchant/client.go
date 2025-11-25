package merchant

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchant/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 查询平台账户实时余额
// https://pay.weixin.qq.com/doc/v3/partner/4012476700
func (comp *Client) FundBalance(ctx context.Context, accountType string) (*response.ResponseFundBalance, error) {

	result := &response.ResponseFundBalance{}

	endpoint := "/v3/merchant/fund/balance/" + accountType
	_, err := comp.SafeRequestV3(ctx, endpoint, &object.StringMap{}, http.MethodGet, &object.HashMap{}, nil, result)

	return result, err
}

// 图片上传API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter2_1_1.shtml
func (comp *Client) UploadImg(ctx context.Context, params *request.RequestMediaUpload) (*response.ResponseMediaUpload, error) {

	result := &response.ResponseMediaUpload{}

	var files *object.HashMap
	if params.File != "" {
		files = &object.HashMap{
			"file": params.File,
		}
	}

	var formData *kernel.UploadForm
	if params.Meta != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  "file",
					Value: params.Meta.Filename,
				},
			},
		}
	}
	options, _ := object.StructToHashMap(params.Meta)

	_, err := comp.BaseClient.HttpUploadJson(ctx, "/v3/merchant/media/upload", files, formData, options, nil, nil, &result)
	return result, err
}
