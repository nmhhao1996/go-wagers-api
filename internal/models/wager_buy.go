package models

import "time"

// WagerBuy is the model for a wager buy
type WagerBuy struct {
	ID          int       `gorm:"column:id"`
	WagerID     int       `gorm:"column:wager_id"`
	BuyingPrice float64   `gorm:"column:buying_price"`
	BoughtAt    time.Time `gorm:"column:bought_at"`
}

// TableName returns the table name for the model
func (WagerBuy) TableName() string {
	return "wager_buys"
}
