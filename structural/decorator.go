package structural

type Coffee interface {
  GetCost() float32
  GetDescription() string
}

type SimpleCoffee struct {

}

func (s *SimpleCoffee) GetCost() float32 {
  return 10.0
}

func (s *SimpleCoffee) GetDescription() string {
  return "Simple Coffee"
}

// type CoffeeDecorator struct {
//   coffee Coffee
// }

// func (c *CoffeeDecorator) GetCost() float32 {
//   return c.coffee.GetCost()
// }

// func (c *CoffeeDecorator) GetDescription() string {
//   return c.coffee.GetDescription();
// }

type MilkCoffee struct {
  coffee Coffee
}

func GetMilkCoffee(coffee Coffee) Coffee {
  return &MilkCoffee{coffee: coffee}
}

func (c *MilkCoffee) GetCost() float32 {
  return c.coffee.GetCost() + 5.0
}

func (c *MilkCoffee) GetDescription() string {
  return c.coffee.GetDescription() + " + milk";
}


type SugarCoffee struct {
  coffee Coffee
}

func GetSugarCoffee(coffee Coffee) Coffee {
  return &SugarCoffee{coffee: coffee}
}

func (c *SugarCoffee) GetCost() float32 {
  return c.coffee.GetCost() + 2.0
}

func (c *SugarCoffee) GetDescription() string {
  return c.coffee.GetDescription() + " + sugar";
}

