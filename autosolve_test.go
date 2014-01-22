package mastermind

import (
	"testing"
	"math/rand"
)

func TestAutoSolve(t *testing.T) {
	rand.Seed(4711)
	if AutoSolve(true) != 5 {
		t.Error("Hidden not guessed in 5 tries")
	}
}
