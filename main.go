package main

func main() {
	generator := Generator{Sequence: createSequence(100)}
	generator.generate()
}
