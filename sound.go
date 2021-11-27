package main

import (
	"encoding/binary"
	"github.com/jsnctl/gotechre/waveforms"
	"math"
	"math/rand"
	"os"
)

const (
	SampleRate = 44100
)

type Generator struct {
	Sequence Sequence
	Waveform waveforms.Waveform
}

func newGenerator(sequence Sequence, waveform waveforms.Waveform) Generator {
	generator := Generator{
		Sequence: sequence,
		Waveform: waveform,
	}
	return generator
}

func (generator *Generator) generate() {
	durations := make([]float64, 100)
	minDuration := 1.0E-2
	maxDuration := 5.0E-2
	for i := range durations {
		durations[i] = minDuration + rand.Float64() * (maxDuration - minDuration)
	}

	output := "sound.bin"
	f, _ := os.Create(output)

	for j, seed := range generator.Sequence.Stack {
		nSamples := int(durations[j] * SampleRate)
		tau := math.Pi * 2

		var angle = tau / float64(nSamples)

		frequency := 5.0 * seed

		for i := 0; i <= nSamples; i++ {
			sample := 5.0 * math.Sin(angle * frequency * float64(i))
			var buf [8]byte
			binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
			f.Write(buf[:])
		}
	}


}
