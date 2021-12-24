package shared

import "math"

func FloatingPointEqual(expected float64, actual float64, threshold float64) bool {
	absDiff := math.Abs(expected - actual)
	if absDiff < threshold {
		return true
	}
	return false
}
