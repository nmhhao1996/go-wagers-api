package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/nmhhao1996/go-wagers-api/internal/models"
	"github.com/nmhhao1996/go-wagers-api/internal/wager/repository"
	"github.com/nmhhao1996/go-wagers-api/pkg/converter"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockDeps struct {
	repo *repository.MockRepository
}

func initUsecase(t *testing.T) (Usecase, mockDeps) {
	repo := repository.NewMockRepository(t)

	return New(repo), mockDeps{
		repo: repo,
	}
}

func TestWagerUsecase_Create(t *testing.T) {
	now := time.Now()

	type mockRepoCreate struct {
		expCall bool
		input   models.Wager
		output  models.Wager
		err     error
	}

	tcs := map[string]struct {
		givenInput     CreateInput
		mockRepoCreate mockRepoCreate
		wantRes        models.Wager
		wantErr        error
	}{
		"success": {
			givenInput: CreateInput{
				TotalWagerValue:   100,
				Odds:              100,
				SellingPercentage: 40,
				SellingPrice:      100,
			},
			mockRepoCreate: mockRepoCreate{
				expCall: true,
				input: models.Wager{
					TotalWagerValue:     100,
					Odds:                100,
					SellingPercentage:   40,
					SellingPrice:        100,
					CurrentSellingPrice: 100,
					PlacedAt:            now,
				},
				output: models.Wager{
					ID:                  1,
					TotalWagerValue:     100,
					Odds:                100,
					SellingPercentage:   40,
					SellingPrice:        100,
					CurrentSellingPrice: 100,
					PlacedAt:            now,
				},
			},
			wantRes: models.Wager{
				ID:                  1,
				TotalWagerValue:     100,
				Odds:                100,
				SellingPercentage:   40,
				SellingPrice:        100,
				CurrentSellingPrice: 100,
				PlacedAt:            now,
			},
		},
		"error: invalid selling price": {
			givenInput: CreateInput{
				TotalWagerValue:   100,
				Odds:              100,
				SellingPercentage: 100,
				SellingPrice:      100,
			},
			wantErr: ErrInvalidSellingPrice,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()

			orgGetTimeNowFn := getTimeNowFn
			getTimeNowFn = func() time.Time {
				return now
			}
			defer func() {
				getTimeNowFn = orgGetTimeNowFn
			}()

			uc, deps := initUsecase(t)

			if tc.mockRepoCreate.expCall {
				deps.repo.EXPECT().
					Create(ctx, tc.mockRepoCreate.input).
					Return(tc.mockRepoCreate.output, tc.mockRepoCreate.err)
			}

			// Act
			res, err := uc.Create(ctx, tc.givenInput)

			// Assert
			if tc.wantErr != nil {
				require.EqualError(t, err, tc.wantErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantRes, res)
			}
		})
	}
}

