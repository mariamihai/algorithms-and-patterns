package state

import "fmt"

type HasItemState struct {
	machine *VendingMachine
}

func (state *HasItemState) AddItem(count int) error {
	state.machine.incrementItemCount(count)

	return nil
}

func (state *HasItemState) RequestItem() error {
	if state.machine.itemCount == 0 {
		state.machine.SetState(state.machine.NoItem)

		return fmt.Errorf("no item present")
	}

	state.machine.SetState(state.machine.ItemRequested)

	return nil
}

func (state *HasItemState) InsertMoney(_ int) error {
	return fmt.Errorf("no item selected")
}

func (state *HasItemState) DisperseItem() error {
	return fmt.Errorf("no item selected")
}
