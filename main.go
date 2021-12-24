package main

func main() {
	generator := newGenerator(createSequence(100))
	generator.generate()
}
