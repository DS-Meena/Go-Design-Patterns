package main

// Caretake manages and safekeeps the Momentos
type Caretaker struct {
	momentos []*Momento // can also use stack
}

func (c *Caretaker) AddMomento(m *Momento) {
	c.momentos = append(c.momentos, m)
}

func (c *Caretaker) GetMomento(index int) *Momento {
	return c.momentos[index]
}
