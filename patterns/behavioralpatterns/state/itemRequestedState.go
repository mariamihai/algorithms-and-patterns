package state

import "fmt"

type ItemRequestedState struct {
	machine *VendingMachine
}

func (state *ItemRequestedState) AddItem(_ int) error {
	return fmt.Errorf("item disperse in progress")
}

func (state *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (state *ItemRequestedState) InsertMoney(money int) error {
	if money < state.machine.itemPrice {
		return fmt.Errorf("Not enough money. %d more needed.", state.machine.itemPrice-money)
	}

	state.machine.SetState(state.machine.HasMoney)

	return nil
}

func (state *ItemRequestedState) DisperseItem() error {
	return fmt.Errorf("insert money first")
}
