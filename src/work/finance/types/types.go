package types

import "fmt"

type SDKPlatform string

const (
	// Linux
	SDKPlatformLinuxAMD64 SDKPlatform = "linux_amd64"
	SDKPlatformLinuxARM64 SDKPlatform = "linux_arm64"

	// Windows（常见为 amd64）
	SDKPlatformWindowsAMD64 SDKPlatform = "windows_amd64"

	// macOS（有 .dylib 才能用；官方通常不发）
	SDKPlatformDarwinARM64 SDKPlatform = "darwin_arm64"
)

// lib 文件名
func (p SDKPlatform) LibFilename() (string, error) {
	switch p {
	case SDKPlatformLinuxAMD64, SDKPlatformLinuxARM64:
		return "libWeWorkFinanceSdk_C.so", nil
	case SDKPlatformWindowsAMD64:
		return "WeWorkFinanceSdk_C.dll", nil
	case SDKPlatformDarwinARM64:
		return "libWeWorkFinanceSdk_C.so", nil
	default:
		return "", fmt.Errorf("未支持的平台: %s", p)
	}
}

type ChatData struct {
	Seq              int    `json:"seq"`
	MsgId            string `json:"msgid"`
	PublicKeyVer     int    `json:"publickey_ver"`
	EncryptRandomKey string `json:"encrypt_random_key"`
	EncryptChatMsg   string `json:"encrypt_chat_msg"`
}

type ChatMessage struct {
	Id          string   `json:"id"`
	From        string   `json:"from"`
	ToList      []string `json:"toList"`
	Action      string   `json:"action"`
	Type        string   `json:"type"`
	DecryptData []byte   `json:"decryptData"`
}

type MediaData struct {
	OutIndexBuf string `json:"outindexbuf,omitempty"`
	IsFinish    bool   `json:"is_finish,omitempty"`
	Data        []byte `json:"data,omitempty"`
}
