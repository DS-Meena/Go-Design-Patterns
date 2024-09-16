package main

func main() {
	editor := NewTextEditor()
	editor.Write("Hello ") // Hello
	editor.Write("World!") // Hello World!
	editor.ShowDocument()
	editor.Undo() // Hello
	editor.ShowDocument()
	editor.Redo() // Hello World!
	editor.ShowDocument()
}
