package main

// Intermediary Decorator
type CoffeeDecorator struct {
	Coffee Coffee
}

func (c *CoffeeDecorator) GetCost() int {
	return c.Coffee.GetCost()
}

func (c *CoffeeDecorator) GetDescription() string {
	return c.Coffee.GetDescription()
}

// Concrete Decorators
type Milk struct {
	CoffeeDecorator
}

func (m *Milk) GetCost() int {
	return m.CoffeeDecorator.GetCost() + 5
}

func (m *Milk) GetDescription() string {
	return m.CoffeeDecorator.GetDescription() + ", milk"
}

type Whip struct {
	CoffeeDecorator
}

func (w *Whip) GetCost() int {
	return w.CoffeeDecorator.GetCost() + 7
}

func (w *Whip) GetDescription() string {
	return w.CoffeeDecorator.GetDescription() + ", whip"
}
