package models

import "time"

// Wager is the model for a wager
type Wager struct {
	ID                  int       `gorm:"primaryKey"`
	TotalWagerValue     int       `gorm:"column:total_wager_value"`
	Odds                int       `gorm:"column:odds"`
	SellingPercentage   int       `gorm:"column:selling_percentage"`
	SellingPrice        float64   `gorm:"column:selling_price"`
	CurrentSellingPrice float64   `gorm:"column:current_selling_price"`
	PlacedAt            time.Time `gorm:"column:placed_at"`
	PercentageSold      *int      `gorm:"column:percentage_sold"`
	AmountSold          *int      `gorm:"column:amount_sold"`
}

// TableName returns the table name for the model
func (Wager) TableName() string {
	return "wagers"
}
