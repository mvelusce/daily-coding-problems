package main

import (
	"fmt"
)

// Flight is tuple with origin and destination
type Flight struct {
	origin      string
	destination string
}

// ComputeItinerary computes itinerary
func ComputeItinerary(start string, flights []Flight) []string {
	fmt.Println(start)

	fs := make(map[string]string)

	for _, f := range flights {
		fs[f.origin] = f.destination
	}

	itinerary := []string{start}

	for i := 0; i < len(itinerary); i++ {
		x := itinerary[i]
		if s, ok := fs[x]; ok {
			itinerary = append(itinerary, s)
			delete(fs, x)
		}
	}

	return itinerary
}

func main() {
	fs := []Flight{
		{"SFO", "HKO"}, {"YYZ", "SFO"}, {"YUL", "YYZ"}, {"HKO", "ORD"},
	}
	res := ComputeItinerary("YUL", fs)
	fmt.Println(res)
}
