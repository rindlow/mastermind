package mastermind

import (
	"testing"
)

func TestEliminate_1(t *testing.T) {
	c := make([]Combination, 2)
	c[0].SetColorAt(Red, 0)
	c[0].SetColorAt(Green, 1)
	c[0].SetColorAt(Blue, 2)
	c[0].SetColorAt(Black, 3)
	c[1].SetColorAt(Red, 0)
	c[1].SetColorAt(Yellow, 1)
	c[1].SetColorAt(Black, 2)
	c[1].SetColorAt(White, 3)

	x := make([]Try, 1)
	x[0].comb.SetColorAt(Red, 0)
	x[0].comb.SetColorAt(Blue, 1)
	x[0].comb.SetColorAt(Black, 2)
	x[0].comb.SetColorAt(Green, 3)
	x[0].eval.correct = 1
	x[0].eval.present = 3

	res := Eliminate(c, x)
	if len(res) != 1 {
		t.Error("res should be 1 element long, is", len(res))
	}
	if res[0] != c[0] {
		t.Error("res[0] should be equal to c[0]")
	}
}
