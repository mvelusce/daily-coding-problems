package main

import (
	"fmt"

	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fs := []Flight{
		{"SFO", "HKO"}, {"YYZ", "SFO"}, {"YUL", "YYZ"}, {"HKO", "ORD"},
	}
	res := ComputeItinerary("YUL", fs)
	fmt.Println(res)

	assert.Equal(t, []string{"YUL", "YYZ", "SFO", "HKO", "ORD"}, res)
}

func TestMainMoreItineraries(t *testing.T) {
	fs := []Flight{
		{"A", "B"}, {"A", "C"}, {"B", "C"}, {"C", "A"},
	}
	res := ComputeItinerary("A", fs)
	fmt.Println(res)

	assert.Equal(t, []string{"A", "B", "C", "A", "C"}, res)
}
