package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileShare struct {
	response.ResponseWork

	ShareURL string `json:"share_url"`
}
