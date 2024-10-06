package creational

// Animal is an interface with a Speak method
type Animal interface {
	Speak() string
}

// Dog implements Animal
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat implements Animal
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

func GetAnimal(animal string) Animal {
	switch animal {
	case "dog":
		return new(Dog)
	case "cat":
		return new(Cat)
	default:
		return nil
	}
}
