package request

type RequestWeDriveFileMove struct {
	UserID   string   `json:"userid"`
	FatherID string   `json:"fatherid"`
	Replace  bool     `json:"replace"`
	FileID   []string `json:"fileid"`
}
