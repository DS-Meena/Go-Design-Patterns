package main

import "fmt"

func main() {
	coffee := &SimpleCoffee{}
	fmt.Printf("Cost :%d; Description: %s\n", coffee.GetCost(), coffee.GetDescription())

	// coffee + milk
	coffeeWithMilk := &Milk{
		CoffeeDecorator: CoffeeDecorator{
			Coffee: coffee},
	}
	fmt.Printf("Cost :%d; Description: %s\n", coffeeWithMilk.GetCost(), coffeeWithMilk.GetDescription())

	// coffee + milk + whip
	coffeeWithMilkAndWhip := &Whip{
		CoffeeDecorator: CoffeeDecorator{
			Coffee: coffeeWithMilk,
		},
	}
	fmt.Printf("Cost :%d; Description: %s\n", coffeeWithMilkAndWhip.GetCost(), coffeeWithMilkAndWhip.GetDescription())

}
