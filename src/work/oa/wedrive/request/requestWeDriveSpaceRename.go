package request

type RequestWeDriveSpaceRename struct {
	UserID    string `json:"userid"`
	SpaceID   string `json:"spaceid"`
	SpaceName string `json:"space_name"`
}
