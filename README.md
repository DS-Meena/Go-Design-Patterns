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