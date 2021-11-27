package waveforms

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var nSamples int
var frequency float64
var angle float64
const testSampleRate = 10

func Init() {
	nSamples = testSampleRate * 1
	frequency = 1.0
	angle = 1.0
}

func TestSine(t *testing.T) {
	Init()
	result := make([]float64, nSamples)
	expectedResult := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		result[i] = sine(angle, frequency, i)
		expectedResult[i] = math.Sin(frequency * angle * float64(i))
	}

	assert.Equal(t, expectedResult, result)
}

func TestSquare(t *testing.T) {
	Init()
	result := make([]float64, nSamples)
	expectedResult := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		result[i] = square(angle)
		expectedResult[i] = math.Sin(frequency * angle * float64(i))
	}

	assert.Equal(t, expectedResult, result)
}
