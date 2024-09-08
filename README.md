# Go-Design-Patterns

## What are design patterns?

- Degn patterns identify solutions to common programming problems, but they are not actual code or algorithms themselves.
- Pattesirns are a form of best practices for approaching a problem
- Usage of patterns can be significantly influenced by the language

Example -> patterns useful in object-oriented programming (OOP) may not work in non-OOP languages.

## Golang Features That Affect Design Patterns 😓

- No support for traditional OOP classes like in C# or Java
    -   No static members or constructors - affects patterns like Singleton
- No support for class-based inheritance or method overloading
    - Affects patterns like Visitor

- No support for exceptions - error handling is via return values
    - Affects  patterns like Builder

- No support for abstract classes
    - Affects patterns like Abstract Factory and Bridge

But still the syntax of golang is much simpler than old languages like C++, that makes many design patterns much obvious.

## Design Pattern Categories 🐛

| Creational 🏗️ | Structural 🏢 | Behavioral 🧑‍🦰 |
| --- | ---  | --- |
| **Abstract Factory** | **Adapter** | **Chain of Responsibility** |
| **Builder** | Bridge | Command |
| **Factory** | Composite | Interpreter |
| Lazy Initialization | **Decorator** | **Iterator** |
| Multiton | **Facade** | Mediator |
| Object Pool | Flyweight | Momento |
| Protype | **Proxy** | **Observer** |
| **Singleton** |  | **Strategy** |
|  |  | Visitor |

# Is-A and Has-A Relationships in Golang 🔗

While Golang doesn't have traditional class-based inheritance, it still supports concepts similar to Is-A and Has-A relationships through its unique features. 🎭

## Is-A Relationship 🧬

Golang doesn't have inheritance, but it uses interfaces to achieve polymorphism, which can be seen as a form of Is-A relationship. 🔄

```go
type Vehicle interface {
    Move()
}

type Car struct{}

func (c Car) Move() {
    fmt.Println("Car is moving")
}

// Car "is a" Vehicle because it implements the Vehicle interface 🚗
```

## Has-A Relationship 🧩

Has-A relationships are implemented in Golang through composition, which is a key feature of the language. 🔧

Composition in Go refers to the practice of building complex types by combining simpler ones. It's a way of structuring your code where you embed one type within another, allowing the outer type to use the methods and fields of the embedded type. 🏗️

```go
type Engine struct{}

func (e Engine) Start() {
    fmt.Println("Engine started")
}

type Car struct {
    engine Engine  // Car "has an" Engine 🚙
}

func (c Car) StartCar() {
    c.engine.Start()
    fmt.Println("Car is ready to move")
}
```

In Golang, composition is often preferred over inheritance, following the principle "composition over inheritance". This approach allows for more flexible and modular code design. 🧠💡

## Troubleshooting

### Not able to do ping 🤮

ping: www.google.com: Temporary failure in name resolution

This usually happens due to misconfigured DNS 🛜. To resolve this:
1. Open the /etc/resolv.conf file:

    `
    sudo nano /etc/resolv.conf
    `

2. Add working DNS servers, such as Google's public DNS or Cloudflare's DNS, to ensure proper resolution:
    
    `
    nameserver 8.8.8.8
    nameserver 1.1.1.1
    `
3. Save the changes (Ctrl + O, then Enter) and exit (Ctrl + X).

If you are using WSL, then by default WSL regenerates the `resolv.conf` file on every launch, you'll need to disable this feature: ♿

1. Open or create the WSL configuration file:

    `sudo nano /etc/wsl.conf`

2. Add the following lines (or confirm they're already present):

    `
    [network]
    generateResolvConf = false
    `
3. Save and exit the file.

# References

[Course on LinkedIn Learning](https://www.linkedin.com/learning/go-design-patterns/) by Joe Marini

[Refactoring Guru](https://refactoring.guru/)

## How to Run 

Type `go run *.go` inside any folder. It will run all the files and you should not encounter any undefined struct error.