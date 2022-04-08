package response

import "github.com/go-playground/validator/v10"

type response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewErrReponse(code int, message string, err error) response {
	errors := errors(err)
	return response{
		Code:    code,
		Status:  "error",
		Message: message,
		Errors:  errors,
		Data:    nil,
	}
}

func NewReponse(code int, message string, data interface{}) response {
	return response{
		Code:    code,
		Status:  "success",
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func errors(err error) []string {
	var errors []string

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, e.Error())
		}
		return errors
	}

	errors = append(errors, err.Error())
	return errors
}
