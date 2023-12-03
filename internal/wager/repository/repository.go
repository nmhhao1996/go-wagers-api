package repository

import (
	"context"

	"github.com/nmhhao1996/go-wagers-api/internal/core/pagination"
	"github.com/nmhhao1996/go-wagers-api/internal/models"
)

// Repository is the interface for the wager repository
//
//go:generate mockery --name Repository --inpackage --with-expecter
type Repository interface {
	// Create creates a wager
	Create(ctx context.Context, m models.Wager) (models.Wager, error)
	// List lists all wagers
	List(ctx context.Context, pagQuery pagination.Query) ([]models.Wager, error)
	// CreateBuy creates a wager buy
	CreateBuy(ctx context.Context, m models.WagerBuy) (models.WagerBuy, error)
	// GetByID gets a wager by ID
	GetByID(ctx context.Context, id int) (models.Wager, error)
	// Update updates a wager
	Update(ctx context.Context, m models.Wager) error
}
