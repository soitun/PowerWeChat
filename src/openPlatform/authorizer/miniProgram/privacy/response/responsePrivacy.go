package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Setting struct {
	PrivacyKey   string `json:"privacy_key"`
	PrivacyText  string `json:"privacy_text"`
	PrivacyLabel string `json:"privacy_label"`
}

type OwnerSetting struct {
	ContactPhone         string `json:"contact_phone"`
	ContactEmail         string `json:"contact_email"`
	ContactQq            string `json:"contact_qq"`
	ContactWeixin        string `json:"contact_weixin"`
	StoreExpireTimestamp string `json:"store_expire_timestamp"`
	ExtFileMediaId       string `json:"ext_file_media_id"`
	NoticeMethod         string `json:"notice_method"`
}

type Desc struct {
	PrivacyKey  string `json:"privacy_key"`
	PrivacyDesc string `json:"privacy_desc"`
}

type PrivacyDesc struct {
	PrivacyDescList []Desc `json:"privacy_desc_list"`
}

type Sdk struct {
	PrivacyKey   string `json:"privacy_key"`
	PrivacyText  string `json:"privacy_text"`
	PrivacyLabel string `json:"privacy_label"`
}

type SdkPrivacyInfo struct {
	SdkName    string `json:"sdk_name"`
	SdkBizName string `json:"sdk_biz_name"`
	SdkList    []Sdk  `json:"sdk_list"`
}

type ResponseGet struct {
	response.ResponseOpenPlatform

	CodeExist          int              `json:"code_exist"`
	PrivacyList        []string         `json:"privacy_list"`
	SettingList        []Setting        `json:"setting_list"`
	UpdateTime         int              `json:"update_time"`
	OwnerSetting       OwnerSetting     `json:"owner_setting"`
	PrivacyDesc        PrivacyDesc      `json:"privacy_desc"`
	SdkPrivacyInfoList []SdkPrivacyInfo `json:"sdk_privacy_info_list"`
}

type ResponseUpload struct {
	response.ResponseOpenPlatform

	ExtFileMediaID string `json:"ext_file_media_id"`
}
