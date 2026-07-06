package main

import "fmt"

func main() {
	builder := NewBuilder()
	director := &Director{builder: builder}

	obj1 := director.ConstructMinimal()
	obj2 := director.ConstructFull()

	fmt.Printf("Minimal: %+v\nFull: %+v\n", obj1, obj2)
}
