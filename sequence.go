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

	muValues := []float64{5, 10, 15}
	longMuValues := []float64{1.05, 1.55, 1.95}

	sequence.Stack = make([]float64, numberOfSteps)
	for i, _ := range sequence.Stack {
		if i % 8 == 0 {
			longMu := longMuValues[rand.Intn(len(longMuValues))]
			sequence.Stack[i] = longMu * normal(muValues[rand.Intn(len(muValues))], 0.02)
		} else {
			sequence.Stack[i] = normal(muValues[rand.Intn(len(muValues))], 0.02)
		}
	}
	return sequence
}
