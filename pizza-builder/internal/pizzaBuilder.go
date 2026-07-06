package internal

type PizzaBuilder interface {
	SetSize(size Size) PizzaBuilder
	SetCrust(crust CrustType) PizzaBuilder
	AddTopping(topping Topping) PizzaBuilder
	AddSauce(sauce Sauce) PizzaBuilder
	AddCheese(cheese Cheese) PizzaBuilder
	Build() *Pizza
}

type pizzaBuild struct {
	pizza *Pizza
}

// AddCheese implements [PizzaBuilder].
func (p *pizzaBuild) AddCheese(cheese Cheese) PizzaBuilder {
	p.pizza.Cheese = cheese
	switch cheese {
	case Mozzerella:
		p.pizza.Price += 1
	case Cheddar:
		p.pizza.Price += 2
	case Vegan:
		p.pizza.Price += 3
	}
	return p
}

// AddSauce implements [PizzaBuilder].
func (p *pizzaBuild) AddSauce(sauce Sauce) PizzaBuilder {
	p.pizza.Sauce = sauce
	switch sauce {
	case Tomato:
		p.pizza.Price += 2
	case Pesto:
		p.pizza.Price += 3
	case White:
		p.pizza.Price += 1
	default:
		p.pizza.Price += 1
	}
	return p
}

// AddTopping implements [PizzaBuilder].
func (p *pizzaBuild) AddTopping(topping Topping) PizzaBuilder {
	p.pizza.Toppings = append(p.pizza.Toppings, topping)
	switch topping {
	case Mushrooms:
		p.pizza.Price += 2
	case Pepperoni:
		p.pizza.Price += 2
	case Onions:
		p.pizza.Price += 1
	case Olives:
		p.pizza.Price += 2
	}
	return p
}

// Build implements [PizzaBuilder].
func (p *pizzaBuild) Build() *Pizza {
	return p.pizza
}

// SetCrust implements [PizzaBuilder].
func (p *pizzaBuild) SetCrust(crust CrustType) PizzaBuilder {
	p.pizza.Crust = crust
	switch crust {
	case Thick:
		p.pizza.Price += 2
	case Thin:
		p.pizza.Price += 1
	case Stuffed:
		p.pizza.Price += 3
	default:
		p.pizza.Price += 1
	}
	return p
}

// SetSize implements [PizzaBuilder].
func (p *pizzaBuild) SetSize(size Size) PizzaBuilder {
	p.pizza.Size = size
	switch size {
	case L:
		p.pizza.Price += 5
	case M:
		p.pizza.Price += 10
	case S:
		p.pizza.Price += 15
	default:
		p.pizza.Price += 10
	}
	return p
}

func NewPizzaBuilder() PizzaBuilder {
	return &pizzaBuild{pizza: &Pizza{Price: 0}}
}
