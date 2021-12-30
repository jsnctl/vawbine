package waveforms

import (
	"github.com/jsnctl/vawbine/shared"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var nSamples int
var frequency float64
var angleIncrement float64

const testSampleRate = 44000

func Init() {
	nSamples = testSampleRate * 1
	frequency = 1000.0
	angleIncrement = (2 * math.Pi) / float64(nSamples)
}

func TestSine(t *testing.T) {
	Init()
	result := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		angle := float64(i) * angleIncrement
		result[i] = Sine(angle, frequency)
	}
	expectedFirstTen := []float64{
		0, 0.14231483827328514, 0.28173255684142967, 0.4154150130018864,
		0.5406408174555976, 0.6548607339452851, 0.7557495743542582,
		0.8412535328311811, 0.9096319953545183, 0.9594929736144975,
	}

	for i, expected := range expectedFirstTen {
		assert.True(t,
			shared.FloatingPointEqual(
				expected,
				result[i],
				1e-10,
			),
		)
	}
}

func TestSquare(t *testing.T) {
	Init()
	result := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		angle := float64(i) * angleIncrement
		result[i] = Square(angle, frequency)
	}

	assert.Equal(t, 1.0, result[0])
	assert.Equal(t, -1.0, result[23])
	assert.Equal(t, 1.0, result[45])
	assert.Equal(t, -1.0, result[67])
}

func TestThud(t *testing.T) {
	Init()
	result := make([]float64, nSamples)
	for i := 0; i < nSamples; i++ {
		angle := float64(i) * angleIncrement
		result[i] = Thud(angle, frequency)
	}

	assert.Equal(t, 0.0, result[0])
}