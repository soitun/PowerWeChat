package request

type RequestAddJoinWay struct {
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseId     int      `json:"room_base_id"`
	ChatIdList     []string `json:"chat_id_list"`
	State          string   `json:"state"`
}
