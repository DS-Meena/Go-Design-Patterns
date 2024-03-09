## Structural Pattern

## Adapater

Purpose: 
- Allows the interface of an existing subsystem or API to be used as another interface without modifying the code of the existing API

Scenarios:
- Enables incompatible objects to work together without having to make changes to either one.

How to do: 
- create a new struct that uses old struct and its receiver methods.

![Adapter Pattern](image.png)

## Facade

Purpose:
- Provide a simple, front-facing interface to a more complex system, library, or API

Scenarios:
- Improve usability of a more complex API
- Serve as a starting point for refactoring
- Reduce tight coupling between parts of a system

How to use:
- Instead of directly interacting with the API, we use inbetween facade methods.

![Facade Pattern](image-1.png)