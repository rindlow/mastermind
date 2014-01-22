package mastermind

import (
	"testing"
)

func TestPermutations(t *testing.T) {
	perm := Permutations()
	allblack := Combination{[4]color{Black, Black, Black, Black}}
	allyellow := Combination{[4]color{Yellow, Yellow, Yellow, Yellow}}
	if len(perm) != 6 * 6 * 6 * 6 {
		t.Error("Number of permutations shoud be 1296, is", len(perm))
	}
	if perm[0] != allblack {
		t.Error("First permutation should be all black, is", perm[0])
	}
	if perm[len(perm)-1] != allyellow {
		t.Error("First permutation should be all yellow, is", 
			perm[len(perm)-1])
	}
}
