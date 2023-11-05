package types

type RemoteResponse struct {
	StatusCode int
	Status     string `json:"status"`
	Message    string `json:"message"`
	Items      any    `json:"return"`
}
