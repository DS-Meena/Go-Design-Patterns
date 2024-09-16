# Flyweight Pattern 🪶

The Flyweight pattern is a structural design pattern that aims to minimize memory usage or computational expenses by sharing as much as possible with similar objects. It's particularly useful when dealing with a large number of objects that have some shared state. 🧠💾

## Key components of the Flyweight pattern:

- Flyweight: The interface through which flyweights can receive and act on extrinsic state. 🔄
- Concrete Flyweight: Implements the Flyweight interface and stores intrinsic state. 🧱
- Flyweight Factory: Creates and manages flyweight objects. 🏭

### When to Use

1. When Memory is limited 📚
2. When objects share data: 🔁
    - Intrinsic data: shared among objects and doesn't change for different objects.
    - Extrinsic data: changes based on client input and differs from one object to another
3. Creation of object is expensive. 💰

### How it solves

- From Object, remove all the Extrinsic data and keep Intrinsic data (`Conrete Flyweight`)
- This flyweight object can be immutable.
- Extrinsic datat can be passed to the flyweight object in method parameter (`Interface method`).
- Once the flyweight object is created, it is cached and reused whenever required (`Flyweight factory`).

## Example: Text Editor with Character Objects

Imagine a text editor that creates an object for each character. For a large document, this could consume a lot of memory. Using the Flyweight pattern, we can share character objects for each letter of the alphabet.

```golang
go run *.go
Character: H with Font: Arial and Size: 12
Character: e with Font: Arial and Size: 12
Character: l with Font: Arial and Size: 12
Character: l with Font: Arial and Size: 12
Character: o with Font: Arial and Size: 12
Character: , with Font: Arial and Size: 12
Character:   with Font: Arial and Size: 12
Character: F with Font: Arial and Size: 12
Character: l with Font: Arial and Size: 12
Character: y with Font: Arial and Size: 12
Character: w with Font: Arial and Size: 12
Character: e with Font: Arial and Size: 12
Character: i with Font: Arial and Size: 12
Character: g with Font: Arial and Size: 12
Character: h with Font: Arial and Size: 12
Character: t with Font: Arial and Size: 12
Character: ! with Font: Arial and Size: 12
Character:   with Font: Arial and Size: 12
Character: 🪶 with Font: Arial and Size: 12
```

The CharacterFactory ensures that only one object is created for each unique character. This significantly reduces memory usage for large texts. 🚀💡