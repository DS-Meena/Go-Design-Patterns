## Creational Pattern
Creational design patterns deal with object creation mechanisms, trying to create objects in a manner suitable to the situation.

The most common types of creational design patterns include Factory Method, Abstract Factory, Builder, Prototype, and Singleton.

## Singleton

Purpose:
- The Singleton pattern restricts the instantiation of a class to a single instance and provides global access.
- Allows for lazy initialization of the class

Scenarios:
- Situations where you want to ensure there is only one instance of a class: logging, configuration, telemetry, debugging.

How to do this:

- By checking if global instance variable is nil
- For goroutines, you can use the sync.Once method to achieve Singleton.

![Singleton Flowchart](image-2.png)