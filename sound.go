package main

import (
	"encoding/binary"
	"github.com/jsnctl/vawbine/shared"
	"github.com/jsnctl/vawbine/waveforms"
	"math"
	"math/rand"
	"os"
)

type Generator struct {
	Sequence Sequence
}

func newGenerator(sequence Sequence) Generator {
	generator := Generator{
		Sequence: sequence,
	}
	return generator
}

var f *os.File

func (generator *Generator) generate() {
	f, _ = os.Create(shared.OutputFile)
	durations := []float64{0.025}
	for _, seed := range generator.Sequence.Stack {
		duration := durations[rand.Intn(len(durations))]
		waveFn := waveforms.GetRandomWaveFn()
		note(seed, duration, waveFn)
	}
}

func note(seed float64, duration float64, waveFn func(float64, float64) float64) {
	nSamples := int(duration * shared.SampleRate)
	tau := math.Pi * 2

	var angleIncr = tau / float64(nSamples)

	for i := 0; i <= nSamples; i++ {
		angle := angleIncr * float64(i)
		sample := waveFn(angle, seed)
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		write(buf)
	}
}

func write(buf [8]byte) {
	f.Write(buf[:])
}
