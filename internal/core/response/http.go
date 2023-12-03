package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeOK = 0
)

const (
	MessageOK = "Success"
)

func WithOK(c *gin.Context, data any) {
	WithCode(c, http.StatusOK, data)
}

func WithNoContent(c *gin.Context) {
	WithCode(c, http.StatusNoContent, nil)
}

func WithCreated(c *gin.Context, data any) {
	WithCode(c, http.StatusCreated, data)
}

func WithCode(c *gin.Context, code int, data any) {
	c.JSON(code, data)
}

func WithError(c *gin.Context, err error) {
	c.JSON(parseErrorToResponse(err))
}

func WithErrorMapping(c *gin.Context, err error, mapping ErrorMapping) {
	if e, ok := mapping[err]; ok {
		err = e
	}

	c.JSON(parseErrorToResponse(err))
}

type ErrorMapping map[error]error
