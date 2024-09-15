package main

import "fmt"

// Implementor
type Device interface {
	TurnOn()
	TurnOff()
	SetChannel(channel int)
}

// Concrete Implementors
type TV struct{}

func (t *TV) TurnOn()                { fmt.Println("TV turned on") }
func (t *TV) TurnOff()               { fmt.Println("TV turned off") }
func (t *TV) SetChannel(channel int) { fmt.Printf("TV channel set to %d\n", channel) }

type Radio struct{}

func (r *Radio) TurnOn()                { fmt.Println("Radio turned on") }
func (r *Radio) TurnOff()               { fmt.Println("Radio turned off") }
func (r *Radio) SetChannel(channel int) { fmt.Printf("Radio frequency set to %d\n", channel) }
