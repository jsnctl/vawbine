package waveforms

import "math"

type Waveform string

const (
	Sine Waveform = "sine"
	Square = "square"
	Sawtooth = "sawtooth"
)

func (waveform Waveform) isValid() bool {
	switch waveform {
	case Sine, Square, Sawtooth:
		return true
	}
	return false
}

func sine(angle float64, frequency float64, index int) float64 {
	return math.Sin(angle * frequency * float64(index))
}

func square(angle float64) float64 {
	if angle < math.Pi {
		return 1.0
	}
	return -1.0
}
