package strategy

type Addition struct {
	//...
}

func (add Addition) executeOperation(a, b int) int {
	return a + b
}
