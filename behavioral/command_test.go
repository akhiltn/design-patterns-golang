
package behavioral

import (
	"testing"
)

// Test if LightOnCommand turns the light on.
func TestLightOnCommand(t *testing.T) {
	light := &Light{ state: false }
	lightOnCommand := &LightOnCommand{light: light}
	remote := &RemoteControl{}

	// Set the command and press the button to execute.
	remote.setCommand(lightOnCommand)
	remote.Press()

	if !light.state {
		t.Errorf("Expected light state to be true, got %v", light.state)
	}
}

// Test if LightOffCommand turns the light off.
func TestLightOffCommand(t *testing.T) {
	light := &Light{state: true} // Initially, the light is on.
	lightOffCommand := &LightOffCommand{light: light}
	remote := &RemoteControl{}

	// Set the command and press the button to execute.
	remote.setCommand(lightOffCommand)
	remote.Press()

	if light.state {
		t.Errorf("Expected light state to be false, got %v", light.state)
	}
}
