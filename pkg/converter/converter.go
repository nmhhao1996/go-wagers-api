package converter

import (
	"bytes"
	"encoding/json"
)

// Convert bytes to buffer helper
func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

// Convert bytes to any
func BytesToAny(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// Convert any to bytes
func AnyToBytes(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Convert any to string
func MapStringToSlice(maps map[string]any) []any {
	var results []any
	for key, value := range maps {
		results = append(results, key, value)
	}
	return results
}

// ToPointer convert any to pointer
func ToPointer[T any](v T) *T {
	return &v
}
