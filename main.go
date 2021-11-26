package main

func play() {
	sequence := createSequence(8, 1)
	for {
		sequence.proceed()
	}
}

func main() {
	play()
}
