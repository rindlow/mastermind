package mastermind

func Permutations() []Combination {
	perm := make([]Combination, 0, 6*6*6*6)
	colors := []color{Black, White, Red, Green, Blue, Yellow}
	for _, a := range colors {
		for _, b := range colors {
			for _, c := range colors {
				for _, d := range colors {
					perm = append(perm,
						Combination{a, b, c, d})
				}
			}
		}
	}
	return perm
}
