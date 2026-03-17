package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	response.ResponseWork
	ChatID string `json:"chat_id"`
}
