package mastermind

type Evaluation struct {
	correct int
	present int
}

func (e *Evaluation) Correct() int {
	return e.correct
}

func (e *Evaluation) Present() int {
	return e.present
}

func (e *Evaluation) IncCorrect() {
	e.correct++
}

func (e *Evaluation) IncPresent() {
	e.present++
}
