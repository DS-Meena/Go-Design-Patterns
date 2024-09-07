# Observer
Observer defines a pattern where a given object maintains a list of dependent objects that are notified when the state of the main object changes.

## Purpose:
- Provides the ability for a subject to notify a set of "observers" about changes to the subject
- Used to maintain loose coupling between elements of a system that interact with each other

## Scenarios:
- Systems where state changes of a subject cannot be predicted and the list of interested "listeners" cannot be known in advance or changes during runtime

![Observer Pattern](../image.png)