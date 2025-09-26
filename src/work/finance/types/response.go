package types

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseChatData struct {
	response.ResponseWork
	ChatData []ChatData `json:"chatdata"`
}

type Text struct {
	Content string `json:"content"`
}

type Image struct {
	Md5Sum    string `json:"md5sum"`
	Filesize  int    `json:"filesize"`
	SdkFileId string `json:"sdkfileid"`
}

type Revoke struct {
	PreMsgId string `json:"pre_msgid"`
}
type Disagree struct {
	UserId       string `json:"userid"`
	DisagreeTime int64  `json:"disagree_time"`
}

type Agree struct {
	Userid    string `json:"userid"`
	AgreeTime int64  `json:"agree_time"`
}

type Voice struct {
	Md5Sum     string `json:"md5sum"`
	VoiceSize  int    `json:"voice_size"`
	PlayLength int    `json:"play_length"`
	SdkFileId  string `json:"sdkfileid"`
}

type Video struct {
	Md5Sum     string `json:"md5sum"`
	Filesize   int    `json:"filesize"`
	PlayLength int    `json:"play_length"`
	SdkFileId  string `json:"sdkfileid"`
}

type Card struct {
	CorpName string `json:"corpname"`
	Userid   string `json:"userid"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Address   string  `json:"address"`
	Title     string  `json:"title"`
	Zoom      int     `json:"zoom"`
}

type Emotion struct {
	Type      int    `json:"type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	ImageSize int    `json:"imagesize"`
	Md5Sum    string `json:"md5sum"`
	SdkFileId string `json:"sdkfileid"`
}

type File struct {
	Md5Sum    string `json:"md5sum"`
	Filename  string `json:"filename"`
	FileExt   string `json:"fileext"`
	Filesize  int    `json:"filesize"`
	SdkFileId string `json:"sdkfileid"`
}

type Link struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkUrl     string `json:"link_url"`
	ImageUrl    string `json:"image_url"`
}

type Weapp struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Username    string `json:"username"`
	DisplayName string `json:"displayname"`
}

type ChatRecordItem struct {
	Type         string `json:"type"`
	MsgTime      int    `json:"msgtime"`
	Content      string `json:"content"`
	FromChatroom bool   `json:"from_chatroom"`
}

type ChatRecord struct {
	Title string           `json:"title"`
	Item  []ChatRecordItem `json:"item"`
}

type CollectDetail struct {
	Id   int    `json:"id"`
	Ques string `json:"ques"`
	Type string `json:"type"`
}

type Collect struct {
	RoomName   string          `json:"room_name"`
	Creator    string          `json:"creator"`
	CreateTime string          `json:"create_time"`
	Title      string          `json:"title"`
	Details    []CollectDetail `json:"details"`
}

type RedPacket struct {
	Type        int    `json:"type"`
	Wish        string `json:"wish"`
	TotalCnt    int    `json:"totalcnt"`
	TotalAmount int    `json:"totalamount"`
}

type Meeting struct {
	Topic     string `json:"topic"`
	StartTime int    `json:"starttime"`
	EndTime   int    `json:"endtime"`
	Address   string `json:"address"`
	Remarks   string `json:"remarks"`
	MeetingId int    `json:"meetingid"`
}

type ChatInfo struct {
	Content          string `json:"content"`
	MeetingId        int64  `json:"meeting_id"`
	NotificationType int    `json:"notification_type"`
}

type Doc struct {
	Title      string `json:"title"`
	DocCreator string `json:"doc_creator"`
	LinkUrl    string `json:"link_url"`
}

type ChatDataMeta struct {
	MsgId   string   `json:"msgid"`
	Action  string   `json:"action"`
	From    string   `json:"from"`
	Tolist  []string `json:"tolist"`
	RoomId  string   `json:"roomid"`
	MsgTime int64    `json:"msgtime"`
	MsgType string   `json:"msgtype"`
}

type ResponseChatDataText struct {
	ChatDataMeta
	Text Text `json:"text,omitempty"`
}
type ResponseChatDataImage struct {
	ChatDataMeta
	Image Image `json:"image,omitempty"`
}
type ResponseChatDataRevoke struct {
	ChatDataMeta
	Revoke Revoke `json:"revoke,omitempty"`
}
type ResponseChatDataDisagree struct {
	ChatDataMeta
	Disagree Disagree `json:"disagree,omitempty"`
}
type ResponseChatDataAgree struct {
	ChatDataMeta
	Agree Agree `json:"agree,omitempty"`
}
type ResponseChatDataVoice struct {
	ChatDataMeta
	Voice Voice `json:"voice,omitempty"`
}
type ResponseChatDataVideo struct {
	ChatDataMeta
	Video Video `json:"video,omitempty"`
}
type ResponseChatDataCard struct {
	ChatDataMeta
	Card Card `json:"card,omitempty"`
}
type ResponseChatDataLocation struct {
	ChatDataMeta
	Location Location `json:"location,omitempty"`
}
type ResponseChatDataEmotion struct {
	ChatDataMeta
	Emotion Emotion `json:"emotion,omitempty"`
}
type ResponseChatDataFile struct {
	ChatDataMeta
	File File `json:"file,omitempty"`
}
type ResponseChatDataLink struct {
	ChatDataMeta
	Link Link `json:"link,omitempty"`
}
type ResponseChatDataWeApp struct {
	ChatDataMeta
	WeApp Weapp `json:"weapp,omitempty"`
}
type ResponseChatDataChatRecord struct {
	ChatDataMeta
	ChatRecord ChatRecord `json:"chatrecord,omitempty"`
}
type ResponseChatDataCollect struct {
	ChatDataMeta
	Collect Collect `json:"collect,omitempty"`
}
type ResponseChatDataRedPacket struct {
	ChatDataMeta
	RedPacket RedPacket `json:"redpacket,omitempty"`
}
type ResponseChatDataMeeting struct {
	ChatDataMeta
	Meeting Meeting `json:"meeting,omitempty"`
}

