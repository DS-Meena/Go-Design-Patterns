package main

import "fmt"

// Visitor interface
type CharacterVisitor interface {
	VisitSoldier(*Soldier)
	VisitTank(*Tank)
	VisitHealer(*Healer)
}

// Concrete visitor
type HealthBoostVisitor struct{}

func (v *HealthBoostVisitor) VisitSoldier(s *Soldier) {
	fmt.Println("Boosting Soldier's health by 10%")
}

func (v *HealthBoostVisitor) VisitTank(t *Tank) {
	fmt.Println("Boosting Tank's health by 20%")
}

func (v *HealthBoostVisitor) VisitHealer(h *Healer) {
	fmt.Println("Boosting Healer's health by 15%")
}
