## Composite Design Pattern 🌳

The Composite Design Pattern is a structural pattern that allows you to compose objects into tree structures to represent part-whole hierarchies. It lets clients treat individual objects and compositions of objects uniformly. 🧩

Part-whole hierarchy - objects are composed of smaller parts, which can themselves be composed of even smaller parts.

## When to Use 🤔

- When you want to represent part-whole hierarchies of objects
- When you want clients to be able to ignore the difference between compositions of objects and individual objects

## Example: File System 📁

Let's consider a file system as an example. A file system contains directories and files. A directory can contain files and other directories. This forms a tree-like structure, perfect for the Composite pattern.

In this example:

- FileSystemComponent is the component interface 🔧
- File is the leaf 🍃
- Directory is the composite 📂

```
Total size of root: 450
Total size of subDir: 300
```

This structure allows you to treat both files and directories uniformly. You can easily calculate the size of a directory by summing up the sizes of its components, which could be files or other directories. 🧮

The Composite pattern makes it easy to add new types of components to the system without changing the existing code, adhering to the Open-Closed Principle. 🚀