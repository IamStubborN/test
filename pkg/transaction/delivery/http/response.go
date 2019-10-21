package http

type response struct {
	Error string `json:"error"`
	Balance float64 `json:"balance"`
}
