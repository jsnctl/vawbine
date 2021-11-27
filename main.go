package main

import (
	"math/rand"
)

func play() {
	stepTime := rand.Intn(5)
	sequence := createSequence(8, float64(stepTime))
	for {
		sequence.proceed()
	}
}

func main() {
	//go play()
	//go play()
	//time.Sleep(100 * time.Second)

	generate()

}
