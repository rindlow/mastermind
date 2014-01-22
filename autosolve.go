package mastermind

import (
	"fmt"
	"math/rand"
)

func AutoSolve(verbose bool) int {
	var i int
	colors := []color{Black, White, Red, Green, Blue, Yellow}
	hidden := new(Combination)
	for i := 0; i < 4; i++ {
		hidden.SetColorAt(colors[rand.Intn(len(colors))], i)
	}

	possible := Permutations()
	tries := make([]Try, 0, 10)
	for i = 0; i < 10; i++ {
		if verbose {
			fmt.Printf("Try %d: %d possible combinations\n",
				i+1, len(possible))
		}
		guess := possible[rand.Intn(len(possible))]
		if verbose {
			fmt.Printf("  Trying %v\n", guess)
		}
		eval := hidden.Evaluate(&guess)
		if verbose {
			fmt.Printf("  %d correct, %d present\n",
				eval.Correct(), eval.Present())
		}
		if eval.Correct() == 4 {
			if verbose {
				fmt.Printf("SOLVED!\n")
			}
			break
		}

		t := Try{guess, *eval}
		tries = append(tries, t)

		possible = Eliminate(possible, tries)
		if len(possible) < 1 {
			if verbose {
				fmt.Printf("No possibilities left\n")
			}
			break
		}
	}
	if verbose {
		fmt.Printf("Hidden was %v\n\n", *hidden)
	}
	return i + 1
}
