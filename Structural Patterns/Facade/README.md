# Facade
The Facade pattern is mainly used to simplify the interface to a more complex object **or** to provide a context-specific interface to a more generic API.

## Purpose:
- Provide a simple, front-facing interface to a more complex system, library, or API

## Scenarios:
- Improve usability of a more complex API
- Serve as a starting point for refactoring
- Reduce tight coupling between parts of a system

## How to use:
- Instead of directly interacting with the API, we use inbetween facade methods (Facade layer).

![Facade Pattern](../image-1.png)

## Differences

### Facade vs Proxy 🏠🔐

- **Facade:** Simplifies a complex system by providing a unified interface 🎭
- **Proxy:** Controls access to an object, acting as a surrogate or placeholder 🚪

Key difference: Facade simplifies, while Proxy controls access.

### Facade vs Adapter 🏠🔌

- **Facade:** Provides a simplified interface to a complex subsystem 🎭
- **Adapter:** Allows incompatible interfaces to work together 🔄

Key difference: Facade simplifies existing interfaces, while Adapter makes incompatible interfaces compatible.

Remember, Facade is about simplification 🧘, Proxy is about access control 🔒, and Adapter is about compatibility 🤝.