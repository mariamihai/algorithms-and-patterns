package main

import (
	"code/patterns/behavioralpatterns/state"
	"fmt"
)

func stateExample() {
	// Context
	machine := state.NewVendingMachine()
	machine.SetItems(1, 10)

	err := machine.RequestItem()
	if err != nil {
		fmt.Println("Couldn't request item")
	}

	err = machine.InsertMoney(5)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.InsertMoney(10)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.DisperseItem()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.RequestItem()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.AddItem(10)
	if err != nil {
		fmt.Println(err)
	}

	err = machine.DisperseItem()
	if err != nil {
		fmt.Println(err)
	}

	err = machine.InsertMoney(5)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	//ACTION: Requesting item
	//ACTION: Inserting money
	//Not enough money. 5 more needed.
	//ACTION: Inserting money
	//ACTION: Dispersing item
	//ACTION: Requesting item
	//item out of stock
	//ACTION: Adding items
	//ACTION: Dispersing item
	//no item selected
	//ACTION: Inserting money
	//no item selected
}
