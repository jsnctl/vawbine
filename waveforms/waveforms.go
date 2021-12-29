package waveforms

import (
	"math"
	"math/rand"
)

func GetRandomWaveFn() func(angle float64, frequency float64) float64 {
	waveFns := []func(angle float64, frequency float64) float64 {
		Saw,
		Sine,
	}
	return waveFns[rand.Intn(len(waveFns))]
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func Square(angle float64, frequency float64) float64 {
	if Sine(angle, frequency) >= 0 {
		return 1.0
	}
	return -1.0
}

func Saw(angle float64, frequency float64) float64 { //untested
	phase := angle * frequency
	val := 2.0*(phase*(1.0/(math.Pi))) - 1.0
	return 100 * val
}
