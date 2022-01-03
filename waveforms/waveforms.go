package waveforms

import (
	"math"
	"math/rand"
)

func GetRandomWaveFn() func(angle float64, frequency float64) float64 {
	waveFns := []func(angle float64, frequency float64) float64{
		SquareWithDecay,
		SineWithDecay,
	}
	return waveFns[rand.Intn(len(waveFns))]
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func SineWithDecay(angle float64, frequency float64) float64 {
	return Sine(angle, frequency) * math.Exp(-angle/5)
}

func Torricelli(angle float64, frequency float64) float64 {
	return 2*Sine(angle, frequency/4) - 0.8*SquareWithDecay(angle, frequency*frequency)
}

func Additive(angle float64, frequency float64) float64 {
	if angle < 0.5*math.Pi {
		return Sine(angle, frequency) +
			0.5*Sine(angle, frequency/100) +
			0.3*SquareWithDecay(angle, frequency/4)
	}
	return 3 * Square(angle, frequency) * math.Exp(-angle/5) * rand.Float64()
}

func Square(angle float64, frequency float64) float64 {
	if Sine(angle, frequency) >= 0 {
		return 1.0
	}
	return -1.0
}

func SquareWithDecay(angle float64, frequency float64) float64 {
	return Square(angle, frequency) * math.Exp(-angle/3)
}

func Thud(angle float64, _ float64) float64 { //untested
	frequency := 100 * math.Exp(-angle)
	multipliers := []float64{5, 15}
	return multipliers[rand.Intn(len(multipliers))] * Sine(angle, frequency)
}

func Snare(angle float64, frequency float64) float64 { //untested
	sweepFrequency := 10 * frequency * math.Exp(-angle*3)
	return 10 * Sine(angle, sweepFrequency) * float64(rand.Intn(10)) * math.Exp(-angle/10)
}

func Lipsmack(angle float64, frequency float64) float64 { //untested
	sweepFrequency := 50 * frequency * math.Exp(-angle*5)
	return 10 * Sine(angle, sweepFrequency) * math.Exp(-angle*10)
}
