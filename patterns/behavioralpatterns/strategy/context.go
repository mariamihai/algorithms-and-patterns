package strategy

type Context struct {
	strategy Strategy
	A, B     int
}

func InitContext(a, b int, strategy Strategy) *Context {
	return &Context{
		strategy: strategy,
		A:        a,
		B:        b,
	}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() int {
	return c.strategy.executeOperation(c.A, c.B)
}
