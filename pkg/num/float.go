package num

import "math"

// RoundFloat64 rounds a float64 to a given precision
func RoundFloat64(n float64, prec int) float64 {
	ratio := math.Pow(10, float64(prec))
	return math.Round(n*ratio) / ratio
}
