# Design patterns

## Creational Patterns

## Structural Patterns

## Behavioral Patterns

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

Example: [AoC 2022, day 11](https://adventofcode.com/2022/day/11) (the operation done by each monkey)