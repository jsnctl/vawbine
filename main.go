package main

func main() {
	generator := newGenerator(stutter(100, 30, 0))
	generator.generate()
}
