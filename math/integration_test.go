package math

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const delta = 0.001

func TestDefiniteIntegration(t *testing.T) {
	I := DefiniteIntegration(-1, 1, 20, func(float64) float64 {
		return 1
	})
	assert.InDelta(t, I, 2.0, delta)

	I = DefiniteIntegration(-1, 1, 20, func(t float64) float64 {
		return math.Sin(t)
	})
	assert.InDelta(t, I, 0.0, delta)
}
