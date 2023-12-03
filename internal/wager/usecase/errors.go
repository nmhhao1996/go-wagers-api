package usecase

import "errors"

var (
	ErrInvalidSellingPrice = errors.New("invalid selling price")
	ErrWagerNotFound       = errors.New("wager not found")
	ErrBuyingPriceTooHigh  = errors.New("buying price too high")
	ErrWagerSoldOut        = errors.New("wager sold out")
)
