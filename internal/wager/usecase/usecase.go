package usecase

import (
	"context"

	"github.com/nmhhao1996/go-wagers-api/internal/models"
)

// Usecase is the interface for the wager usecase
type Usecase interface {
	// Create creates a wager
	Create(ctx context.Context, inp CreateInput) (models.Wager, error)
	// List lists all wagers
	List(ctx context.Context, inp ListInput) ([]models.Wager, error)
	// Buy buys a wager
	Buy(ctx context.Context, inp BuyInput) (models.WagerBuy, error)
}
