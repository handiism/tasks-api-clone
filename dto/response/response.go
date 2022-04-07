package response

type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func NewResponse(message string, code int, status string, data interface{}) Response {
	return Response{
		Message: message,
		Code:    code,
		Status:  status,
		Data:    data,
	}
}
