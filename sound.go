package main

import (
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

type Note struct {
	wave []float64
}

type Buffer struct {
	values []Note
}

func (generator *Generator) generate() {
	f, _ = os.Create(shared.OutputFile)
	durations := []float64{0.1, 0.11, 0.09, 0.05}

	buffer := Buffer{}

	for _, seed := range generator.Sequence.Stack {
		duration := durations[rand.Intn(len(durations))]
		output := createNote(seed, duration, waveforms.TriangleWithDecay, 0, 1)
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

func createNote(seed float64, duration float64, waveFn func(float64, float64) float64, shift float64, amplitude float64) Note {
	note := Note{}
	nSamples := int(duration * shared.SampleRate)
	tau := math.Pi * 2

	var angleIncr = tau / float64(nSamples)

	var shiftIncr float64
	if shift != 0 {
		shiftIncr = shift / float64(nSamples)
	}

	for i := 0; i <= nSamples; i++ {
		angle := (angleIncr + shiftIncr) * float64(i)
		sample := amplitude * waveFn(angle, seed)
		note.wave = append(note.wave, sample)
	}
	return note
}

func reverb(note Note, delay int, decay float64) Note {
	var ghost []float64
	ghost = note.wave
	for i, _ := range note.wave {
		if i == 0 || i >= len(note.wave) - 1 || i < delay  {
			continue
		}
		note.wave[i] = note.wave[i] * (decay*ghost[i-(delay-1)])
	}
	return note
}

func distortion(note Note) Note {
	// udo zolzer
	// https://dsp.stackexchange.com/questions/13142/digital-distortion-effect-algorithm
	for i, _ := range note.wave {
		preValue := note.wave[i]
		if preValue > 0 {
			note.wave[i] = 1 - math.Exp(-preValue)
		} else {
			note.wave[i] = -1 + math.Exp(preValue)
		}
	}
	return note

}
