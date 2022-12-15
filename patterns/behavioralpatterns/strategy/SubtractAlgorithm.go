package strategy

type Subtraction struct {
	//...
}

func (subtr Subtraction) executeOperation(a, b int) int {
	return a - b
}
