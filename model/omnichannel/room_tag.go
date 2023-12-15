package omnichannel

type RoomTags []RoomTag
type RoomTag struct {
	AppId          int    `json:"app_id"`
	CreatedAt      string `json:"created_at"`
	Id             int    `json:"id"`
	Name           string `json:"name"`
	RoomTagCreated string `json:"room_tag_created"`
	UpdatedAt      string `json:"updated_at"`
}
