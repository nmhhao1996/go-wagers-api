package usecase

import "github.com/nmhhao1996/go-wagers-api/internal/core/pagination"

// CreateInput is the input for the Create method
type CreateInput struct {
	TotalWagerValue   int
	Odds              int
	SellingPercentage int
	SellingPrice      float64
}

// ListInput is the input for the List method
type ListInput struct {
	PagQuery pagination.Query
}

// BuyInput is the input for the Buy method
type BuyInput struct {
	WagerID     int
	BuyingPrice float64
}
