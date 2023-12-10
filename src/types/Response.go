package types

type RemoteResponse struct {
	StatusCode int
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Items      any    `json:"return"`
}
