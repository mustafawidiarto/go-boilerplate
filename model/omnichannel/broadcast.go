package omnichannel

type Broadcast struct {
	BroadcastJobID int64         `json:"broadcast_job_id"`
	BroadcastLogs  BroadcastLogs `json:"broadcast_logs"`
}
