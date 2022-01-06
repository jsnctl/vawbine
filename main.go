package main

func main() {
	generator := newGenerator(stutter(60, 40, 0))
	generator.generate()
}
