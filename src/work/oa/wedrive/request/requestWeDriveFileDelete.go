package request

type RequestWeDriveFileDelete struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}
