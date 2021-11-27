package main

func main() {
	generator := Generator{Sequence: createSequence(5000)}
	generator.generate()
}
