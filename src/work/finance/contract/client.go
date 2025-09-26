package contract

import "github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/types"

type Client interface {
	GetChatData(seq uint64, limit uint32, proxy string, passwd string, timeout int) ([]*types.ResponseChatData, error)

	DecryptData(encryptRandomKey string, encryptMsg string, specificPrivateKey string) (msg *types.ChatMessage, err error)

	GetMediaData(indexBuf string, sdkFileId string, proxy string, passwd string, timeout int) (*types.MediaData, error)
}
