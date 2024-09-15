package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
}

// concrete class with Clone method
func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: c.name}
}

func main() {
	original := &ConcretePrototype{name: "David"}
	clone := original.Clone()

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Clone: %v\n", clone)
}
