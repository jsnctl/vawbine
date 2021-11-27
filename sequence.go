package main

import (
	"math/rand"
)

type Sequence struct {
	numberOfSteps int
	Stack []int
	Cursor int
}

func createSequence(numberOfSteps int) Sequence {
	sequence := Sequence{
		numberOfSteps: numberOfSteps,
	}
	sequence.Stack = make([]int, numberOfSteps)
	for i, _ := range sequence.Stack  {
		sequence.Stack[i] = i + rand.Intn(15)
	}
	sequence.Cursor = 0
	return sequence
}


