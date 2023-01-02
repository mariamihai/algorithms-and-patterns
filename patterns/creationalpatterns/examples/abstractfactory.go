package main

import (
	"code/patterns/creationalpatterns/abstractfactory"
	"fmt"
)

func abstractFactoryExample() {
	adidasFactory, _ := abstractfactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractfactory.GetSportsFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	fmt.Printf("Adidas shoe (logo %s, size %d)\n", adidasShoe.GetLogo(), adidasShoe.GetSize())
	fmt.Printf("Adidas shirt (logo %s, size %d)\n", adidasShirt.GetLogo(), adidasShirt.GetSize())

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	fmt.Printf("Nike shoe (logo %s, size %d)\n", nikeShoe.GetLogo(), nikeShoe.GetSize())
	fmt.Printf("Nike shirt (logo %s, size %d)\n\n", nikeShirt.GetLogo(), nikeShirt.GetSize())

	// Output
	// Adidas shoe (logo adidas, size 40)
	// Adidas shirt (logo adidas, size 40)
	// Nike shoe (logo nike, size 40)
	// Nike shirt (logo nike, size 40)
}
