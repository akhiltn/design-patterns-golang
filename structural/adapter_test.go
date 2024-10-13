package structural

import (
	"testing"
)

func TestAdapter(t *testing.T) {

	// Create the bird and the adapter
	var toy IToyBird
	bird := new(Bird)
	toy = NewBirdAdapter(bird)

	// Get the captured output and check if it's correct
	expected := bird.Chirp()
	actual := toy.MakeSound()

	if actual != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, actual)
	}
}
