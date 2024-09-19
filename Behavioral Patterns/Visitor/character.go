package main

// Character interface
type Character interface {
	Accept(visitor CharacterVisitor)
}

// Concrete character types
type Soldier struct{}

func (s *Soldier) Accept(visitor CharacterVisitor) {
	visitor.VisitSoldier(s)
}

type Tank struct{}

func (t *Tank) Accept(visitor CharacterVisitor) {
	visitor.VisitTank(t)
}

type Healer struct{}

func (h *Healer) Accept(visitor CharacterVisitor) {
	visitor.VisitHealer(h)
}
