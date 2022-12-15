package strategy

type Strategy interface {
	executeOperation(a, b int) int
}
