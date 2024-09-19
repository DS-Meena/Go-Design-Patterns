# Visitor Pattern 🏃‍♂️👀

The Visitor pattern is a behavioral design pattern that allows adding new operations to existing object structures without modifying them. It separates the algorithm from an object structure by moving the algorithm into a separate class. 🧩🔀

## Example Scenario 🎭

Imagine a game with different character types (Soldier, Tank, Healer). We want to implement various operations on these characters without modifying their classes. 🎮👥

```bash
go run *.go
Boosting Soldier's health by 10%
Boosting Tank's health by 20%
Boosting Healer's health by 15%
```

## When to Use 🤔

- When you need to perform operations on a complex object structure, such as an abstract syntax tree or a composite pattern. 🌳
- When you want to add new operations to existing object structures without modifying them. ➕🔒
- When you have many distinct and unrelated operations to perform on objects in a structure. 🔀🔧

### Single vs Double Dispatch 🚀

**Single Dispatch:** The method to be invoked is determined by the runtime type of a single object. Most object-oriented languages use this by default. 1️⃣🎯

**Double Dispatch:** The method invoked is determined by the runtime type of two objects. The Visitor pattern uses double dispatch to determine which visit_* method to call based on both the visitor and the element types. 2️⃣🎯

```golang
character.Accept(healthBooster)
// Here, method do be called is decided by character and healthBooster.
// character.Accept()
// visitor.VisitSoldier(s)  
```

## Visitor vs Strategy Pattern 🆚

**Visitor Pattern:**

- Separates an algorithm from an object structure 🏗️
- Allows adding new operations without modifying the objects 🆕
- Uses double dispatch 2️⃣🚀
- Best for when you have a fixed set of element classes but often need to add new operations 📊

**Strategy Pattern:**

- Defines a family of algorithms and makes them interchangeable 🔄
- Encapsulates each algorithm 📦
- Uses single dispatch 1️⃣🚀
- Best for when you need to switch between different algorithms dynamically 🔀

The main difference is that the Visitor pattern is about adding operations to an object structure, while the Strategy pattern is about making algorithms interchangeable. 🔑🧠