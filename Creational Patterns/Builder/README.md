
# Builder
Builder simplifies the creation of complex objects that have many possible representations.

## Purpose 
- *Encapsulates an object's construction process* along with specifying the various parts of a comlex API
- Enables flexible creation of an object that can have many different representations
- Increase code readability for complex types

## Scenarios:

Objects that have complex APIs, multiple constructor options, and several different possible representations.

## How to do this:
- Create a builder object, with same fields and many setter methods.
- Use Build method, to return a finished object from builder object.

![Builder Flowhcart](../image.png)