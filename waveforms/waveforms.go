package waveforms

import (
	"math"
	"math/rand"
)

func GetRandomWaveFn() func(angle float64, frequency float64) float64 {
	waveFns := []func(angle float64, frequency float64) float64{
		Thud,
		Snare,
	}
	return waveFns[rand.Intn(len(waveFns))]
}

func Sine(angle float64, frequency float64) float64 {
	return math.Sin(angle * frequency)
}

func SineWithDecay(angle float64, frequency float64) float64 {
	return Sine(angle, frequency) * math.Exp(-angle/5)
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

func Triangle(angle float64, frequency float64) float64 {
	sineValue := Sine(angle, frequency)
	return (4/(2*math.Pi))*math.Asin(sineValue)
}

func TriangleWithDecay(angle float64, frequency float64) float64 {
	return Triangle(angle, frequency) * math.Exp(-angle/10)
}

func Sawtooth(angle float64, frequency float64) float64 {
	cotValue := math.Tan(angle * frequency)
	return (2/math.Pi)*math.Atan(1.0/cotValue)
}

func Thud(angle float64, _ float64) float64 { //untested
	frequency := 100 * math.Exp(-angle)
	multipliers := []float64{5, 15}
	return multipliers[rand.Intn(len(multipliers))] * Sine(angle, frequency)
}

func Snare(angle float64, frequency float64) float64 { //untested
	sweepFrequency := 10 * frequency * math.Exp(-angle*3)
	return 0.5*Sine(angle, sweepFrequency) * float64(rand.Intn(10)) * math.Exp(-angle/10)
}

func Lipsmack(angle float64, frequency float64) float64 { //untested
	sweepFrequency := 50 * frequency * math.Exp(-angle*5)
	return 10 * Sine(angle, sweepFrequency) * math.Exp(-angle*10)
}
