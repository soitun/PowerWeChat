package finance

import (
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/types"
	"strings"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/drivers"
)

func NewClient(cfg *FinanceConfig) (contract.Client, error) {
	if cfg == nil {
		return nil, fmt.Errorf("FinanceConfig 不能为空")
	}

	// 平台归一化
	if cfg.SDKPlatform != "" {
		p, err := CanonicalizePlatform(string(cfg.SDKPlatform))
		if err != nil {
			return nil, err
		}
		cfg.SDKPlatform = p
	}

	// 解析库路径
	libPath, err := ResolveLibPath(cfg)
	if err != nil {
		return nil, err
	}

	// 友好提示：常见官方只提供 Linux 包
	switch cfg.SDKPlatform {
	case types.SDKPlatformLinuxAMD64, types.SDKPlatformLinuxARM64:
		// ok
	default:
		// 不强拦，交给加载器报更具体错误（如缺 OpenSSL 3）
	}

	cli, err := drivers.NewDynamicWithPath(libPath, cfg.CorpID, cfg.CorpSecret)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "libssl.so.3") || strings.Contains(msg, "libcrypto.so.3") {
			return nil, fmt.Errorf("%v（可能缺少 OpenSSL 3 运行库）", err)
		}
		return nil, err
	}
	return cli, nil
}
