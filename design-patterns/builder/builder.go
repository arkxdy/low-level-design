package main

// 2. The Builder Interface
type Builder interface {
	SetPartA(val string) Builder
	SetPartB(val string) Builder
	SetPartC(val string) Builder
	Build() *ComplexObject
}

// 3. The Concrete Builder
type concreteBuilder struct {
	obj *ComplexObject
}

func NewBuilder() Builder {
	return &concreteBuilder{obj: &ComplexObject{}}
}

func (b *concreteBuilder) SetPartA(val string) Builder { b.obj.PartA = val; return b }
func (b *concreteBuilder) SetPartB(val string) Builder { b.obj.PartB = val; return b }
func (b *concreteBuilder) SetPartC(val string) Builder { b.obj.PartC = val; return b }

func (b *concreteBuilder) Build() *ComplexObject {
	res := b.obj
	b.obj = &ComplexObject{} // Reset for the next build
	return res
}
