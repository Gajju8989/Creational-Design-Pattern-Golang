package main

import (
	"design_patterns/prototype"
	"fmt"
)

func main() {
	original := &prototype.ConcretePrototype{Value: 42}
	clone := original.Clone()

	fmt.Println(original)
	fmt.Println(clone)
	fmt.Printf("Are they the same object? %v\n", original == clone)
}