type MeetingNotification struct {
	Content          string `json:"content"`
	MeetingId        int64  `json:"meeting_id"`
	NotificationType int    `json:"notification_type"`
}

type ResponseChatDataMeetingNotification struct {
	ChatDataMeta
	Info MeetingNotification `json:"info"`
}

type ResponseChatDataDoc struct {
	ChatDataMeta
	Doc Doc `json:"doc,omitempty"`
}

type MarkdownInfo struct {
	Markdown string `json:"markdown"`
}

type ResponseChatDataMarkdown struct {
	ChatDataMeta
	MarkdownInfo MarkdownInfo `json:"info,omitempty"`
}

type NewsInfoItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}

type NewsInfo struct {
	Item []NewsInfoItem `json:"item"`
}

type ResponseChatDataNews struct {
	ChatDataMeta
	Info NewsInfo `json:"info"`
}

type Calendar struct {
	Title        string   `json:"title"`
	CreatorName  string   `json:"creatorname"`
	AttendeeName []string `json:"attendeename"`
	StartTime    int      `json:"starttime"`
	EndTime      int      `json:"endtime"`
	Place        string   `json:"place"`
	Remarks      string   `json:"remarks"`
}

type ResponseChatDataCalendar struct {
	ChatDataMeta
	Calendar Calendar `json:"calendar"`
}

type MixedItem struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Mixed struct {
	Item []MixedItem `json:"item"`
}

type ResponseChatDataMixed struct {
	ChatDataMeta
	Mixed Mixed `json:"mixed"`
}

type DemoFileData struct {
	Filename     string `json:"filename"`
	DemoOperator string `json:"demooperator"`
	StartTime    int    `json:"starttime"`
	EndTime      int    `json:"endtime"`
}

type ShareScreenData struct {
	Share     string `json:"share"`
	StartTime int    `json:"starttime"`
	EndTime   int    `json:"endtime"`
}

type MeetingVoiceCall struct {
	EndTime         int               `json:"endtime"`
	SdkFileId       string            `json:"sdkfileid"`
	DemoFileData    []DemoFileData    `json:"demofiledata"`
	ShareScreenData []ShareScreenData `json:"sharescreendata"`
}

type ResponseChatDataMeetingVoiceCall struct {
	ChatDataMeta
	MeetingVoiceCall MeetingVoiceCall `json:"meeting_voice_call"`
}

type VoipDocShare struct {
	Filename  string `json:"filename"`
	Md5Sum    string `json:"md5sum"`
	Filesize  int    `json:"filesize"`
	SdkFileId string `json:"sdkfileid"`
}

type ResponseChatDataVoipDocShare struct {
	ChatDataMeta
	VoipDocShare VoipDocShare `json:"voip_doc_share"`
}

type ExternalRedPacket struct {
	Type        int    `json:"type"`
	Wish        string `json:"wish"`
	TotalCnt    int    `json:"totalcnt"`
	TotalAmount int    `json:"totalamount"`
}

type ResponseChatDataExternalRedPacket struct {
	ChatDataMeta
	RedPacket ExternalRedPacket `json:"redpacket"`
}

type SphFeed struct {
	FeedType int    `json:"feed_type"`
	SphName  string `json:"sph_name"`
	FeedDesc string `json:"feed_desc"`
}

type ResponseChatDataSphFeed struct {
	ChatDataMeta
	SphFeed SphFeed `json:"sphfeed"`
}

type VoipTextInfo struct {
	CallDuration int `json:"callduration"`
	InviteType   int `json:"invitetype"`
}

type ResponseChatDataVoipText struct {
	ChatDataMeta
	Info VoipTextInfo `json:"info"`
}

type QyDiskFile struct {
	Filename string `json:"filename"`
}

type ResponseChatDataQyDiskFile struct {
	ChatDataMeta
	Info QyDiskFile `json:"info"`
}

type SwitchCorpLog struct {
	MsgId  string `json:"msgid"`
	Action string `json:"action"`
	Time   int64  `json:"time"`
	User   string `json:"user"`
}

type RobotInfo struct {
	RobotId       string `json:"robot_id"`
	Name          string `json:"name"`
	CreatorUserid string `json:"creator_userid"`
}

type ResponseRobotInfo struct {
	Data RobotInfo `json:"data"`
}
