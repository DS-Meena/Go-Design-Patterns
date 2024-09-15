# Singleton Pattern 🔒

The Singleton pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance. It's useful when exactly one object is needed to coordinate actions across the system. 🏗️

## Purpose
- The Singleton pattern restricts the instantiation of a class to a single instance and provides global access.
- Allows for lazy initialization of the class

## Scenarios
- Situations where you want to ensure there is only one instance of a class: logging, configuration, telemetry, debugging.

## How to do this

1. By checking if global instance variable is nil.
2. For goroutines, you can use the sync.Once method to achieve Singleton.
3. Double-checked Locking [Java]

![Singleton Flowchart](../image-2.png)

**To note**

While using double-checked locking, you might encounter some issues like:
1. Instruction Reordering
2. L1 Cache Issue

These can cause, creation of two objects. Using `volatile` keyword can solve this.

The Go example above using sync.Once is a safe and efficient way to implement the Singleton pattern without the issues of double-checked locking. 👍
