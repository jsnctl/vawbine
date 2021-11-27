package main

func main() {
	generator := newGenerator(createSequence(1000), "sine")
	generator.generate()
}
