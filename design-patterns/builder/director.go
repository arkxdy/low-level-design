package main

// 4. The Director
// The Director defines the "recipes" for different product variations.
type Director struct {
	builder Builder
}

func (d *Director) ConstructMinimal() *ComplexObject {
	return d.builder.SetPartA("Alpha").Build()
}

func (d *Director) ConstructFull() *ComplexObject {
	return d.builder.SetPartA("Alpha").SetPartB("Beta").SetPartC("Gamma").Build()
}
