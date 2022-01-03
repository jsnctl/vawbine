package main

import (
	"encoding/binary"
	"math"
)

func monophonic(note Note) {
	for _, value := range note.wave {
		writeNoteValue(value)
	}
}

func polyphony(left Note, right Note) {
	for i := 0; i < min(len(left.wave), len(right.wave)); i++ {
		writeNoteValue(left.wave[i])
		writeNoteValue(right.wave[i])
	}
}

func writeNoteValue(value float64) {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(value)))
	f.Write(buf[:])
}
