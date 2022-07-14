package web

type Response struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}
