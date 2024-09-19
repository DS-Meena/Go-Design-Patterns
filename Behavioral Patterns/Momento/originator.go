package main

// Momento stores the state of the Originator
type Momento struct {
	state string
}

// Originator is the object whose state we want to save
type Originator struct {
	state string
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) CreateMomento() *Momento {
	return &Momento{state: o.state}
}

func (o *Originator) RestoreState(m *Momento) {
	o.state = m.state
}
