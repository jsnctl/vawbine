package waveforms

import (
	"math"
)

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func Square(angle float64, frequency float64) float64 {
	if Sine(angle, frequency) >= 0 {
		return 1.0
	}
	return -1.0
}
