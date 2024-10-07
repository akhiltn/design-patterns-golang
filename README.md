# design-patterns-golang
GOF design pattern in GO


## Creational Patterns:
- Singleton: Ensures a class has only one instance.
```mermaid
classDiagram
    class Singleton {
        - int count
        - Singleton()
        + GetInstance()$ Singleton
    }
```
- Factory: Hides the instantiation logic.
```mermaid
classDiagram
    class Animal {
        <<interface>>
        + Speak() string
    }

    class Dog {
        + Speak() string
    }

    class Cat {
        + Speak() string
    }

    Animal <|-- Dog
    Animal <|-- Cat

    class AnimalFactory {
        + GetAnimal(animal string) Animal
    }

    AnimalFactory --> Animal : "creates"
```
- Builder: Constructs complex objects step by step.
```mermaid
classDiagram
    class House {
        - string Foundation
        - string Roof
        - string Wall
    }

    class HouseBuilder {
        - House house
        + buildFoundation() HouseBuilder
        + buildRoof() HouseBuilder
        + buildWall() HouseBuilder
        + build() House
    }

    class Director {
        - HouseBuilder builder
        + setBuilder(b HouseBuilder) void
        + getHouse() House
    }

    Director --> HouseBuilder : "uses"
    HouseBuilder --> House : "builds"
```

## Structural Patterns:
- Adapter: Bridges incompatibility between interfaces.
- Decorator: Adds responsibilities to objects dynamically.
- Composite: Treats individual objects and compositions uniformly.

## Behavioral Patterns:
- Strategy: Encapsulates algorithms within a family.
- Observer: Notifies dependents of state changes.
- Command: Encapsulates method invocation.

## Concurrency Patterns (specific to Go):
- Pipeline: Passes data through stages of transformations.
- Worker Pools: Distributes tasks among workers.
- Channel Oriented Patterns: Communication between goroutines.
- Context Pattern: Useful for controlling goroutine lifecycles, canceling operations, and managing deadlines.

