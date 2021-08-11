package helper

import (
	"strings"
	"time"
)

type (
	Response struct {
		Timestamp time.Time   `json:"timestamp"`
		Status    bool        `json:"status"`
		Message   string      `json:"message"`
		Errors    interface{} `json:"errors"`
		Data      interface{} `json:"data"`
	}

	MessageError struct {
		Timestamp time.Time   `json:"timestamp"`
		Errors    interface{} `json:"errors"`
		Message   interface{} `json:"message"`
	}

	Message struct {
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
	}
)

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Timestamp: time.Now(),
		Status:    status,
		Message:   message,
		Errors:    nil,
		Data:      data,
	}
	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Timestamp: time.Now(),
		Status:    false,
		Message:   message,
		Errors:    splittedError,
		Data:      data,
	}
	return res
}

func ResponseErrorMessage(message string, err string) MessageError {
	splittedError := strings.Split(err, "\n")
	res := MessageError{
		Message:   message,
		Timestamp: time.Now(),
		Errors:    splittedError,
	}
	return res
}

func ResponseMessage(message string) Message {
	res := Message{
		Message:   message,
		Timestamp: time.Now(),
	}
	return res
}

// func (r *Response) Error(err interface{}) *Response {
// 	r.Errors = &MessageError{
// 		Message: err,
// 	}
// 	return r
// }
