package main

import "fmt"

func main() {
	originator := &Originator{}
	caretaker := &Caretaker{}

	// set initial state
	originator.SetState("State #1")
	fmt.Printf("Originator current state: %s\n", originator.GetState())

	// Save to momento
	caretaker.AddMomento(originator.CreateMomento())

	// change the state
	originator.SetState("State #2")
	fmt.Printf("Originator current state: %s\n", originator.GetState())

	// Restor efrom momento
	originator.RestoreState(caretaker.GetMomento(0))
	fmt.Printf("Originator current state: %s\n", originator.GetState())

}
