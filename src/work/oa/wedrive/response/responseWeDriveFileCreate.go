package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileCreate struct {
	response.ResponseWork

	FileID string `json:"fileid"`
	Url    string `json:"url"`
}
