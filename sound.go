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

func generate() {

	durations := make([]float64, 100)
	minDuration := 0.001
	maxDuration := 0.5
	for i := range durations {
		durations[i] = minDuration + rand.Float64() * (maxDuration - minDuration)
	}

	frequencies := make([]float64, 100)
	min := 1000.0
	max := 3000.0
	for i := range frequencies {
		frequencies[i] = min + rand.Float64() * (max - min)
	}

	output := "sound.bin"
	f, _ := os.Create(output)

	for j, frequency := range frequencies {
		nSamples := int(durations[j] * SampleRate)
		tau := math.Pi * 2

		var angle = tau / float64(nSamples)

		for i := 0; i <= nSamples; i++ {
			sample := 5.0 * math.Sin(angle * frequency * float64(i))
			fmt.Printf("%v \n", sample)
			var buf [8]byte
			binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
			f.Write(buf[:])
		}
	}


}
