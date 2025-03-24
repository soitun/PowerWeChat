package request

type RequestExternalContactRemark struct {
	UserID           string   `json:"userid,omitempty"`
	ExternalUserID   string   `json:"external_userid,omitempty"`
	Remark           string   `json:"remark,omitempty"`
	Description      string   `json:"description,omitempty"`
	RemarkCompany    string   `json:"remark_company,omitempty"`
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid,omitempty"`
}
