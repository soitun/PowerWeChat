package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveSpaceCreate struct {
	UserID    string           `json:"userid"`
	SpaceName string           `json:"space_name"`
	AuthInfo  []*power.HashMap `json:"auth_info"`
}
