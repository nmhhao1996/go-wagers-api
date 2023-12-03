package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/nmhhao1996/go-wagers-api/internal/models"
	"github.com/nmhhao1996/go-wagers-api/internal/wager/repository"
	"github.com/nmhhao1996/go-wagers-api/pkg/log"
)

var (
	getTimeNowFn = getTimeNow
)

func getTimeNow() time.Time {
	return time.Now()
}

type implUsecase struct {
	l    log.Logger
	repo repository.Repository
}

func (uc implUsecase) Buy(ctx context.Context, inp BuyInput) (models.WagerBuy, error) {
	w, err := uc.repo.GetByID(ctx, inp.WagerID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return models.WagerBuy{}, ErrWagerNotFound
		}

		return models.WagerBuy{}, err
	}

	if w.CurrentSellingPrice <= 0 {
		return models.WagerBuy{}, ErrWagerSoldOut
	}

	if inp.BuyingPrice > w.CurrentSellingPrice {
		return models.WagerBuy{}, ErrBuyingPriceTooHigh
	}

	wb, err := uc.repo.CreateBuy(ctx, models.WagerBuy{
		WagerID:     inp.WagerID,
		BuyingPrice: inp.BuyingPrice,
		BoughtAt:    getTimeNowFn(),
	})
	if err != nil {
		return models.WagerBuy{}, err
	}

	w.CurrentSellingPrice -= inp.BuyingPrice

	if w.AmountSold == nil {
		w.AmountSold = new(int)
	}
	*w.AmountSold += 1

	if w.PercentageSold == nil {
		w.PercentageSold = new(int)
	}
	*w.PercentageSold = int((1 - (w.CurrentSellingPrice / w.SellingPrice)) * 100)

	if err := uc.repo.Update(ctx, w); err != nil {
		return models.WagerBuy{}, err
	}

	return wb, nil

}

func (uc implUsecase) List(ctx context.Context, inp ListInput) ([]models.Wager, error) {
	ms, err := uc.repo.List(ctx, inp.PagQuery)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (uc implUsecase) validateCreateInput(inp CreateInput) error {
	if inp.SellingPrice <= float64(inp.TotalWagerValue)*(float64(inp.SellingPercentage)/100) {
		return ErrInvalidSellingPrice
	}

	return nil
}

func (uc implUsecase) Create(ctx context.Context, inp CreateInput) (models.Wager, error) {
	if err := uc.validateCreateInput(inp); err != nil {
		return models.Wager{}, err
	}

	m := models.Wager{
		TotalWagerValue:     inp.TotalWagerValue,
		Odds:                inp.Odds,
		SellingPercentage:   inp.SellingPercentage,
		SellingPrice:        inp.SellingPrice,
		CurrentSellingPrice: inp.SellingPrice,
		PlacedAt:            getTimeNowFn(),
	}

	m, err := uc.repo.Create(ctx, m)
	if err != nil {
		return models.Wager{}, err
	}

	return m, nil
}

// New creates a new wager usecase
func New(l log.Logger, repo repository.Repository) Usecase {
	return implUsecase{
		l:    l,
		repo: repo,
	}
}
