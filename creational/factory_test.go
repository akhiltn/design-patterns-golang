package creational

import "testing"

// Define your test cases
var testcase = map[string]string{
    "dog":   "Woof!",
    "cat":   "Meow!",
    "human": "Hello",
}

// TestGetAnimal tests the GetAnimal function
func TestFactory(t *testing.T) {
    for k, v := range testcase {
        animal := GetAnimal(k)
        if animal == nil {
            if k == "human" {
                continue // Expecting nil for "human"
            }
            t.Errorf("Expected an animal, got nil for %s", k)
            continue
        }
        if animal.Speak() != v {
            t.Errorf("Expected %s, got %s", v, animal.Speak())
        }
    }
}
