package main

// Receiver
type Document struct {
	content string
}

func (d *Document) Write(text string) {
	d.content += text
}

func (d *Document) Erase(length int) {
	d.content = d.content[:len(d.content)-length]
}

func (d *Document) String() string {
	return d.content
}
