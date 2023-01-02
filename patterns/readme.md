# Design patterns

<details>
    <summary>Table of Content</summary>

- [Creational Patterns](#creational-patterns)
  - [Builder](#builder)
  - [Factory Method](#factory-method)
  - [Abstract Factory](#abstract-factory)
- [Structural Patterns](#structural-patterns)
- [Behavioral Patterns](#behavioral-patterns)
  - [State](#state)
  - [Strategy](#strategy)
- [Links](#links)

</details>

## Creational Patterns

### Builder

Construct complex objects step by step.

![img.png](../img/patterns/creational/builder.png)

- create different representations of the object by using the same construction code
- the Builder doesn’t allow other objects to access the product while it’s being built


- builder
  - extract the construction code
  - can execute a series of steps on a builder object
  - call only those steps that are necessary for producing a particular configuration of an object
  - can have several builder classes that implement the same steps but in a different way
  - builders need to provide their own methods for returning results as different products might be created from different builders
    (the products don't follow a common interface)


- director
  - not mandatory
  - extract a series of calls to the builder steps 
  - it defines the order in which to execute the building steps
  - provides reusability - can be used for grouping different construction routines
  - the director class completely hides the details of product construction from the client code
  - the client only needs to associate a builder with a director, launch the construction with the director, and get the result from the builder.


- used when
  - the products are quite complex and require extensive configuration
  - want to create different representations of the products (same base product; for eg., stone and wooden houses)
  - the representations are based on similar steps with different details only
  - isolate complex product construction logic from business logic ([Single Responsibility Principle](../principles/solid/readme.md))

[Code here](./creationalpatterns/examples/builder.go).

### Factory Method

Hide the creation of the object and expose it under an interface.

![img.png](../img/patterns/creational/factorymethod.png)

- the objects (aka products) are created inside a factory function
- the creator, the one which defines the factory method, has other primary responsibility instead of creating products
- it has some business logic related to the created products
- the client uses an instance of a concrete creator but doesn't know which concrete product it uses


- used when the exact types and dependencies are not known beforehand
- the product creation is separated from its use
- can be reusing existing products instead of creating new ones each time


- used when
  - want to avoid tight coupling between the products and the creator
  - the product creation is in one place only ([Single Responsibility Principle](../principles/solid/readme.md))
  - makes it easy to add new product types ([Open / Closed Principle](../principles/solid/readme.md))

[Code here](./creationalpatterns/examples/factorymethod.go).

### Abstract Factory

Create families of objects

![img.png](../img/patterns/creational/abstractfactory.png)

- factory of factories
- add new products or families of products easily
- the products from factories are compatible with each other
- loose coupling between concrete products and client code
- the product creation is in one place only ([Single Responsibility Principle](../principles/solid/readme.md))
- makes it easy to add new products or variants  ([Open / Closed Principle](../principles/solid/readme.md))

[Code here](./creationalpatterns/examples/abstractfactory.go).

## Structural Patterns

## Behavioral Patterns

### State

An object alters its behavior when its internal state changes.

![img.png](../img/patterns/behavioralpatterns/state.png)

- there are a finite number of states which a program can be in
- for each one the program behaves differently
- can change from one to another easily by some predefined and finite rules (transitions)
- any change to the transition logic may require changing state conditions in every method


- create new types for each possible state and extract all state-specific behavior in them
- transition by replacing the current active state with the new state
- create new states or change existing one independently of each other


- used when
  - needing to organize code and decrease complexity ([Single Responsibility Principle](../principles/solid/readme.md))
  - there are a lot of states and want to add new states independently of existing ones ([Open / Closed Principle](../principles/solid/readme.md))

[Code here](./behavioralpatterns/examples/state.go).

### Strategy

For designing a family of interchangeable algorithms hidden behind an abstraction.

![img.png](../img/patterns/behavioralpatterns/strategy.png)

- strategy
  - interface defining the work
  - one method that triggers the algorithm for the strategy implementation


- context 
  - contains a reference to the chosen strategy implementation
  - it delegates the work to a linked strategy
  - the client passes to the context the chosen strategy
  - the context is independent of the concrete strategy


- used when
  - using variations of an algorithm
  - you can alter the context's behavior at runtime and associate a different strategy to it
  - respects [Open / Closed Principle](../principles/solid/readme.md) - introduce different strategies or policies without changing the context

[Code here](./behavioralpatterns/examples/strategy.go).

Example: [AoC 2022, day 11](https://adventofcode.com/2022/day/11) (the operation done by each monkey)

## Links

- https://refactoring.guru/design-patterns
- https://golangbyexample.com/all-design-patterns-golang/
- https://medium.com/swlh/design-patterns-in-go-d90e7866deff
- https://dev.to/mauriciolinhares/gof-design-patterns-that-still-make-sense-in-go-27k5
- https://github.com/tmrts/go-patterns