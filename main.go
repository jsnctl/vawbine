package main

func main() {
	generator := newGenerator(createSequence(300))
	generator.generate()
}
