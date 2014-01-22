package mastermind

import (
	"testing"
)

func TestEvaluate_0(t *testing.T) {
	c := Combination{[4]color{Red, Red, Red, Red}}
	m := new(Combination)

	e := c.Evaluate(m)
	if e.Correct() != 0 {
		t.Error("Correct should be 0, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_1(t *testing.T) {
	c := Combination{[4]color{Red, Red, Red, Red}}
	m := Combination{[4]color{Red, Red, Red, Red}}

	e := c.Evaluate(&m)
	if e.Correct() != 4 {
		t.Error("Correct should be 4, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_2(t *testing.T) {
	c := Combination{[4]color{Red, Red, Red, Yellow}}
	m := Combination{[4]color{Red, Yellow, Yellow, Yellow}}

	e := c.Evaluate(&m)
	if e.Correct() != 2 {
		t.Error("Correct should be 2, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_3(t *testing.T) {
	c := Combination{[4]color{Black, White, Red, Yellow}}
	m := Combination{[4]color{Red, White, Yellow, Blue}}

	e := c.Evaluate(&m)
	if e.Correct() != 1 {
		t.Error("Correct should be 1, is", e.Correct())
	}
	if e.Present() != 2 {
		t.Error("Present should be 2, is", e.Present())
	}
}

func TestEvaluate_4(t *testing.T) {
	c := Combination{[4]color{Red, Red, Red, Yellow}}
	m := Combination{[4]color{Red, Yellow, Red, Yellow}}

	e := c.Evaluate(&m)
	if e.Correct() != 3 {
		t.Error("Correct should be 3, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}
