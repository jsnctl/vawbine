package main

import (
	"math/rand"
)

type Sequence struct {
	numberOfSteps int
	Stack         []float64
}

func normal(mu float64, sigma float64) float64 {
	return rand.NormFloat64()*sigma + mu
}

func createSequence(numberOfSteps int) Sequence {
	sequence := Sequence{
		numberOfSteps: numberOfSteps,
	}

	muValues := []float64{50, 100, 150}
	longMuValues := []float64{1.05, 1.55, 1.95}

	sequence.Stack = make([]float64, numberOfSteps)
	for i, _ := range sequence.Stack {
		if i%8 == 0 {
			longMu := longMuValues[rand.Intn(len(longMuValues))]
			sequence.Stack[i] = longMu * normal(muValues[rand.Intn(len(muValues))], 0.02)
		} else {
			sequence.Stack[i] = normal(muValues[rand.Intn(len(muValues))], 0.02)
		}
	}
	return sequence
}

func stutter(length int) Sequence {
	sequence := Sequence{
		numberOfSteps: length,
	}

	root := 50.0
	sequence.Stack = make([]float64, length)

	for i, _ := range sequence.Stack {
		sequence.Stack[i] = root
	}

	return sequence
}

func ramp(length int, start float64, step float64) Sequence {
	sequence := Sequence{
		numberOfSteps: length,
	}
	sequence.Stack = make([]float64, length)
	for i, _ := range sequence.Stack {
		sequence.Stack[i] = start
		start = start + step
	}
	return sequence
}
