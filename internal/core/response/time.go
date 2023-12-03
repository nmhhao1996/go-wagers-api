package response

import (
	"encoding/json"
	"time"
)

// Date Response
func marshalTime(dt time.Time, format string) ([]byte, error) {
	t := dt.Format(format)
	dstr, err := json.Marshal(t)
	if err != nil {
		return []byte{}, err
	}

	return dstr, nil
}

// DateResponse is a custom time.Time type that marshals to and from JSON in YYYY-MM-DD format
type DateResponse time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d DateResponse) MarshalJSON() ([]byte, error) {
	return marshalTime(time.Time(d), time.DateOnly)
}

// DateTimeResponse is a custom time.Time type that marshals to and from JSON in YYYY-MM-DD HH:MM:SS format
type DateTimeResponse time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d DateTimeResponse) MarshalJSON() ([]byte, error) {
	return marshalTime(time.Time(d), time.DateTime)
}

type TimestampResponse time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d TimestampResponse) MarshalJSON() ([]byte, error) {
	t := time.Time(d).Unix()
	dstr, err := json.Marshal(t)
	if err != nil {
		return []byte{}, err
	}

	return dstr, nil
}
