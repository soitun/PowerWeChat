package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileRename struct {
	response.ResponseWork

	File *power.HashMap `json:"file"`
}
