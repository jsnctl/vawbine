package main

func main() {
	generator := newGenerator(stutter(100))
	generator.generate()
}
