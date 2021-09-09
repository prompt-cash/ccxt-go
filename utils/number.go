package utils

import "math"

// Get nr with only numDecimals (rounding the rest).
func GetNumberDecimals(nr float32, numDecimals int) float32 {
	shift := math.Pow(10, float64(numDecimals))
	return float32(math.Round(float64(nr)*shift) / shift)
}

// Returns the % difference between value1 and value2.
// The % difference > 0 if value1 > value2, < 0 otherwise
func GetDiffPercent(value1 float32, value2 float32) float32 {
	if value2 == 0.0 {
		return 0.0
	}
	return ((value1 - value2) / value2) * 100.0 // ((y2 - y1) / y1)*100 - positive % if value1 > value2
}
