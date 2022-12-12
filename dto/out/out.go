package out

type DefaultResponse struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

type ErrorPayload struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}
