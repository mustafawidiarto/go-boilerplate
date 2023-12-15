package sdk

type WebhookBotRoomRequest struct {
	Channel        string `json:"channel"`
	ChannelDetails struct {
		ChannelID int64  `json:"channel_id"`
		Name      string `json:"name"`
		Phone     string `json:"phone"`
	} `json:"channel_details"`
	IsResolved bool   `json:"is_resolved"`
	IsWaiting  bool   `json:"is_waiting"`
	Source     string `json:"source"`
}

type SendSystemEventMessageRequest struct {
	RoomID          string      `json:"room_id"`
	SystemEventType string      `json:"system_event_type"`
	Message         string      `json:"message,omitempty"`
	SubjectEmail    string      `json:"subject_email,omitempty"`
	ObjectEmail     []string    `json:"object_email,omitempty"`
	UpdatedRoomName string      `json:"updated_room_name,omitempty"`
	Payload         interface{} `json:"payload,omitempty"`
	Extras          interface{} `json:"extras,omitempty"`
}
