# Proxy Design Pattern 🎭

The Proxy design pattern provides a surrogate or placeholder for another object to control access to it. It's like having a representative that manages interactions with the real object. 🕴️

## Purpose 🎯

The main purposes of the Proxy pattern are:

- Control access to the real object 🔒
- Perform actions before or after requests reach the real object 🔄
- Provide a level of indirection to support distributed, controlled, or intelligent access 🌐

## Scenarios to Use 🛠️

The Proxy pattern is useful in several scenarios:

- Virtual Proxy: Delaying the creation and initialization of expensive objects 💰
- Remote Proxy: Representing an object that's in a different address space 🌍
- Protection Proxy: Controlling access rights to an object 🛡️
- Smart Reference: Performing additional actions when an object is accessed 🧠

Example: 
1. Access restriction to inappropriate websites in an organization.
2. Cachinng
3. Preprocessing & Post processing.

# Caching Example:
Let's consider a virtual proxy example for a high-resolution image viewer:

In this example, the ProxyImage acts as a surrogate for the RealImage. It controls access to the RealImage, only creating and loading it when necessary. This can significantly improve performance when dealing with resource-intensive objects. 🚀

```
Loading photo1.jpg from disk
Displaying photo1.jpg
Displaying photo1.jpg
Loading photo2.jpg from disk
Displaying photo2.jpg
```

The Proxy pattern allows for a more efficient use of system resources and provides a way to add additional behaviors without changing the RealImage class, adhering to the Open/Closed principle of SOLID. 🧱