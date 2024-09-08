package main

import "fmt"

// Image interface
type Image interface {
	Display()
}

// Real Image is the expensive object
type RealImage struct {
	filename string
}

func (r *RealImage) Display() {
	fmt.Printf("Displaying %s\n", r.filename)
}

func (r *RealImage) LoadFromDisk() {
	fmt.Printf("Loading %s from disk\n", r.filename)
}
