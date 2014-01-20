package mastermind

import (
	"testing"
)

func TestIncCorrect(t *testing.T) {
	e := new(Evaluation)
	if e.correct != 0 {
		t.Error("Evaluation.Correct not zero")
	}
	e.IncCorrect()
	if e.correct != 1 {
		t.Error("Evaluation.Correct not incremented")
	}
}

func TestIncPresent(t *testing.T) {
	e := new(Evaluation)
	if e.present != 0 {
		t.Error("Evaluation.Present not zero")
	}
	e.IncPresent()
	if e.present != 1 {
		t.Error("Evaluation.Present not incremented")
	}
}
