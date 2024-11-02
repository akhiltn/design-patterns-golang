# design-patterns-golang
GOF design pattern in GO


## Creational Patterns:
- Singleton: Ensures a class has only one instance.
```mermaid
---
title: SingletonPattern
---
classDiagram

    class Singleton {
        - int count
        - Singleton()
        + GetInstance()$ Singleton
    }
```
- Factory: Hides the instantiation logic.
```mermaid
---
title: FactoryPattern
---
classDiagram

    AnimalFactory *-- Animal : "creates"
    Animal <|.. Dog
    Animal <|.. Cat


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
```
- Builder: Constructs complex objects step by step.
```mermaid
---
title: BuilderPattern
---
classDiagram

    Director --> HouseBuilder : "uses"
    HouseBuilder --> House : "builds"

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
```

## Structural Patterns:
- Adapter: Bridges incompatibility between interfaces.
```mermaid
---
title: AdapterPattern
---
classDiagram

    IToyBird <|-- BirdAdapter
    BirdAdapter *-- Bird

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
```
- Decorator: Adds responsibilities to objects dynamically.
```mermaid
classDiagram
    Coffee <|.. SimpleCoffee
    Coffee <|.. CoffeeDecorator
    CoffeeDecorator <|-- MilkCoffee : "decorates"
    CoffeeDecorator <|-- SugarCoffee : "decorates"

    class Coffee {
        <<interface>>
        + GetCost() float32
        + GetDescription() string
    }

    class SimpleCoffee {
        + GetCost() float32
        + GetDescription() string
    }

    class CoffeeDecorator {
        - coffee Coffee
        + GetCost() float32
        + GetDescription() string
    }

    class MilkCoffee {
        + GetCost() float32
        + GetDescription() string
    }

    class SugarCoffee {
        + GetCost() float32
        + GetDescription() string
    }
```
- Composite: Treats individual objects and compositions uniformly.
```mermaid
---
title: CompositePattern
---
classDiagram

    Component <|-- Composite
    Component <|-- Leaf
    Composite *-- "1..*" Component

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
```

## Behavioral Patterns:
- Strategy: Encapsulates algorithms within a family.
```mermaid
---
title: StrategyPattern
---
classDiagram

    PaymentGateway <|.. CreditcardPayment
    PaymentGateway <|.. BitcoinPayment
    PaymentGateway <|.. PaypalPayment
    ShoppingCart <-- PaymentGateway
    
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
```
- Observer: Notifies dependents of state changes.
```mermaid
---
title: ObserverPattern
---
classDiagram
    class Subject {
        +Attach(Observer)
        +Detach(Observer)
        +Notify()
    }

    class ConcreteSubject {
        -observers: Observer[]
        +SetState(state: string)
        +state: string
        +Attach(Observer)
        +Detach(Observer)
        +Notify()
    }

    class Observer {
        +Update(state: string)
    }

    class ConcreteObserver {
        +name: string
        +Update(state: string)
    }

    Subject <|.. ConcreteSubject
    Observer <|.. ConcreteObserver
    ConcreteSubject --> Observer : Notifies
```
- Command: Encapsulates method invocation.
```mermaid
---
title: Command Pattern Diagram
---
classDiagram
    Command <|.. LightOnCommand
    Command <|.. LightOffCommand
    LightOnCommand *-- Light
    LightOffCommand *-- Light
    RemoteControl *-- Command

    class Light {
        +state: bool
    }
    
    class LightOnCommand {
        +light: Light
        +Execute()
    }
    
    class LightOffCommand {
        +light: Light
        +Execute()
    }

    class Command {
        <<interface>>
        +Execute()
    }

    class RemoteControl {
        +command: Command
        +setCommand(command: Command)
        +Press()
    }
```

## Concurrency Patterns (specific to Go):
- Pipeline: Passes data through stages of transformations.
- Worker Pools: Distributes tasks among workers.
- Channel Oriented Patterns: Communication between goroutines.
- Context Pattern: Useful for controlling goroutine lifecycles, canceling operations, and managing deadlines.

