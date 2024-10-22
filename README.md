# design-patterns-golang
GOF design pattern in GO


## Creational Patterns:
- Singleton: Ensures a class has only one instance.
```mermaid
---
title: SingletonPattern
---
classDiagram

namespace SingletonPattern{
    class Singleton {
        - int count
        - Singleton()
        + GetInstance()$ Singleton
    }
}
```
- Factory: Hides the instantiation logic.
```mermaid
classDiagram

    AnimalFactory *-- Animal : "creates"
    Animal <|-- Dog
    Animal <|-- Cat


namespace FactoryPattern{
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

    class AnimalFactory {
        + GetAnimal(animal string) Animal
    }
}
```
- Builder: Constructs complex objects step by step.
```mermaid
classDiagram

    Director --> HouseBuilder : "uses"
    HouseBuilder --> House : "builds"

namespace BuilderPattern{
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
}
```

## Structural Patterns:
- Adapter: Bridges incompatibility between interfaces.
```mermaid
classDiagram

    IToyBird <|-- BirdAdapter
    BirdAdapter *-- Bird

namespace AdapterPattern{
    class IToyBird {
        <<interface>>
        + MakeSound()
    }

    class Bird {
        + Fly()    
        + Chirp()
    }

    class BirdAdapter {
        - Bird bird
        + BirdAdapter(Bird)
        + MakeSound()
    }
}
```
- Decorator: Adds responsibilities to objects dynamically.
```mermaid
classDiagram
    Coffee <|.. SimpleCoffee
    Coffee <|.. CoffeDecorator
    CoffeDecorator <|-- SugarCoffeeDecorator : "decorates"
    CoffeDecorator <|-- MilkCoffeeDecorator : "decorates"

namespace DecoratorPattern{
    class Coffee{
        <<interface>>
        + GetCost() int
        + GetDescription() string
    }
    
    class SimpleCoffee

    class CoffeDecorator{
        - coffee Coffee
    }

    class SugarCoffeeDecorator{
        + GetCost() int
        + GetDescription() string
    }

    class MilkCoffeeDecorator{
        + GetCost() int
        + GetDescription() string
    }
```
- Composite: Treats individual objects and compositions uniformly.
```mermaid
classDiagram

    Component <|-- Composite
    Component <|-- Leaf
    Composite *-- "1..*" Component

    namespace CompositePattern {
        class Component {
            <<interface>>
            + Operation()
        }

        class Composite {
            - components: []Component
            + Add(component: Component)
            + Remove(component: Component)
            + Operation()
        }

        class Leaf {
            + Operation()
        }
    }
```

## Behavioral Patterns:
- Strategy: Encapsulates algorithms within a family.
```mermaid
classDiagram
    PaymentGateway <|.. CreditcardPayment
    PaymentGateway <|.. BitcoinPayment
    PaymentGateway <|.. PaypalPayment
    ShoppingCart <-- PaymentGateway
    
    namespace StrategyPattern {
        class PaymentGateway {
            <<interface>>
            Pay(amount float) string
        }

        class CreditcardPayment {
            Pay(amount float) string
        }

        class BitcoinPayment {
            Pay(amount float) string
        }
        
        class PaypalPayment {
            Pay(amount float) string
        }
        
        class ShoppingCart {
            +payment PaymentGateway
        }
    }
```
- Observer: Notifies dependents of state changes.
- Command: Encapsulates method invocation.

## Concurrency Patterns (specific to Go):
- Pipeline: Passes data through stages of transformations.
- Worker Pools: Distributes tasks among workers.
- Channel Oriented Patterns: Communication between goroutines.
- Context Pattern: Useful for controlling goroutine lifecycles, canceling operations, and managing deadlines.

