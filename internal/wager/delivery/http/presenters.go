package http

import (
	"github.com/nmhhao1996/go-wagers-api/internal/core/response"
	"github.com/nmhhao1996/go-wagers-api/internal/models"
	"github.com/nmhhao1996/go-wagers-api/internal/wager/usecase"
	"github.com/nmhhao1996/go-wagers-api/pkg/num"
)

const (
	pricePrecision = 2
)

type createRequest struct {
	TotalWagerValue   int     `json:"total_wager_value"`
	Odds              int     `json:"odds"`
	SellingPercentage int     `json:"selling_percentage"`
	SellingPrice      float64 `json:"selling_price"`
}

func (r *createRequest) adjust() {
	r.SellingPrice = num.RoundFloat64(r.SellingPrice, pricePrecision)
}

func (r createRequest) validate() error {
	if r.TotalWagerValue <= 0 {
		return errInvalidTotalWagerValue
	}

	if r.Odds <= 0 {
		return errInvalidOdds
	}

	if r.SellingPercentage <= 0 || r.SellingPercentage > 100 {
		return errInvalidSellingPercentage
	}

	return nil
}

func (r createRequest) toInput() usecase.CreateInput {
	return usecase.CreateInput(r)
}

type wagerResponse struct {
	ID                  int                        `json:"id"`
	TotalWagerValue     int                        `json:"total_wager_value"`
	Odds                int                        `json:"odds"`
	SellingPercentage   int                        `json:"selling_percentage"`
	SellingPrice        float64                    `json:"selling_price"`
	CurrentSellingPrice float64                    `json:"current_selling_price"`
	PercentageSold      *int                       `json:"percentage_sold"`
	AmountSold          *int                       `json:"amount_sold"`
	PlacedAt            response.TimestampResponse `json:"placed_at"`
}

func newWagerResponse(m models.Wager) wagerResponse {
	return wagerResponse{
		ID:                  m.ID,
		TotalWagerValue:     m.TotalWagerValue,
		Odds:                m.Odds,
		SellingPercentage:   m.SellingPercentage,
		SellingPrice:        m.SellingPrice,
		CurrentSellingPrice: m.CurrentSellingPrice,
		PercentageSold:      m.PercentageSold,
		AmountSold:          m.AmountSold,
		PlacedAt:            response.TimestampResponse(m.PlacedAt),
	}
}

type createResponse = wagerResponse

func newCreateResponse(m models.Wager) createResponse {
	return newWagerResponse(m)
}

type listResponse = []wagerResponse

func newListResponse(ms []models.Wager) listResponse {
	res := make([]wagerResponse, len(ms))

	for i := range ms {
		res[i] = newWagerResponse(ms[i])
	}

	return res
}

type buyRequest struct {
	BuyingPrice float64 `json:"buying_price"`
}

func (r *buyRequest) adjust() {
	r.BuyingPrice = num.RoundFloat64(r.BuyingPrice, pricePrecision)
}

func (r buyRequest) validate() error {
	if r.BuyingPrice <= 0 {
		return errInvalidBuyingPrice
	}

	return nil
}

type buyResponse struct {
	ID          int                        `json:"id"`
	WagerID     int                        `json:"wager_id"`
	BuyingPrice float64                    `json:"buying_price"`
	BoughtAt    response.TimestampResponse `json:"bought_at"`
}

func newBuyResponse(m models.WagerBuy) buyResponse {
	return buyResponse{
		ID:          m.ID,
		WagerID:     m.WagerID,
		BuyingPrice: m.BuyingPrice,
		BoughtAt:    response.TimestampResponse(m.BoughtAt),
	}
}
