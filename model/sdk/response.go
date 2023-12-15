package sdk

type QiscusSDKErrorResponse struct {
	Errors interface{} `json:"error"`
	Status int         `json:"status"`
}

type QiscusSDKGeneralResponse struct {
	Results Results `json:"results"`
	Status  int     `json:"status"`
}

type LoadCommentsResponse struct {
	Results Results `json:"results"`
	Status  int     `json:"status"`
}

type Results struct {
	Comments     Comments         `json:"comments"`
	Rooms        Rooms            `json:"rooms"`
	Participants RoomParticipants `json:"participants"`
}
