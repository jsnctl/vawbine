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

type Buffer struct {
	values [][]float64
}

func (generator *Generator) generate() {
	f, _ = os.Create(shared.OutputFile)
	durations := []float64{0.4, 0.3, 0.1}

	buffer := Buffer{}

	for _, seed := range generator.Sequence.Stack {
		duration := durations[rand.Intn(len(durations))]
		output := note(seed, duration, waveforms.SineWithDecay, 0, 1)
		buffer.values = append(buffer.values, output)
	}

	for _, value := range buffer.values {
		monophonic(value)
	}
}

func combineNotes(left []float64, right []float64, decay float64) []float64 { // untested
	var result []float64
	choices := [][]float64{left, right}
	for i, _ := range longer(left, right) {
		choice := choices[rand.Intn(2)]
		result = append(result, decay*choice[i])
	}
	return result
}

func monophonic(note []float64) {
	for _, value := range note {
		var buf = make([]byte, 8)
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(value)))
		f.Write(buf[:])
	}
}

func polyphony(left [][]byte, right [][]byte) {
	for i := 0; i < min(len(left), len(right)); i++ {
		writeByte(left[i])
		writeByte(right[i])
	}
}

func longer(left []float64, right []float64) []float64 {
	if len(left) <= len(right) {
		return right
	}
	return left
}

func min(left int, right int) int {
	if left > right {
		return right
	}
	return left
}

func note(seed float64, duration float64, waveFn func(float64, float64) float64, shift float64, amplitude float64) []float64 {
	nSamples := int(duration * shared.SampleRate)
	tau := math.Pi * 2

	var angleIncr = tau / float64(nSamples)

	var shiftIncr float64
	if shift != 0 {
		shiftIncr = shift / float64(nSamples)
	}

	var note []float64
	for i := 0; i <= nSamples; i++ {
		angle := (angleIncr + shiftIncr) * float64(i)
		sample := amplitude * waveFn(angle, seed)
		note = append(note, sample)
	}

	var ghost []float64
	ghost = note
	delays := []int{200}
	delay := delays[rand.Intn(len(delays))]
	for i, _ := range note {
		if i == 0 || i >= len(note) - 1 || i < delay  {
			continue
		}
		note[i] = note[i] * 0.55*ghost[i-(delay-1)]
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
