package repository

import (
	"context"

	pkgErrors "github.com/go-errors/errors"
	"github.com/nmhhao1996/go-wagers-api/internal/core/pagination"
	"github.com/nmhhao1996/go-wagers-api/internal/models"
	"gorm.io/gorm"
)

type implRepository struct {
	db *gorm.DB
}

func (r implRepository) Update(ctx context.Context, m models.Wager) error {
	if m.ID <= 0 {
		return pkgErrors.New(ErrNotFound)
	}

	db := r.db.WithContext(ctx)
	if err := db.Save(&m).Error; err != nil {
		return convErr(err)
	}

	return nil
}

// CreateBuy implements Repository.
func (r implRepository) CreateBuy(ctx context.Context, m models.WagerBuy) (models.WagerBuy, error) {
	db := r.db.WithContext(ctx).Model(&models.WagerBuy{})

	if err := db.Create(&m).Error; err != nil {
		return models.WagerBuy{}, convErr(err)
	}

	return m, nil
}

// GetByID implements Repository.
func (r implRepository) GetByID(ctx context.Context, id int) (models.Wager, error) {
	db := r.db.WithContext(ctx).Model(&models.Wager{})
	var m models.Wager

	if err := db.First(&m, id).Error; err != nil {
		return models.Wager{}, convErr(err)
	}

	return m, nil
}

func (r implRepository) List(ctx context.Context, pagQuery pagination.Query) ([]models.Wager, error) {
	db := r.db.WithContext(ctx).Model(&models.Wager{})

	var ms []models.Wager
	if err := db.
		Order("placed_at DESC").
		Limit(pagQuery.Limit).
		Offset(pagQuery.Offset()).
		Find(&ms).Error; err != nil {
		return nil, convErr(err)
	}

	return ms, nil
}

func (r implRepository) Create(ctx context.Context, m models.Wager) (models.Wager, error) {
	db := r.db.WithContext(ctx).Model(&models.Wager{})

	if err := db.Create(&m).Error; err != nil {
		return models.Wager{}, convErr(err)
	}

	return m, nil
}

// New creates a new wager repository
func New(db *gorm.DB) Repository {
	return implRepository{
		db: db,
	}
}
