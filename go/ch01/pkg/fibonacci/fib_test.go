package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroAndOne(t *testing.T) {
	assert.Equal(t, Fib(0), uint(0))
	assert.Equal(t, Fib2(0), uint(0))
	assert.Equal(t, Fib3(0), uint(0))
	assert.Equal(t, Fib(1), uint(1))
	assert.Equal(t, Fib2(1), uint(1))
	assert.Equal(t, Fib3(1), uint(1))
}

var fibonacci = []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}

func TestTwoThroughFifteen(t *testing.T) {
	for n := uint(2); n < 16; n++ {
		assert.Equal(t, Fib(n), fibonacci[n])
		assert.Equal(t, Fib2(n), fibonacci[n])
		assert.Equal(t, Fib3(n), fibonacci[n])
	}
}
