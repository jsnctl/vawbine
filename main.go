package main

func main() {
	generator := newGenerator(createSequence(50))
	generator.generate()
}
