package state

import "fmt"

// For testing purposes a lot of things are exported or not properly encapsulated.

type VendingMachine struct {
	currentState State

	HasItem       State
	ItemRequested State
	HasMoney      State
	NoItem        State

	itemCount int
	itemPrice int
}

func NewVendingMachine() *VendingMachine {
	machine := &VendingMachine{}

	hasItemState := &HasItemState{
		machine: machine,
	}

	hasMoneyState := &HasMoneyState{
		machine: machine,
	}

	itemRequestedState := &ItemRequestedState{
		machine: machine,
	}

	noItemState := &NoItemState{
		machine: machine,
	}

	machine.SetState(hasItemState)

	machine.HasItem = hasItemState
	machine.HasMoney = hasMoneyState
	machine.ItemRequested = itemRequestedState
	machine.NoItem = noItemState

	return machine
}

func (vm *VendingMachine) GetState() State {
	return vm.currentState
}

func (vm *VendingMachine) SetState(state State) {
	vm.currentState = state
}

func (vm *VendingMachine) SetItems(count int, money int) {
	vm.itemCount = count
	vm.itemPrice = money
}

func (vm *VendingMachine) GetItems() (int, int) {
	return vm.itemCount, vm.itemPrice
}

func (vm *VendingMachine) incrementItemCount(nrItems int) {
	vm.itemCount += nrItems
}

func (vm *VendingMachine) decrementItemCountByOne() {
	vm.itemCount -= 1
}

func (vm *VendingMachine) AddItem(count int) error {
	fmt.Println("ACTION: Adding items")
	return vm.currentState.AddItem(count)
}

func (vm *VendingMachine) RequestItem() error {
	fmt.Println("ACTION: Requesting item")
	return vm.currentState.RequestItem()
}

func (vm *VendingMachine) InsertMoney(money int) error {
	fmt.Println("ACTION: Inserting money")
	return vm.currentState.InsertMoney(money)
}

func (vm *VendingMachine) DisperseItem() error {
	fmt.Println("ACTION: Dispersing item")
	return vm.currentState.DisperseItem()
}
