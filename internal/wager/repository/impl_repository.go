package repository

import (
	"context"

	"github.com/nmhhao1996/go-wagers-api/internal/core/pagination"
	"github.com/nmhhao1996/go-wagers-api/internal/models"
	"github.com/nmhhao1996/go-wagers-api/pkg/log"
	"gorm.io/gorm"
)

type implRepository struct {
	l  log.Logger
	db *gorm.DB
}

func (r implRepository) Update(ctx context.Context, m models.Wager) error {
	if m.ID <= 0 {
		return ErrNotFound
	}

	db := r.db.WithContext(ctx)
	if err := db.Save(&m).Error; err != nil {
		r.l.Errorf(ctx, "wager.repository.update.Save: %v", err)
		return convErr(err)
	}

	return nil
}

// CreateBuy implements Repository.
func (r implRepository) CreateBuy(ctx context.Context, m models.WagerBuy) (models.WagerBuy, error) {
	db := r.db.WithContext(ctx).Model(&models.WagerBuy{})

	if err := db.Create(&m).Error; err != nil {
		r.l.Errorf(ctx, "wager.repository.createBuy.Create: %v", err)
		return models.WagerBuy{}, convErr(err)
	}

	return m, nil
}

// GetByID implements Repository.
func (r implRepository) GetByID(ctx context.Context, id int) (models.Wager, error) {
	db := r.db.WithContext(ctx).Model(&models.Wager{})
	var m models.Wager

	if err := db.First(&m, id).Error; err != nil {
		r.l.Errorf(ctx, "wager.repository.getByID.First: %v", err)
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
		r.l.Errorf(ctx, "wager.repository.list.Find: %v", err)
		return nil, convErr(err)
	}

	return ms, nil
}

func (r implRepository) Create(ctx context.Context, m models.Wager) (models.Wager, error) {
	db := r.db.WithContext(ctx).Model(&models.Wager{})

	if err := db.Create(&m).Error; err != nil {
		r.l.Errorf(ctx, "wager.repository.create.Create: %v", err)
		return models.Wager{}, convErr(err)
	}

	return m, nil
}

// New creates a new wager repository
func New(l log.Logger, db *gorm.DB) Repository {
	return implRepository{
		l:  l,
		db: db,
	}
}
