# Chain of Responsibility Design Pattern 🔗

The Chain of Responsibility is a behavioral design pattern that allows passing requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain. 🤝

### Purpose 🎯

The main purpose of this pattern is to decouple the sender of a request from its receivers by giving multiple objects a chance to handle the request. It promotes loose coupling and allows you to add or remove responsibilities dynamically. 🔀

### When to Use 🕒

- When you want to decouple a request's sender and receiver 📤📥
- When multiple objects may handle a request, and the handler isn't known in prior 🤔
- When you want to issue a request to one of several objects without specifying the receiver explicitly 🎭

### Benefits 🌟

- Reduces coupling between components 🔓
- Increases flexibility in assigning responsibilities to objects 🔄
- Allows dynamic changes to the chain at runtime 🔀

Remember, while the Chain of Responsibility pattern can be powerful, it's important to use it judiciously to avoid creating overly complex chains that might be hard to debug or maintain. 🧠

## Examples: -

1. ATM machine
2. Vending machine 
3. Design Logger (error, info, debug)

### 4. Ticket system

In our example, we have a chain of support handlers. Each handler decides whether it can process the request or should pass it to the next handler in the chain. This allows for flexible and decoupled handling of different types of support requests. 🛠️

Output:
```
go run *.go
First Level Support: Handling general inquiry
Second Level Support: Handling technical issue
```