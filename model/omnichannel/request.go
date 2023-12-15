package omnichannel

type SendMessageAsBotRequest struct {
	SenderEmail string `json:"sender_email"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	RoomID      string `json:"room_id"`
}