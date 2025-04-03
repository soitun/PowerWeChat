package request

type RequestWeDriveFileSetting struct {
	UserID    string `json:"userid"`
	FileID    string `json:"fileid"`
	AuthScope int    `json:"auth_scope"`
	Auth      int    `json:"auth"`
}
