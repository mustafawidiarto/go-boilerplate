package sdk

type RoomParticipants []RoomParticipant
type RoomParticipant struct {
	Active    bool        `json:"active"`
	AvatarUrl string      `json:"avatar_url"`
	Extras    interface{} `json:"extras"`
	UserID    string      `json:"user_id"`
	Username  string      `json:"username"`

	Email                    string `json:"email"`
	LastCommentReceivedIDStr string `json:"last_comment_received_id_str"`
	IDStr                    string `json:"id_str"`
	LastCommentReadIDStr     string `json:"last_comment_read_id_str"`
	ID                       int    `json:"id"`
	LastCommentReadID        int    `json:"last_comment_read_id"`
	LastCommentReceivedID    int    `json:"last_comment_received_id"`
}
