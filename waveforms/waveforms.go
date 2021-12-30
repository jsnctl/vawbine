package waveforms

import (
	"math"
	"math/rand"
)

func GetRandomWaveFn() func(angle float64, frequency float64) float64 {
	waveFns := []func(angle float64, frequency float64) float64 {
		Lipsmack,
		Thud,
		Snare,
		Additive,
	}
	return waveFns[rand.Intn(len(waveFns))]
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func SineWithDecay(angle float64, frequency float64) float64 {
	if angle < math.Pi {
		return 3 * Sine(angle, 0.5*frequency) - (0.3 * Additive(angle, 3*frequency))
	}
	return 3*Sine(angle, frequency) * math.Exp(-angle/5)
}

func Additive(angle float64, frequency float64) float64 {
	if angle < 0.5 * math.Pi {
		return Sine(angle, frequency) +
			0.5 * Sine(angle, frequency/100) +
			0.3 * SquareWithDecay(angle, frequency/4)
	}
	return 3*Square(angle, frequency) * math.Exp(-angle/5) * rand.Float64()
}

func Square(angle float64, frequency float64) float64 {
	if Sine(angle, frequency) >= 0 {
		return 1.0
	}
	return -1.0
}

func SquareWithDecay(angle float64, frequency float64) float64 {
	return Square(angle, frequency) * math.Exp(-angle)
}

func Thud(angle float64, _ float64) float64 { //untested
	frequency := 100*math.Exp(-angle)
	multipliers := []float64{5, 15}
	return multipliers[rand.Intn(len(multipliers))]*Sine(angle, frequency)
}

func Snare(angle float64, frequency float64) float64 { //untested
	sweepFrequency := 10*frequency*math.Exp(-angle*3)
	return 10 * Sine(angle, 0.5*sweepFrequency) * float64(rand.Intn(10)) * math.Exp(-angle*4)
}

func Lipsmack(angle float64, frequency float64) float64 { //untested
	sweepFrequency := 50*frequency*math.Exp(-angle*5)
	return 10 * Sine(angle, sweepFrequency) * math.Exp(-angle*10)
}
