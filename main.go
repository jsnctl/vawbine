package main

func main() {
	generator := newGenerator(ramp(120, 50, 5))
	generator.generate()
}
