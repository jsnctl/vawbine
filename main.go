package main

func main() {
	pattern := ramp(5, 21, 7)
	generator := newGenerator(arp(100, pattern.Stack))
	generator.generate()
}
