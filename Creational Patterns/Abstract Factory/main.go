package main

import "fmt"

func OrderCar(factory CarFactory) {
	car := factory.CreateCar()
	engine := factory.CreateEngine()

	engine.Start()
	car.Drive()
}

func main() {
	sportsFactory := &SportsCarFactory{}
	luxuryFactory := &LuxuryCarFactory{}

	fmt.Println("Ordering a sports car...")
	OrderCar(sportsFactory)

	fmt.Println("Ordering a Luxury car...")
	OrderCar(luxuryFactory)
}
