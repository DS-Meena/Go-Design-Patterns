# Decorator Pattern 🎨

The Decorator pattern is a structural design pattern that allows behavior to be added to an individual object, either statically or dynamically, without affecting the behavior of other objects from the same class. 🔧

This does by placing the objects inside special wrapper objects called Decorators. Decorator has both is-a and has-a relationship with the interface or base class.

## Purpose 🎯

The main purpose of the Decorator pattern is to:

- Add responsibilities to objects dynamically and transparently 🔄
- Provide a flexible alternative to subclassing for extending functionality 🌿
- Allow for the combination of behaviors in ways that are not possible with static inheritance 🧩

## Scenarios to Use 🛠️

The Decorator pattern is useful in the following scenarios:

- When you need to add responsibilities to objects dynamically without affecting other objects 🚀
- When extension by subclassing is impractical or impossible 🚫
- When you want to add or remove responsibilities at runtime 🕒

## Example

In this example, we have a Coffee interface that defines the basic operations. The SimpleCoffee struct implements this interface. We then have decorator structs (Milk and Whip) that also implement the Coffee interface and contain a Coffee object. These decorators add their own behavior to the contained Coffee object. 🍵☕

```
Cost :10; Description: Simple coffee
Cost :15; Description: Simple coffee, milk
Cost :22; Description: Simple coffee, milk, whip
```

This pattern allows us to add new behaviors (like adding milk or whip) to our coffee objects dynamically, without changing their code or using inheritance. It provides a flexible way to extend functionality at runtime. 🚀

## Why Intermediary decorator

The `CoffeeDecorator` struct serves as an intermediary layer in the Decorator pattern implementation. 🏗️ While it's true that you can directly add the `Coffee` interface to each concrete decorator, using the `CoffeeDecorator` struct provides several benefits:

1. Code Reusability 🔄: The `CoffeeDecorator` implements the basic forwarding behavior for both `GetCost()` and `GetDescription()` methods. This avoids duplicating this code in each concrete decorator.
2. Simplified Concrete Decorators 🧩: Concrete decorators like `Milk` and `Whip` can focus solely on their specific additions, making them cleaner and easier to maintain.
3. Extensibility 🚀: If we need to add common functionality to all decorators in the future, we can do so in the `CoffeeDecorator` struct without modifying each concrete decorator.
4. Adherence to the Open/Closed Principle 🚪: The `CoffeeDecorator` allows for easy addition of new decorators without modifying existing code.

While it's possible to implement the Decorator pattern without this intermediate struct, using it is generally considered a best practice in object-oriented design 🏆, especially for more complex systems where you might have multiple levels of decoration. 🎭