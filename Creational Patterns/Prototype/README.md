# Prototype Pattern 🧬

The Prototype pattern is a creational design pattern that allows you to copy existing objects without making your code dependent on their classes. It's particularly useful when the cost of creating a new object is more expensive than copying an existing one. 🔄💰

## How it works 🛠️

1. Define a prototype interface with a clone() method.
2. Implement the clone() method in concrete classes.
3. Create new objects by copying existing instances.

## Scenarios to Use 🎯

- When you need to create objects based on existing objects 🏗️
- When object creation is expensive or complex 💸
- When you want to reduce the number of subclasses 📉
- When you need to create objects at runtime 🏃‍♂️

The Prototype pattern is a powerful tool in your design pattern toolkit, offering flexibility and efficiency in object creation. Use it wisely to streamline your code and improve performance! 🚀

## Example 👨‍💻

In the example, we create a conrete class and add a clone method in it.

```
go run *.go
Original: &{David}
Clone: &{David}
```