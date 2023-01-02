package main

import (
	"code/patterns/creationalpatterns/factorymethod"
	"fmt"
)

// Simple factory implementation

func factoryMethodExample() {
	ak47, _ := factorymethod.GetGun("ak47")
	musket, _ := factorymethod.GetGun("musket")

	fmt.Printf("Gun %s (power %d)\n", ak47.GetName(), ak47.GetPower())
	fmt.Printf("Gun %s (power %d)\n", musket.GetName(), musket.GetPower())
}