func TestWagerUsecase_Buy(t *testing.T) {
	now := time.Now()

	type mockRepoGetByID struct {
		expCall bool
		input   int
		output  models.Wager
		err     error
	}

	type mockRepoCreateBuy struct {
		expCall bool
		input   models.WagerBuy
		output  models.WagerBuy
		err     error
	}

	type mockRepoUpdate struct {
		expCall bool
		input   models.Wager
		err     error
	}

	type mockRepo struct {
		getByID   mockRepoGetByID
		createBuy mockRepoCreateBuy
		update    mockRepoUpdate
	}

	tcs := map[string]struct {
		givenInput BuyInput
		mockRepo   mockRepo
		wantRes    models.WagerBuy
		wantErr    error
	}{
		"success": {
			givenInput: BuyInput{
				WagerID:     1,
				BuyingPrice: 10,
			},
			mockRepo: mockRepo{
				getByID: mockRepoGetByID{
					expCall: true,
					input:   1,
					output: models.Wager{
						ID:                  1,
						TotalWagerValue:     1000,
						Odds:                20,
						SellingPercentage:   10,
						SellingPrice:        129,
						CurrentSellingPrice: 32,
						PercentageSold:      converter.ToPointer(75),
						AmountSold:          converter.ToPointer(3),
						PlacedAt:            time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				createBuy: mockRepoCreateBuy{
					expCall: true,
					input: models.WagerBuy{
						WagerID:     1,
						BuyingPrice: 10,
						BoughtAt:    now,
					},
					output: models.WagerBuy{
						ID:          1,
						WagerID:     1,
						BuyingPrice: 10,
						BoughtAt:    now,
					},
				},
				update: mockRepoUpdate{
					expCall: true,
					input: models.Wager{
						ID:                  1,
						TotalWagerValue:     1000,
						Odds:                20,
						SellingPercentage:   10,
						SellingPrice:        129,
						CurrentSellingPrice: 22,
						PercentageSold:      converter.ToPointer(82),
						AmountSold:          converter.ToPointer(4),
						PlacedAt:            time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantRes: models.WagerBuy{
				ID:          1,
				WagerID:     1,
				BuyingPrice: 10,
				BoughtAt:    now,
			},
		},
		"error: wager not found": {
			givenInput: BuyInput{
				WagerID:     1,
				BuyingPrice: 10,
			},
			mockRepo: mockRepo{
				getByID: mockRepoGetByID{
					expCall: true,
					input:   1,
					err:     repository.ErrNotFound,
				},
			},
			wantErr: ErrWagerNotFound,
		},
		"error: wager sold out": {
			givenInput: BuyInput{
				WagerID:     1,
				BuyingPrice: 10,
			},
			mockRepo: mockRepo{
				getByID: mockRepoGetByID{
					expCall: true,
					input:   1,
					output: models.Wager{
						ID:                  1,
						TotalWagerValue:     1000,
						Odds:                20,
						SellingPercentage:   10,
						SellingPrice:        129,
						CurrentSellingPrice: 0,
						PercentageSold:      converter.ToPointer(100),
						AmountSold:          converter.ToPointer(3),
						PlacedAt:            time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: ErrWagerSoldOut,
		},
		"error: buying price is too high": {
			givenInput: BuyInput{
				WagerID:     1,
				BuyingPrice: 100000,
			},
			mockRepo: mockRepo{
				getByID: mockRepoGetByID{
					expCall: true,
					input:   1,
					output: models.Wager{
						ID:                  1,
						TotalWagerValue:     1000,
						Odds:                20,
						SellingPercentage:   10,
						SellingPrice:        129,
						CurrentSellingPrice: 32,
						PercentageSold:      converter.ToPointer(75),
						AmountSold:          converter.ToPointer(3),
						PlacedAt:            time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: ErrBuyingPriceTooHigh,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			// Arrange
			ctx := context.Background()

			orgGetTimeNowFn := getTimeNowFn
			getTimeNowFn = func() time.Time {
				return now
			}
			defer func() {
				getTimeNowFn = orgGetTimeNowFn
			}()

			uc, deps := initUsecase(t)

			if tc.mockRepo.getByID.expCall {
				deps.repo.EXPECT().
					GetByID(mock.Anything, tc.mockRepo.getByID.input).
					Return(tc.mockRepo.getByID.output, tc.mockRepo.getByID.err)
			}

			if tc.mockRepo.createBuy.expCall {
				deps.repo.EXPECT().
					CreateBuy(mock.Anything, tc.mockRepo.createBuy.input).
					Return(tc.mockRepo.createBuy.output, tc.mockRepo.createBuy.err)
			}

			if tc.mockRepo.update.expCall {
				deps.repo.EXPECT().
					Update(mock.Anything, tc.mockRepo.update.input).
					Return(tc.mockRepo.update.err)
			}

			// Act
			res, err := uc.Buy(ctx, tc.givenInput)

			// Assert
			if tc.wantErr != nil {
				require.EqualError(t, err, tc.wantErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.wantRes, res)
			}
		})
	}
}
