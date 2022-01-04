package main

func main() {
	generator := newGenerator(ramp(400, 40, 0.1))
	generator.generate()
}
