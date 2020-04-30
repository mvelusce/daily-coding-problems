package main

import (
	"math"
	"strings"
)

const PageWidth int = 80

type CachedResults struct {
	badness float64
	text    string
}

var badnessMem = map[int]CachedResults{}

func badness(lineLength int, pageWidth int) float64 {

	if lineLength > pageWidth {
		return math.MaxFloat64
	} else {
		return math.Pow(float64(pageWidth-lineLength), 3)
	}
}

func textJustification(text string, i int) (float64, string) {

	if len(text) <= PageWidth {
		return badness(len(text), PageWidth), text
	}

	words := strings.Split(text, " ")
	minBadness := math.MaxFloat64
	finalText := ""
	for j := 1; j < len(words); j++ {
		badnessIJ := badness(len(strings.Join(words[:j], " ")), PageWidth)
		badnessJ := math.MaxFloat64
		partText := ""
		if cachedRes, ok := badnessMem[i+j]; ok {
			badnessJ = cachedRes.badness
			partText = cachedRes.text
		} else {
			badnessJ, partText = textJustification(strings.Join(words[j:], " "), i+j)
			badnessMem[i+j] = CachedResults{
				badness: badnessJ,
				text:    partText,
			}
		}

		if badnessIJ+badnessJ < minBadness {
			finalText = strings.Join(words[:j], " ") + "\n" + partText
			minBadness = badnessIJ + badnessJ
		}
	}
	return minBadness, finalText
}

func main() {

	badness, res := textJustification("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", 0)

	println(badness)
	println(res)
}
