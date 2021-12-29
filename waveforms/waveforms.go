package waveforms

import (
	"math"
	"math/rand"
)

func GetRandomWaveFn() func(angle float64, frequency float64) float64 {
	waveFns := []func(angle float64, frequency float64) float64 {
		SineWithDecay,
		Thud,
	}
	return waveFns[rand.Intn(len(waveFns))]
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func SineWithDecay(angle float64, frequency float64) float64 {
	if angle < math.Pi {
		return 3 * Sine(angle, 0.5*frequency)
	}
	return 3*  Sine(angle, frequency) * math.Exp(-angle/2)
}

func Additive(angle float64, frequency float64) float64 {
	return Sine(angle, frequency) +
				0.5 * Sine(angle, 5 * frequency) +
				0.3 * Square(angle, 5 * frequency) -
				Sine(angle, 100 * frequency) -
				Square(angle, frequency)
}

func Square(angle float64, frequency float64) float64 {
	if Sine(angle, frequency) >= 0 {
		return 1.0
	}
	return -1.0
}

func Thud(angle float64, _ float64) float64 { //untested
	frequency := 100*math.Exp(-angle)
	return 5*Sine(angle, frequency)
}
