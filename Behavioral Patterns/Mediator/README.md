# Mediator Design Pattern 🔗

The Mediator pattern is a behavioral design pattern that reduces coupling between components of a program by making them communicate indirectly, through a special mediator object. 🧠

## When to Use 🤔

- When you want to reduce chaotic dependencies between objects 🌀
- When you want to reuse individual components more easily 🔄
- When communication between objects becomes too complex 🕸️

It can be used in:
- online auction system
- airline management system
- chat application

## Example: Chat application

We have colleagues `User`, that won't interact with each other directly. `User` has a `ChatRoom` mediator, user will send message to mediator and mediator will show the message to other users.

```bash
go run *.go
[John smith]: Hi there!
[Jane smith]: Stop this
```

This example demonstrates how the Mediator pattern can be implemented, allowing for decoupled communication between users through the chat room.

This pattern is particularly useful in complex systems where multiple objects need to interact. It centralizes control, simplifies maintenance, and makes the system more flexible and extensible. 🚀