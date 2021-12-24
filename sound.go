package main

import (
	"encoding/binary"
	"github.com/jsnctl/gotechre/shared"
	"github.com/jsnctl/gotechre/waveforms"
	"math"
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
	for _, seed := range generator.Sequence.Stack {
		note(5.0*seed, 0.04)
	}
}

func note(seed float64, duration float64) {
	nSamples := int(duration * shared.SampleRate)
	tau := math.Pi * 2

	var angleIncr = tau / float64(nSamples)

	for i := 0; i <= nSamples; i++ {
		angle := angleIncr * float64(i)
		sample := 5.0 * waveforms.Square(angle, seed)
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		write(buf)
	}
}

func write(buf [8]byte) {
	f.Write(buf[:])
}
