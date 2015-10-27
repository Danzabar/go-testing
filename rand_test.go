package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandStr(t *testing.T) {
	rand := randStr(5)

	assert.Equal(t, len(rand), 5)
}

func BenchmarkRand(b *testing.B) {
	randStr(10)
}
