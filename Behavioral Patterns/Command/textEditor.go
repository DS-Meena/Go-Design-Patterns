package main

import (
	"fmt"
)

// Invoker
type TextEditor struct {
	document  *Document
	history   []Command
	redoStack []Command
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		document: &Document{},
	}
}

func (t *TextEditor) Write(text string) {
	command := &WriteCommand{
		text:     text,
		document: t.document,
	}
	t.executeCommand(command)
}

func (t *TextEditor) Undo() {
	if len(t.history) > 0 {
		command := t.history[len(t.history)-1]
		t.history = t.history[:len(t.history)-1]
		command.Undo()
		t.redoStack = append(t.redoStack, command)
	}
}

func (t *TextEditor) Redo() {
	if len(t.redoStack) > 0 {
		command := t.redoStack[len(t.redoStack)-1]
		t.redoStack = t.redoStack[:len(t.redoStack)-1]
		t.executeCommand(command)
	}
}

func (t *TextEditor) executeCommand(command Command) {
	command.Execute()
	t.history = append(t.history, command)
	t.redoStack = nil
}

func (t *TextEditor) ShowDocument() {
	fmt.Println(t.document)
}
