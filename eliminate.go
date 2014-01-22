package mastermind

type Try struct {
	comb Combination
	eval Evaluation
}

func Eliminate(possible []Combination, tries []Try) []Combination {
	res := make([]Combination, 0, len(possible))
Poss:
	for _, p := range possible {
		for _, t := range tries {
			if *p.Evaluate(&t.comb) != t.eval {
				continue Poss
			}
		}
		res = append(res, p)
	}
	return res
}
