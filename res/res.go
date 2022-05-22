package res

import (
	"strings"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type EmptyObj struct{}

func BuildDataResponse(status int, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
	return res
}

func BuildErrorResponse(status int, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  splittedError,
	}
	return res
}
