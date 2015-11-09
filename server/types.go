package server

//Response is a generic API response type
type Response struct {
	Type   string `json:"type"`
	Status int    `json:"status"`
	Code   string `json:"code"`
}
