package main

import (
	"code/patterns/creationalpatterns/builder"
	"fmt"
)

func builderExample() {
	normalBuilder, _ := builder.GetBuilder("normal")
	iglooBuilder, _ := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal house (window - %s, door - %s, floor - %d)\n", normalHouse.WindowType, normalHouse.DoorType, normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("Igloo house (window - %s, door - %s, floor - %d)\n\n", iglooHouse.WindowType, iglooHouse.DoorType, iglooHouse.Floor)

	// Output
	// Normal house (window - wooden window, door - wooden door, floor - 3)
	// Igloo house (window - snow window, door - snow door, floor - 3)
}
