package response

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	pkgErrors "github.com/go-errors/errors"
	"github.com/nmhhao1996/go-wagers-api/pkg/log"
)

const (
	CodeOK = 0
)

const (
	MessageOK = "Success"
)

// Response is the response object
type Response struct {
	l log.Logger
}

// NewResponse creates a new response
func NewResponse(l log.Logger) Response {
	return Response{l: l}
}

func (r Response) WithOK(c *gin.Context, data any) {
	r.WithCode(c, http.StatusOK, data)
}

func (r Response) WithNoContent(c *gin.Context) {
	r.WithCode(c, http.StatusNoContent, nil)
}

func (r Response) WithCreated(c *gin.Context, data any) {
	r.WithCode(c, http.StatusCreated, data)
}

func (r Response) WithCode(c *gin.Context, code int, data any) {
	c.JSON(code, data)
}

func (r Response) WithError(c *gin.Context, err error) {
	r.withError(c, err)
}

func (r Response) WithErrorMapping(c *gin.Context, err error, mapping ErrorMapping) {
	for e, mapE := range mapping {
		if errors.Is(err, e) {
			err = mapE
			break
		}
	}

	r.WithError(c, err)
}

func (r Response) withError(c *gin.Context, err error) {
	if _, ok := err.(*Error); !ok {
		r.logError(c, err)
	}

	c.JSON(parseErrorToResponse(err))
}

func (r Response) logError(ctx context.Context, err error) {
	switch err := err.(type) {
	case *pkgErrors.Error:
		r.l.Errorf(ctx, "Error: %s", err.ErrorStack())
	default:
		r.l.Errorf(ctx, "Unwrapped error: %s", err)
	}
}

type ErrorMapping map[error]error
