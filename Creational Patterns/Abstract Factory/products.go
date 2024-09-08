package main

import "fmt"

// Abstract Product interfaces
type Car interface {
	Drive()
}

type Engine interface {
	Start()
}

// Concrete products
type SedanCar struct{}

func (s *SedanCar) Drive() {
	fmt.Println("Driving a Sedan car")
}

type LuxuryCar struct{}

func (l *LuxuryCar) Drive() {
	fmt.Println("Driving a luxury car")
}

type SportsEngine struct{}

func (s *SportsEngine) Start() {
	fmt.Println("Sports Engine started")
}

type LuxuryEngine struct{}

func (l *LuxuryEngine) Start() {
	fmt.Println("Luxury Engine started")
}
