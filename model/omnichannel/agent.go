package omnichannel

type Agents []Agent
type Agent struct {
	ID                   int64  `json:"id"`
	Email                string `json:"email"`
	Name                 string `json:"name"`
	CurrentCustomerCount int    `json:"current_customer_count"`
	IsAvailable          bool   `json:"is_available"`
}
