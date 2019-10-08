package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceUintInIn(t *testing.T) {
	z := []uint{1, 2, 3}
	assert.Equal(t, SliceUintIn(z, 1), true)
}
