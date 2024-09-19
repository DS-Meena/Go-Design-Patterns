# Memento Design Pattern 🕰️

The Memento Design Pattern is a behavioral pattern that allows you to capture and restore an object's internal state without violating encapsulation. It's like creating a snapshot 📸 of an object's state that can be restored later.

## When to Use 🤔

- When you need to implement undo/redo functionality 🔄
- When you want to create checkpoints of an object's state 🚩
- When direct access to an object's fields/getters/setters violates its encapsulation 🔒

## Example

In this example, we have an Originator (the object whose state we want to save), a Memento (which stores the state), and a Caretaker (which manages the saved states). The Originator can create Mementos and restore its state from them, while the Caretaker keeps track of multiple Mementos. This pattern allows us to implement undo functionality or save checkpoints of an object's state easily. 🎉

```bash
go run *.go
Originator current state: State #1
Originator current state: State #2
Originator current state: State #1
```