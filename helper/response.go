package helper

import (
	"strings"
)

// Response for static shape json return
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Empty object used for data that doesn't want to be null on json
type EmptyObj struct{}

// BuildResponse to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) (res Response) {
	res = Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}

	return
}

// BuildErrorResponse to inject data value to dynamic fail response
func BuildErrorResponse(message string, err string, data interface{}) (res Response) {
	splittedError := strings.Split(err, "\n")
	res = Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}

	return
}
