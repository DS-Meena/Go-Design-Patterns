package main

// Command interface
type Command interface {
	Execute()
	Undo()
}

// Concrete Command for writing text
type WriteCommand struct {
	text     string
	document *Document
}

func (w *WriteCommand) Execute() {
	w.document.Write(w.text)
}

func (w *WriteCommand) Undo() {
	w.document.Erase(len(w.text))
}
