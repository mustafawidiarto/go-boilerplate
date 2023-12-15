package omnichannel

type GenericApiResponse[T any] struct {
	Data T `json:"data"`
	Meta *struct {
		PerPage    int `json:"per_page"`
		TotalCount int `json:"total_count"`
	} `json:"meta"`
	Status int `json:"status"`
}

type OmnichannelData struct {
	AddedAgent    *Agent        `json:"added_agent"`
	CustomerRoom  *CustomerRoom `json:"customer_room"`
	BroadcastLogs BroadcastLogs `json:"broadcast_logs"`
	WaChannel     *WaChannel    `json:"wa_channel"`
}

type OmnichannelErrors struct {
	Errors interface{} `json:"errors"`
	Status int         `json:"status"`
}

type BroadcastEstimatePriceResponse struct {
	CurrentBalance struct {
		Credits     string `json:"credits"`
		FreeSession int    `json:"free_session"`
		Quota       int    `json:"quota"`
	} `json:"current_balance"`
	EstimationCharge struct {
		Credit      int `json:"credit"`
		FreeSession int `json:"free_session"`
		Quota       int `json:"quota"`
	} `json:"estimation_charge"`
	MinimunCredits int    `json:"minimum_credits"`
	PhoneNumber    string `json:"phone_number"`
}
