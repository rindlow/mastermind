package mastermind

import (
	"testing"
)

func TestEvaluate_0(t *testing.T) {
	c := Combination{Red, Red, Red, Red}
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
	c := Combination{Red, Red, Red, Red}
	m := &Combination{Red, Red, Red, Red}

	e := c.Evaluate(m)
	if e.Correct() != 4 {
		t.Error("Correct should be 4, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_2(t *testing.T) {
	c := Combination{Red, Red, Red, Yellow}
	m := &Combination{Red, Yellow, Yellow, Yellow}

	e := c.Evaluate(m)
	if e.Correct() != 2 {
		t.Error("Correct should be 2, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestEvaluate_3(t *testing.T) {
	c := Combination{Black, White, Red, Yellow}
	m := &Combination{Red, White, Yellow, Blue}

	e := c.Evaluate(m)
	if e.Correct() != 1 {
		t.Error("Correct should be 1, is", e.Correct())
	}
	if e.Present() != 2 {
		t.Error("Present should be 2, is", e.Present())
	}
}

func TestEvaluate_4(t *testing.T) {
	c := Combination{Red, Red, Red, Yellow}
	m := &Combination{Red, Yellow, Red, Yellow}

	e := c.Evaluate(m)
	if e.Correct() != 3 {
		t.Error("Correct should be 3, is", e.Correct())
	}
	if e.Present() != 0 {
		t.Error("Present should be 0, is", e.Present())
	}
}

func TestColorString(t *testing.T) {
	var c color
	if color(Black).String() != "Black" {
		t.Error("Black's not Black but", c.String())
	}
	if color(White).String() != "White" {
		t.Error("White's not White but", c.String())
	}
	if color(Red).String() != "Red" {
		t.Error("Red's not Red but", c.String())
	}
	if color(Green).String() != "Green" {
		t.Error("Green's not Green but", c.String())
	}
	if color(Blue).String() != "Blue" {
		t.Error("Blue's not Blue but", c.String())
	}
	if color(Yellow).String() != "Yellow" {
		t.Error("Yellow's not Yellow but", c.String())
	}
}
