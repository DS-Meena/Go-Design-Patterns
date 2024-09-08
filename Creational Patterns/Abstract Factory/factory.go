package main

// Abstract Factory
type CarFactory interface {
	CreateCar() Car
	CreateEngine() Engine
}

// Conrete Factories
type SportsCarFactory struct{}

func (s *SportsCarFactory) CreateCar() Car {
	return &SedanCar{}
}

func (s *SportsCarFactory) CreateEngine() Engine {
	return &SportsEngine{}
}

type LuxuryCarFactory struct{}

func (l *LuxuryCarFactory) CreateCar() Car {
	return &LuxuryCar{}
}
func (l *LuxuryCarFactory) CreateEngine() Engine {
	return &LuxuryEngine{}
}
