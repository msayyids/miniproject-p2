package controller

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func NewErrorResponse(statusCode int, message string) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}
