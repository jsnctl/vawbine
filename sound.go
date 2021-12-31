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
	durations := []float64{0.1, 0.2}
	for _, seed := range generator.Sequence.Stack {
		duration := durations[rand.Intn(len(durations))]
		output := note(seed, duration, waveforms.SquareWithDecay, 0, 1)
		polyphony(output, note(seed, duration-0.05, waveforms.SquareWithDecay, math.Pi/4, 0.8))
		polyphony(output, note(seed+20, duration+0.3, waveforms.SquareWithDecay, math.Pi/8, 0.6))
		polyphony(output, note(seed, 0.05, waveforms.Thud, 0, 1))
	}
}

func polyphony(left [][]byte, right [][]byte) {
	for i := 0; i < min(len(left), len(right)); i++ {
		writeByte(left[i])
		writeByte(right[i])
	}
}

func min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func note(seed float64, duration float64, waveFn func(float64, float64) float64, shift float64, amplitude float64) [][]byte {
	nSamples := int(duration * shared.SampleRate)
	tau := math.Pi * 2

	var angleIncr = tau / float64(nSamples)

	var shiftIncr float64
	if shift != 0 {
		shiftIncr = shift / float64(nSamples)
	}

	var note [][]byte
	for i := 0; i <= nSamples; i++ {
		angle := (angleIncr + shiftIncr) * float64(i)
		sample := amplitude * waveFn(angle, seed)
		var buf = make([]byte, 8)
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		note = append(note, buf)
	}
	return note
}

func write(note [][]byte) {
	for _, buf := range note {
		f.Write(buf[:])
	}
}

func writeByte(b []byte) {
	f.Write(b[:])
}
