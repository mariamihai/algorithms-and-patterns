package state

import "fmt"

type NoItemState struct {
	machine *VendingMachine
}

func (state *NoItemState) AddItem(count int) error {
	state.machine.incrementItemCount(count)

	state.machine.SetState(state.machine.HasItem)

	return nil
}

func (state *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (state *NoItemState) InsertMoney(_ int) error {
	return fmt.Errorf("item out of stock")
}

func (state *NoItemState) DisperseItem() error {
	return fmt.Errorf("item out of stock")
}
