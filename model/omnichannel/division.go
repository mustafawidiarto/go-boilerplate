package omnichannel

type Divisions []Division
type Division struct {
	ID            int64  `json:"id"`
	IsDefaultRole bool   `json:"is_default_role"`
	Name          string `json:"name"`
}
