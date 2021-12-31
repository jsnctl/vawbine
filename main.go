package main

func main() {
	generator := newGenerator(stutter(20))
	generator.generate()
}
