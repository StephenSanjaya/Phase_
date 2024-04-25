package utils

import "fmt"

type HTTPError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Detail  error  `json:"error"`
}

func NewHTTPError(code int, message string, detail error) *HTTPError {
	return &HTTPError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("%d: %v", e.Code, e.Detail)
}
