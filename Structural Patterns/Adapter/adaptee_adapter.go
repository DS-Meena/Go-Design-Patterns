package main

import "fmt"

// Target interface
type Target interface {
	Request() string
}

// Adaptee
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Adaptee's specific Request"
}

// Adapter
type Adapter struct {
	adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return fmt.Sprintf("Adapater: (Translated) %s", a.adaptee.SpecificRequest())
}
