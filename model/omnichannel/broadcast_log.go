package omnichannel

type BroadcastLogs []*BroadcastLog
type BroadcastLog struct {
	ID          int64       `json:"id"`
	MessageID   string      `json:"message_id"`
	Notes       string      `json:"notes"`
	PhoneNumber string      `json:"phone_number"`
	SentAt      string      `json:"sent_at"`
	Status      interface{} `json:"status"`
	Variables   string      `json:"variables"`
}
