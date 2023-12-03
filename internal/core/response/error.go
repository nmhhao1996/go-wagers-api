package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	defaultErrMessage = "SOMETHING_WENT_WRONG"
)

type Error struct {
	httpCode int
	msg      string
}

// NewError creates a new error
func (e Error) Error() string {
	return e.msg
}

func NewError(httpCode int, msg string) *Error {
	return &Error{
		msg:      msg,
		httpCode: httpCode,
	}
}

func parseErrorToResponse(err error) (int, any) {
	switch err := err.(type) {
	case *Error:
		return err.httpCode, gin.H{
			"error": err.msg,
		}
	default:
		return http.StatusInternalServerError, gin.H{
			"error": defaultErrMessage,
		}
	}
}
