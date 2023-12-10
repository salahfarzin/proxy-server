package types

type RemoteResponse struct {
	StatusCode int               `json:"statusCode"`
	Success    bool              `json:"success"`
	Message    string            `json:"message"`
	Items      any               `json:"return"`
	Errors     map[string]string `json:"errors"`
}
