package main

import (
	"mastermind"
	"math/rand"
	"os"
	"fmt"
)

func main() {
	var (
		best = 10
		worst = 0
	)
	rand.Seed(int64(os.Getpid()))
	for i := 0; i < 10; i++ {
		t := mastermind.AutoSolve(true)
		if t < best {
			best = t
		}
		if t > worst {
			worst = t
		}
	}
	fmt.Printf("Seed: %d, best: %d, worst: %d\n", os.Getpid(), best, worst)
}
