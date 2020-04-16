package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOddEvenJumps(t *testing.T) {

	in := []int{10, 13, 12, 14, 15}

	res := oddEvenJumps(in)

	assert.Equal(t, 2, res)
}

func TestOddEvenJumps1(t *testing.T) {

	in := []int{2, 3, 1, 1, 4}

	res := oddEvenJumps(in)

	assert.Equal(t, 3, res)
}

func TestOddEvenJumps2(t *testing.T) {

	in := []int{5, 1, 3, 4, 2}

	res := oddEvenJumps(in)

	assert.Equal(t, 3, res)
}
