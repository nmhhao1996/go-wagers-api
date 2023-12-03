package response

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUtils_Time_DateResponse_MarshallJSON(t *testing.T) {
	type input struct {
		Date DateResponse `json:"date"`
	}

	tcs := map[string]struct {
		givenInput input
		expResult  string
		expErr     error
	}{
		"success": {
			givenInput: input{
				Date: DateResponse(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expResult: `
			{
				"date": "2020-01-01"
			}
			`,
		},
		"with time success": {
			givenInput: input{
				Date: DateResponse(time.Date(2020, 1, 1, 10, 20, 30, 30, time.UTC)),
			},
			expResult: `
			{
				"date": "2020-01-01"
			}
			`,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			res, err := json.Marshal(tc.givenInput)

			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.JSONEq(t, tc.expResult, string(res))
			}
		})
	}
}

func TestUtils_Time_DateTimeResponse_MarshallJSON(t *testing.T) {
	type input struct {
		DateTime DateTimeResponse `json:"date_time"`
	}

	tcs := map[string]struct {
		givenInput input
		expResult  string
		expErr     error
	}{
		"success": {
			givenInput: input{
				DateTime: DateTimeResponse(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expResult: `
			{
				"date_time": "2020-01-01 00:00:00"
			}
			`,
		},
		"with time success": {
			givenInput: input{
				DateTime: DateTimeResponse(time.Date(2020, 1, 1, 10, 20, 30, 30, time.UTC)),
			},
			expResult: `
			{
				"date_time": "2020-01-01 10:20:30"
			}
			`,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			res, err := json.Marshal(tc.givenInput)

			if tc.expErr != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.JSONEq(t, tc.expResult, string(res))
			}
		})
	}
}
