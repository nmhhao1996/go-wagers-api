package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nmhhao1996/go-wagers-api/internal/core/pagination"
	"github.com/nmhhao1996/go-wagers-api/internal/core/response"
	"github.com/nmhhao1996/go-wagers-api/internal/wager/usecase"
	"github.com/nmhhao1996/go-wagers-api/pkg/log"
)

type implHandler struct {
	l  log.Logger
	uc usecase.Usecase
}

func (h implHandler) Buy(c *gin.Context) {
	ctx := c.Request.Context()

	wid, err := strconv.Atoi(c.Param("wager_id"))
	if err != nil {
		h.l.Errorf(ctx, "wager.http.buy.Atoi: %v", err)
		response.WithError(c, errInvalidWagerID)
		return
	}

	var req buyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "wager.http.buy.ShouldBindJSON: %v", err)
		response.WithError(c, errInvalidRequestBody)
		return
	}

	req.adjust()

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "wager.http.buy.validate: %v", err)
		response.WithError(c, err)
		return
	}

	m, err := h.uc.Buy(ctx, usecase.BuyInput{
		WagerID:     wid,
		BuyingPrice: req.BuyingPrice,
	})
	if err != nil {
		h.l.Errorf(ctx, "wager.http.buy.usecase: %v", err)
		response.WithErrorMapping(c, err, errMapping)
		return
	}

	response.WithCreated(c, newBuyResponse(m))
}

func (h implHandler) List(c *gin.Context) {
	ctx := c.Request.Context()

	pag, err := pagination.GetPaginationQueryFromContext(c)
	if err != nil {
		h.l.Errorf(ctx, "wager.http.list.GetPaginationQueryFromContext: %v", err)
		response.WithError(c, errInvalidPaginationQuery)
		return
	}

	ms, err := h.uc.List(ctx, usecase.ListInput{
		PagQuery: pag,
	})
	if err != nil {
		h.l.Errorf(ctx, "wager.http.uc.List: %v", err)
		response.WithErrorMapping(c, err, errMapping)
		return
	}

	response.WithOK(c, newListResponse(ms))
}

func (h implHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Errorf(ctx, "wager.http.create.ShouldBindJSON: %v", err)
		response.WithError(c, errInvalidRequestBody)
		return
	}

	req.adjust()

	if err := req.validate(); err != nil {
		h.l.Errorf(ctx, "wager.http.create.validate: %v", err)
		response.WithError(c, err)
		return
	}

	m, err := h.uc.Create(ctx, req.toInput())
	if err != nil {
		h.l.Errorf(ctx, "wager.http.create.usecase: %v", err)
		response.WithErrorMapping(c, err, errMapping)
		return
	}

	response.WithCreated(c, newCreateResponse(m))
}

// New creates a new wager http handler
func New(l log.Logger, uc usecase.Usecase) Handler {
	return implHandler{
		l:  l,
		uc: uc,
	}
}
