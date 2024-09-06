# Strategy Pattern 🧠

The Strategy Pattern is a behavioral design pattern that enables selecting an algorithm at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable. 🔄

# Key components
Key components of the Strategy Pattern include:

- Context 🎭: Maintains a reference to a Strategy object and may define an interface for clients to interact with the strategy.
- Strategy 📊: Declares an interface common to all supported algorithms. The Context uses this interface to call the algorithm defined by a ConcreteStrategy.
- ConcreteStrategy 🧩: Implements the algorithm using the Strategy interface.

# Example:
In our example: 🛍️

- We define a `PaymentStrategy` interface with a `Pay` method. 💰
- We implement two concrete strategies: `CreditCardPayment` and `BhimPayment`. 💳🌐
- The `ShoppingCart` acts as the context, which can use different payment strategies. 🛒
- In the `main` function, we demonstrate how to switch between strategies at runtime. 🔄

This pattern allows us to easily add new payment methods without changing the `ShoppingCart` code, adhering to the Open/Closed Principle. 🎯

Output:
```
Paid 100.89 using Credit Card 💳
Paid 32.74 using bhim 🌐
```

# How to use 

Create a setter method which takes interface as argument. This way you can decide which conrete strategy to use at runtime.


# Advantages

Benefits of using the Strategy Pattern:

- Flexibility 🤸‍♂️: Algorithms can be switched at runtime.
- Isolation 🏝️: Implementation details of an algorithm are isolated from the code that uses it.
- Reusability ♻️: Strategies can be reused across different contexts.
- Testability 🧪: Individual strategies can be easily unit tested.

This pattern is particularly useful when you have a set of similar algorithms and need to switch between them dynamically, based on the client's needs or runtime conditions. 🔀