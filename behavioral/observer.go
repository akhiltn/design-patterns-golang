package behavioral

import "fmt"

// Observer interface with the Update method.
type Observer interface {
	Update(string)
}

// Subject interface with methods for managing observers.
type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

// ConcreteSubject holds observers and its state.
type ConcreteSubject struct {
	observers []Observer
	state     string
}

// Attach adds an observer to the list.
func (s *ConcreteSubject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

// Detach removes an observer from the list.
func (s *ConcreteSubject) Detach(o Observer) {
	for i, obs := range s.observers {
		if obs == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify updates all observers with the current state.
func (s *ConcreteSubject) Notify() {
	for _, obs := range s.observers {
		obs.Update(s.state)
	}
}

// SetState changes the state and notifies observers.
func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.Notify()
}

// ConcreteObserver reacts to updates from ConcreteSubject.
type ConcreteObserver struct {
	name      string
	lastState string // Stores the last received state for testing
}

// Update receives the new state from the subject.
func (o *ConcreteObserver) Update(state string) {
	o.lastState = state
	fmt.Printf("%s received update: %s\n", o.name, state)
}

// NewConcreteObserver is a constructor for ConcreteObserver.
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}
