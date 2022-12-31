package state_test

import (
	"code/patterns/behavioralpatterns/state"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStates(t *testing.T) {
	machine := state.NewVendingMachine()

	testCases := []struct {
		description       string
		initialItemCount  int
		initialItemPrice  int
		expectedItemCount int
		expectedItemPrice int
		initialState      state.State
		expectedState     state.State
		function          func() error
		errorResult       error
	}{
		// --- Add item
		{
			description:       "Add item when in hasItemState",
			initialItemCount:  10,
			initialItemPrice:  10,
			expectedItemCount: 20,
			expectedItemPrice: 10,
			initialState:      machine.HasItem,
			expectedState:     machine.HasItem,
			function: func() error {
				return machine.AddItem(10)
			},
			errorResult: nil,
		},
		{
			description:       "Add item when in noItemState",
			initialItemCount:  0,
			initialItemPrice:  10,
			expectedItemCount: 10,
			expectedItemPrice: 10,
			initialState:      machine.NoItem,
			expectedState:     machine.HasItem,
			function: func() error {
				return machine.AddItem(10)
			},
			errorResult: nil,
		},
		{
			description:       "Add item (invalid) when in itemRequestedState",
			initialItemCount:  10,
			initialItemPrice:  10,
			expectedItemCount: 10,
			expectedItemPrice: 10,
			initialState:      machine.ItemRequested,
			expectedState:     machine.ItemRequested,
			function: func() error {
				return machine.AddItem(10)
			},
			errorResult: fmt.Errorf("item disperse in progress"),
		},
		{
			description:       "Add item (invalid) when in hasMoneyState",
			initialItemCount:  10,
			initialItemPrice:  10,
			expectedItemCount: 10,
			expectedItemPrice: 10,
			initialState:      machine.HasMoney,
			expectedState:     machine.HasMoney,
			function: func() error {
				return machine.AddItem(10)
			},
			errorResult: fmt.Errorf("item disperse in progress"),
		},
		// --- Request item
		{
			description:       "Request item when in hasItemState and no items",
			initialItemCount:  0,
			initialItemPrice:  10,
			expectedItemCount: 0,
			expectedItemPrice: 10,
			initialState:      machine.HasItem,
			expectedState:     machine.NoItem,
			function: func() error {
				return machine.RequestItem()
			},
			errorResult: fmt.Errorf("no item present"),
		},
		{
			description:       "Request item when in hasItemState and with items",
			initialItemCount:  10,
			initialItemPrice:  10,
			expectedItemCount: 10,
			expectedItemPrice: 10,
			initialState:      machine.HasItem,
			expectedState:     machine.ItemRequested,
			function: func() error {
				return machine.RequestItem()
			},
			errorResult: nil,
		},
		{
			description:       "Request item when in noItemState",
			initialItemCount:  0,
			initialItemPrice:  10,
			expectedItemCount: 0,
			expectedItemPrice: 10,
			initialState:      machine.NoItem,
			expectedState:     machine.NoItem,
			function: func() error {
				return machine.RequestItem()
			},
			errorResult: fmt.Errorf("item out of stock"),
		},
		{
			description:       "Request item when in itemRequestedState",
			initialItemCount:  10,
			initialItemPrice:  10,
			expectedItemCount: 10,
			expectedItemPrice: 10,
			initialState:      machine.ItemRequested,
			expectedState:     machine.ItemRequested,
			function: func() error {
				return machine.RequestItem()
			},
			errorResult: fmt.Errorf("item already requested"),
		},
		{
			description:       "Request item when in hasMoneyState",
			initialItemCount:  10,
			initialItemPrice:  10,
			expectedItemCount: 10,
			expectedItemPrice: 10,
			initialState:      machine.HasMoney,
			expectedState:     machine.HasMoney,
			function: func() error {
				return machine.RequestItem()
			},
			errorResult: fmt.Errorf("item disperse in progress"),
		},
		// --- Insert money
		//...

		// --- disperse item
		//...
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			machine.SetItems(testCase.initialItemCount, testCase.initialItemPrice)
			machine.SetState(testCase.initialState)

			testResult := testCase.function()

			if testCase.errorResult != nil {
				isEqual := assert.EqualErrorf(t, testCase.errorResult, testResult.Error(), "%s: expected %v, got %v", testCase.description, testCase.errorResult, testResult)
				if !isEqual {
					t.Errorf("is not equal")
				}
			}

			assert.IsType(t, testCase.expectedState, machine.GetState())
			resultCount, resultPrice := machine.GetItems()
			assert.Equal(t, testCase.expectedItemCount, resultCount)
			assert.Equal(t, testCase.expectedItemPrice, resultPrice)
		})
	}
}
