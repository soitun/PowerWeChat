package request

type RequestWeDriveFileDownload struct {
	UserID string `json:"userid"`
	FileID string `json:"fileid"`
}
