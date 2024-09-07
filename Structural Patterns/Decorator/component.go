package main

// Component interface
type Coffee interface {
	GetCost() int
	GetDescription() string
}

// Concrete component
type SimpleCoffee struct{}

func (c *SimpleCoffee) GetCost() int {
	return 10
}

func (c *SimpleCoffee) GetDescription() string {
	return "Simple coffee"
}
