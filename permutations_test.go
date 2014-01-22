package mastermind

import (
	"testing"
)

func TestPermutations(t *testing.T) {
	perm := Permutations()
	allblack := Combination{Black, Black, Black, Black}
	allyellow := Combination{Yellow, Yellow, Yellow, Yellow}
	if len(perm) != 1296 {
		t.Error("Number of permutations shoud be 1296, is", len(perm))
	}
	if perm[0] != allblack {
		t.Error("First permutation should be all black, is", perm[0])
	}
	if perm[len(perm)-1] != allyellow {
		t.Error("Last permutation should be all yellow, is",
			perm[len(perm)-1])
	}
}
