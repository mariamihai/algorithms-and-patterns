package strategy_test

import (
	"code/patterns/behavioralpatterns/strategy"
	"fmt"
)

func ExampleExecuteStrategy() {
	add := strategy.Addition{}
	ctx := strategy.InitContext(1, 2, add)
	fmt.Printf("Executing: 1 + 2 = %d\n", ctx.ExecuteStrategy())

	subtract := strategy.Subtraction{}
	ctx.SetStrategy(subtract)
	fmt.Printf("Executing: 1 - 2 = %d\n", ctx.ExecuteStrategy())

	// Output:
	// Executing: 1 + 2 = 3
	// Executing: 1 - 2 = -1
}
