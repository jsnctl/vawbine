package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
)

const (
	SampleRate = 44100
)

type Generator struct {
	Sequence Sequence
}

func (generator *Generator) generate() {

	durations := make([]float64, 100)
	minDuration := 0.001
	maxDuration := 0.5
	for i := range durations {
		durations[i] = minDuration + rand.Float64() * (maxDuration - minDuration)
	}

	output := "sound.bin"
	f, _ := os.Create(output)

	for j, seed := range generator.Sequence.Stack {
		nSamples := int(durations[j] * SampleRate)
		tau := math.Pi * 2

		var angle = tau / float64(nSamples)

		frequency := float64(1000.0 * seed)

		for i := 0; i <= nSamples; i++ {
			sample := 5.0 * math.Sin(angle * frequency * float64(i))
			fmt.Printf("%v \n", sample)
			var buf [8]byte
			binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
			f.Write(buf[:])
		}
	}


}
