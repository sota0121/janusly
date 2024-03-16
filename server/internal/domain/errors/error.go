package errors

import "fmt"

type JanuslyError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *JanuslyError) Error() string {
	errMsg := fmt.Sprintf("[%s] %s", e.Code, e.Message)
	return errMsg
}
