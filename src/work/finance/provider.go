package finance

import (
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/types"
)

func RegisterProvider(app kernel.ApplicationInterface) (contract.Client, error) {

	cfg := app.GetConfig()

	corpID := cfg.Get("corp_id", "").(string)
	secret := cfg.Get("secret", "").(string)

	if corpID == "" || secret == "" {
		return nil, fmt.Errorf("corp_id / wecom_secret 不能为空")
	}

	platRaw := cfg.Get("finance_sdk.platform", "").(string)
	path := cfg.Get("finance_sdk.path", "").(string)
	if path == "" {
		return nil, fmt.Errorf("sdk_path 不能为空（目录或完整库文件路径）")
	}

	var plat types.SDKPlatform
	if platRaw != "" {
		p, err := CanonicalizePlatform(platRaw)
		if err != nil {
			return nil, fmt.Errorf("sdk_platform 非法: %w", err)
		}
		plat = p
	}

	fc := &FinanceConfig{
		SDKPath:     path,
		SDKPlatform: plat,
		CorpID:      corpID,
		CorpSecret:  secret,
	}
	return NewClient(fc)
}
