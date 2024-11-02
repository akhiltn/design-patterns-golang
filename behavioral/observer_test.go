
package behavioral

import (
	"testing"
)

// Test if attaching observers works as expected.
func TestAttachObserver(t *testing.T) {
	subject := &ConcreteSubject{}
	observer := NewConcreteObserver("Observer 1")

	subject.Attach(observer)

	if len(subject.observers) != 1 {
		t.Errorf("Expected 1 observer, got %d", len(subject.observers))
	}
}

// Test if detaching observers works as expected.
func TestDetachObserver(t *testing.T) {
	subject := &ConcreteSubject{}
	observer := NewConcreteObserver("Observer 1")

	subject.Attach(observer)
	subject.Detach(observer)

	if len(subject.observers) != 0 {
		t.Errorf("Expected 0 observers, got %d", len(subject.observers))
	}
}

// Test if observers receive updates when subject's state changes.
func TestObserverNotification(t *testing.T) {
	subject := &ConcreteSubject{}
	observer1 := NewConcreteObserver("Observer 1")
	observer2 := NewConcreteObserver("Observer 2")

	subject.Attach(observer1)
	subject.Attach(observer2)

	// Change the state and check if observers receive it.
	newState := "New State"
	subject.SetState(newState)

	if observer1.lastState != newState {
		t.Errorf("Observer 1 expected state %s, got %s", newState, observer1.lastState)
	}

	if observer2.lastState != newState {
		t.Errorf("Observer 2 expected state %s, got %s", newState, observer2.lastState)
	}
}

// Test if observers no longer receive updates after being detached.
func TestDetachedObserverDoesNotReceiveUpdate(t *testing.T) {
	subject := &ConcreteSubject{}
	observer1 := NewConcreteObserver("Observer 1")
	observer2 := NewConcreteObserver("Observer 2")

	subject.Attach(observer1)
	subject.Attach(observer2)

	subject.Detach(observer1)

	// Change the state and check if observer1 does not receive it.
	newState := "Another State"
	subject.SetState(newState)

	if observer1.lastState == newState {
		t.Errorf("Observer 1 should not have received state %s after detachment", newState)
	}

	if observer2.lastState != newState {
		t.Errorf("Observer 2 expected state %s, got %s", newState, observer2.lastState)
	}
}
