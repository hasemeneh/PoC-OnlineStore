package response

import "fmt"

const JsonContentType = "application/json"

var (
	ErrorInternalServer    = NewResponseError("Internal Server Error", 500)
	ErrorForbiddenResource = NewResponseError("Internal Server Error", 403)
	ErrorBadRequest        = NewResponseError("Internal Server Error", 400)
)

type ResponseError struct {
	ErrorMessage string `json:"error,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("%s with code %d", re.ErrorMessage, re.ErrorCode)
}

func NewResponseError(message string, code int) *ResponseError {
	return &ResponseError{
		ErrorMessage: message,
		ErrorCode:    code,
	}
}
