package state

import "fmt"

type HasMoneyState struct {
	machine *VendingMachine
}

func (state *HasMoneyState) AddItem(_ int) error {
	return fmt.Errorf("item disperse in progress")
}

func (state *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item disperse in progress")
}

func (state *HasMoneyState) InsertMoney(_ int) error {
	return fmt.Errorf("item out of stock")
}

func (state *HasMoneyState) DisperseItem() error {
	state.machine.decrementItemCountByOne()

	if state.machine.itemCount == 0 {
		state.machine.SetState(state.machine.NoItem)
	} else {
		state.machine.SetState(state.machine.HasItem)
	}

	return nil
}
