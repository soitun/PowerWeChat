package finance

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/types"
	"path/filepath"
	"runtime"
	"strings"
)

type FinanceConfig struct {
	SDKPath     string            // 目录 或 具体库文件路径
	SDKPlatform types.SDKPlatform // 建议显式指定；仅用于校验提示
	CorpID      string
	CorpSecret  string
}

// DetectCurrentPlatform 将 runtime.GOOS/GOARCH 归一化
func DetectCurrentPlatform() (types.SDKPlatform, error) {
	switch runtime.GOOS {
	case "linux":
		switch runtime.GOARCH {
		case "amd64":
			return types.SDKPlatformLinuxAMD64, nil
		case "arm64":
			return types.SDKPlatformLinuxARM64, nil
		}
	case "windows":
		switch runtime.GOARCH {
		case "amd64":
			return types.SDKPlatformWindowsAMD64, nil
		}
	case "darwin":
		switch runtime.GOARCH {
		case "arm64":
			return types.SDKPlatformDarwinARM64, nil
		}
	}
	return "", fmt.Errorf("未支持的平台: %s_%s", runtime.GOOS, runtime.GOARCH)
}

// CanonicalizePlatform 平台别名归一化
func CanonicalizePlatform(s string) (types.SDKPlatform, error) {
	k := strings.ToLower(strings.TrimSpace(s))
	k = strings.ReplaceAll(k, "-", "_")
	k = strings.ReplaceAll(k, "x64", "amd64")
	k = strings.ReplaceAll(k, "x86_64", "amd64")
	k = strings.ReplaceAll(k, "aarch64", "arm64")

	switch k {
	case string(types.SDKPlatformLinuxAMD64):
		return types.SDKPlatformLinuxAMD64, nil
	case string(types.SDKPlatformLinuxARM64):
		return types.SDKPlatformLinuxARM64, nil
	case "windows", "win", string(types.SDKPlatformWindowsAMD64):
		return types.SDKPlatformWindowsAMD64, nil
	case "darwin", "macos", string(types.SDKPlatformDarwinARM64):
		return types.SDKPlatformDarwinARM64, nil
	default:
		return "", fmt.Errorf("未知/未支持的 SDKPlatform: %s", s)
	}
}

// ResolveLibPath 统一把 FinanceConfig 解析为“最终库文件绝对路径”
func ResolveLibPath(cfg *FinanceConfig) (string, error) {
	if cfg == nil {
		return "", errors.New("FinanceConfig 不能为空")
	}
	if strings.TrimSpace(cfg.SDKPath) == "" {
		return "", errors.New("SDKPath 不能为空（目录或完整库文件路径）")
	}

	plat := cfg.SDKPlatform
	if plat == "" {
		// 未显式指定则按当前机子推断
		p, err := DetectCurrentPlatform()
		if err != nil {
			return "", err
		}
		plat = p
	}

	// 已经是文件
	lower := strings.ToLower(cfg.SDKPath)
	if strings.HasSuffix(lower, ".so") || strings.HasSuffix(lower, ".dll") || strings.HasSuffix(lower, ".dylib") {
		return cfg.SDKPath, nil
	}

	// 目录 → 追加默认库名
	libName, err := plat.LibFilename()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfg.SDKPath, libName), nil
}
