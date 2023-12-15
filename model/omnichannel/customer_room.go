package omnichannel

type CustomerRoom struct {
	ID             int64  `json:"id"`
	ChannelID      int64  `json:"channel_id"`
	IsResolved     bool   `json:"is_resolved"`
	IsWaiting      bool   `json:"is_Waiting"`
	Name           string `json:"name"`
	RoomID         string `json:"room_id"`
	Source         string `json:"source"`
	UserID         string `json:"user_id"`
	UserAvatarUrl  string `json:"user_avatar_url"`
	IsHandledByBot bool   `json:"is_handled_by_bot"`
}
