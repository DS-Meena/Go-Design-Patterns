# Bridge Design Pattern 🌉

The Bridge design pattern is a structural pattern that decouples an abstraction from its implementation so that the two can vary independently. It's like building a bridge between two independent dimensions. 🏗️

## Key Components 🔑

- **Abstraction:** Defines the abstract interface and maintains a reference to the implementor.
- **Refined Abstraction:** Extends the abstraction and provides more refined operations.
- **Implementor:** Defines the interface for implementation classes.
- **Concrete Implementor:** Implements the Implementor interface.

## When to Use 🤔

- When you want to avoid a permanent binding between an abstraction and its implementation. 🔓
- When both the abstractions and their implementations should be extensible by subclassing. 🌳
- When changes in the implementation of an abstraction should have no impact on clients. 🛡️
- When you have a proliferation of classes resulting from a coupled interface and numerous implementations. 🧩

The Bridge pattern is particularly useful in scenarios where you need to handle multiple variations of a concept, like different types of UI elements across various platforms, or different database drivers for an application. 🚀

## Example

Let's understand the bridge pattern with an example.

1. Abstraction: `RemoteControl` and refined abstraction `AdvancedRemoteControl`
2. Implementor: `Device`
3. Concrete Implementors: `TV` and `Radio` structs.

The Bridge is the connection between the Abstraction (`RemoteControl`) and the Implementor (`Device`). In this case, it's represented by the `device` field in the `RemoteControl` struct:

```
type RemoteControl struct {
	device Device
}
```

This allows the abstraction to work with any implementor of the `Device` interface, creating a "bridge" between the two hierarchies. The abstraction can use the implementor's methods without being tightly coupled to a specific concrete implementation.

This design allows for independent variation of both the remote controls (abstractions) and the devices (implementations), which is the key benefit of the Bridge pattern.

```
go run *.go
Testing TV remote: 
TV turned on
TV turned off

Testing Radio remote:
Radio turned on
Radio turned off
Radio frequency set to 100
```