package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagDeleteUser struct {
	response.ResponseWork

	InvalidUser  string `json:"invalidlist"`  // "userid1|userid2",
	InvalidParty []int  `json:"invalidparty"` // "partyid1|partyid2",
}
