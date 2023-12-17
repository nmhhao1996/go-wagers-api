package http

import (
	"context"

	"github.com/nmhhao1996/go-wagers-api/internal/core/response"
	"github.com/nmhhao1996/go-wagers-api/internal/wager/usecase"
)

var (
	errInvalidPaginationQuery   = response.NewError(400, "INVALID_PAGINATION_QUERY")
	errInvalidRequestBody       = response.NewError(400, "INVALID_REQUEST_BODY")
	errInvalidTotalWagerValue   = response.NewError(400, "INVALID_TOTAL_WAGER_VALUE")
	errInvalidOdds              = response.NewError(400, "INVALID_ODDS")
	errInvalidSellingPercentage = response.NewError(400, "INVALID_SELLING_PERCENTAGE")
	errInvalidWagerID           = response.NewError(400, "INVALID_WAGER_ID")
	errInvalidBuyingPrice       = response.NewError(400, "INVALID_BUYING_PRICE")
	errBuyingPriceTooHigh       = response.NewError(400, "BUYING_PRICE_TOO_HIGH")
	errWagerNotFound            = response.NewError(400, "WAGER_NOT_FOUND")
	errWagerSoldOut             = response.NewError(400, "WAGER_SOLD_OUT")
	errBuyTimeOut               = response.NewError(400, "BUY_TIME_OUT")
)

var errMapping = response.ErrorMapping{
	usecase.ErrInvalidSellingPrice: errInvalidTotalWagerValue,
	usecase.ErrBuyingPriceTooHigh:  errBuyingPriceTooHigh,
	usecase.ErrWagerNotFound:       errWagerNotFound,
	usecase.ErrWagerSoldOut:        errWagerSoldOut,
	context.DeadlineExceeded:       errBuyTimeOut,
}
