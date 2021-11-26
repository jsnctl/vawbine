package main

import (
	"fmt"
	"time"
)

type Sequence struct {
	numberOfSteps int
	stepTime float64
	Stack []int
	Cursor int
}

func createSequence(numberOfSteps int, stepTime float64) Sequence {
	sequence := Sequence{
		numberOfSteps: numberOfSteps,
		stepTime: stepTime,
	}
	sequence.Stack = make([]int, numberOfSteps)
	for i, _ := range sequence.Stack  {
		sequence.Stack[i] = i + 10
	}
	sequence.Cursor = 0
	return sequence
}

func action(timer *Sequence) {
	fmt.Println(timer.Stack[timer.Cursor])
}

func (sequence *Sequence) proceed() {
	action(sequence)
	time.Sleep(time.Duration(sequence.stepTime) * time.Second)
	if sequence.Cursor < (sequence.numberOfSteps - 1) {
		sequence.Cursor += 1
	} else {
		sequence.Cursor = 0
	}
}


