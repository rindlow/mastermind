package mastermind

import (
//	"log"
)

type color int

const (
	Black = iota
	White
	Red
	Green
	Blue
	Yellow
)

type Combination [4]color

func (c *Combination) SetColorAt(col color, i int) {
	c[i] = col
}

func (this *Combination) Evaluate(other *Combination) *Evaluation {
	var found, taken [4]bool
	e := new(Evaluation)
	for i := 0; i < 4; i++ {
		if this[i] == other[i] {
			//log.Printf("Correct, i=%d", i)
			e.IncCorrect()
			found[i] = true
			taken[i] = true
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if this[i] == other[j] &&
				!found[i] && !taken[j] {
				//log.Printf("Present, i=%d, j=%d", i, j)
				e.IncPresent()
				found[i] = true
				taken[j] = true
				break
			}
		}
	}
	return e
}
