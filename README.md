# Go-Design-Patterns

## What are design patterns?

- Degn patterns identify solutions to common programming problems, but they are not actual code or algorithms themselves.
- Pattesirns are a form of best practices for approaching a problem
- Usage of patterns can be significantly influenced by the language

Example -> patterns useful in object-oriented programming (OOP) may not work in non-OOP languages.

## Golang Features That Affect Design Patterns

- No support for traditional OOP classes like in C# or Java
    -   No static members or constructors - affects patterns like Singleton
- No support for class-based inheritance or method overloading
    - Affects patterns like Visitor

- No support for exceptions - error handling is via return values
    - Affects  patterns like Builder

- No support for abstract classes
    - Affects patterns like Abstract Factory and Bridge

## Design Pattern Categories

| Creational | Structural | Behavioral |
| --- | ---  | --- |
| Abstract Factory | **Adapter** | Chain of Responsibility |
| **Builder** | Bridge | Command |
| **Factory** | Composite | Interpreter |
| Lazy Initialization | Decorator | **Iterator** |
| Multiton | **Facade** | Mediator |
| Object Pool | Flyweight | Momento |
| Protype | Proxy | **Observer** |
| Singleton |  | Strategy |
|  |  | Visitor |

[Course on LinkedIn Learning](https://www.linkedin.com/learning/go-design-patterns/) by Joe Marini