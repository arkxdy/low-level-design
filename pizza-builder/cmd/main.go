package main

import (
	"fmt"
	"pizzabuilder/internal"
)

func main() {
	builder := internal.NewPizzaBuilder()
	pizza := builder.
		SetSize(internal.M).
		SetCrust(internal.Stuffed).
		AddTopping(internal.Olives).
		AddTopping(internal.Onions).
		AddTopping(internal.Mushrooms).
		AddSauce(internal.Pesto).
		AddCheese(internal.Vegan).
		Build()
	fmt.Println(pizza.String())
}
