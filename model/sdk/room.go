package sdk

type Rooms []Room
type Room struct {
	RoomAvatarUrl string `json:"room_avatar_url"`
	RoomChannelID string `json:"room_channel_id"`
	RoomID        string `json:"room_id"`
	RoomName      string `json:"room_name"`
	RoomOptions   string `json:"room_options"`
	RoomType      string `json:"room_type"`
}


