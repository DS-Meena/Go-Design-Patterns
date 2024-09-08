package main

import "fmt"

// Handler interface
type SupportHandler interface {
	SetNext(handler SupportHandler)
	Handle(request string)
}

// Concrete handler
type FirstLevelSupport struct {
	next SupportHandler
}

func (f *FirstLevelSupport) SetNext(handler SupportHandler) {
	f.next = handler
}

func (f *FirstLevelSupport) Handle(request string) {
	if request == "General Inquiry" {
		fmt.Println("First Level Support: Handling general inquiry")
	} else if f.next != nil {
		f.next.Handle(request)
	}
}

// Another concrete Handler
type SecondLevelSupport struct {
	next SupportHandler
}

func (s *SecondLevelSupport) SetNext(handler SupportHandler) {
	s.next = handler
}

func (s *SecondLevelSupport) Handle(request string) {
	if request == "Technical Issue" {
		fmt.Println("Second Level Support: Handling technical issue")
	} else if s.next != nil {
		s.next.Handle(request)
	}
}
