package controller

type Response struct {
	Message string      `json:"message"`
	Result  interface{} `json:"detail"`
}

func NewResponse(message string, result interface{}) Response {
	return Response{
		Message: message,
		Result:  result,
	}
}
