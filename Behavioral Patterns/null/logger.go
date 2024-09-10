package main

import "fmt"

// Interface
type Logger interface {
	Log(message string)
}

// Concrete implementation
type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(message string) {
	fmt.Println("Logging: ", message)
}

// Null object implementation
type NullLogger struct{}

func (l *NullLogger) Log(message string) {
	// Do nothing or default behavior
}
