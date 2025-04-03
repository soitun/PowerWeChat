package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceInfo struct {
	response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}
