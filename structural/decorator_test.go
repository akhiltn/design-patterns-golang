package structural

import (
  "testing"
)

func TestDecorator(t *testing.T) {
  var coffee Coffee
  coffee = &SimpleCoffee{}
  coffee = GetMilkCoffee(coffee)
  if coffee.GetDescription() != "Simple Coffee + milk" {
    t.Errorf("Expected 'Simple Coffee + milk' but got '%s'", coffee.GetDescription())
  }
  if coffee.GetCost() != 15.0 {
    t.Errorf("Expected 15.0 but got '%f'", coffee.GetCost())
  }
  coffee = GetSugarCoffee(coffee)
  if coffee.GetDescription() != "Simple Coffee + milk + sugar" {
    t.Errorf("Expected 'Simple Coffee + milk + sugar' but got '%s'", coffee.GetDescription())
  }
  if coffee.GetCost() != 17.0 {
    t.Errorf("Expected 17.0 but got '%f'", coffee.GetCost())
  }
}
