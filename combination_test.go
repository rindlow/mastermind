package mastermind

import (
	"testing"
)

func TestEvaluate_0(t *testing.T) {
	c := new(Combination)
	c.SetColorAt(Red, 0)
	c.SetColorAt(Red, 1)
	c.SetColorAt(Red, 2)
	c.SetColorAt(Red, 3)

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
	c := new(Combination)
	c.SetColorAt(Red, 0)
	c.SetColorAt(Red, 1)
	c.SetColorAt(Red, 2)
	c.SetColorAt(Red, 3)

	m := new(Combination)
	m.SetColorAt(Red, 0)
	m.SetColorAt(Red, 1)
	m.SetColorAt(Red, 2)
	m.SetColorAt(Red, 3)

	e := c.Evaluate(m)
	if e.Correct() != 4 {
		t.Error("Correct should be 4, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_2(t *testing.T) {
	c := new(Combination)
	c.SetColorAt(Red, 0)
	c.SetColorAt(Red, 1)
	c.SetColorAt(Red, 2)
	c.SetColorAt(Yellow, 3)

	m := new(Combination)
	m.SetColorAt(Red, 0)
	m.SetColorAt(Yellow, 1)
	m.SetColorAt(Yellow, 2)
	m.SetColorAt(Yellow, 3)

	e := c.Evaluate(m)
	if e.Correct() != 2 {
		t.Error("Correct should be 2, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_3(t *testing.T) {
	c := new(Combination)
	c.SetColorAt(Black, 0)
	c.SetColorAt(White, 1)
	c.SetColorAt(Red, 2)
	c.SetColorAt(Yellow, 3)

	m := new(Combination)
	m.SetColorAt(Red, 0)
	m.SetColorAt(White, 1)
	m.SetColorAt(Yellow, 2)
	m.SetColorAt(Blue, 3)

	e := c.Evaluate(m)
	if e.Correct() != 1 {
		t.Error("Correct should be 1, is", e.Correct())
	}
	if e.Present() != 2 {
		t.Error("Present should be 2, is", e.Present())
	}
}

func TestEvaluate_4(t *testing.T) {
	c := new(Combination)
	c.SetColorAt(Red, 0)
	c.SetColorAt(Red, 1)
	c.SetColorAt(Red, 2)
	c.SetColorAt(Yellow, 3)

	m := new(Combination)
	m.SetColorAt(Red, 0)
	m.SetColorAt(Yellow, 1)
	m.SetColorAt(Red, 2)
	m.SetColorAt(Yellow, 3)

	e := c.Evaluate(m)
	if e.Correct() != 3 {
		t.Error("Correct should be 3, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}
