package internal

import "fmt"

type Pizza struct {
	Size     Size
	Crust    CrustType
	Toppings []Topping
	Sauce    Sauce
	Cheese   Cheese
	Price    Price
}

func (p *Pizza) String() string {
	return fmt.Sprintf("%s Pizza on %s crust and %s cheese with %v ($%d)",
		p.Size, p.Crust, p.Cheese, p.Toppings, p.Price)
}

type Price int

type PizzaSize struct {
	Size  Size
	Price Price
}

type Size string
type Topping string
type CrustType string
type Sauce string
type Cheese string

const (
	S Size = "Small"
	M Size = "Medium"
	L Size = "Large"
)

const (
	Thin    CrustType = "thin"
	Thick   CrustType = "thick"
	Stuffed CrustType = "stuffed"
)

const (
	Pepperoni Topping = "pepperoni"
	Mushrooms Topping = "mushrooms"
	Onions    Topping = "onions"
	Olives    Topping = "olives"
)

const (
	Tomato Sauce = "tomato"
	Pesto  Sauce = "pesto"
	White  Sauce = "white"
)

const (
	Mozzerella Cheese = "mozzerella"
	Cheddar    Cheese = "cheddar"
	Vegan      Cheese = "vegan"
)
