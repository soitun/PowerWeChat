package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileMove struct {
	response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
