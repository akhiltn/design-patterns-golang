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

type CoffeeDecorator struct {
  coffee Coffee
}

func (c *CoffeeDecorator) GetCost() float32 {
	return c.coffee.GetCost()
}

func (c *CoffeeDecorator) GetDescription() string {
	return c.coffee.GetDescription()
}

type MilkCoffee struct {
  CoffeeDecorator
}

func (c *MilkCoffee) GetCost() float32 {
	return c.CoffeeDecorator.GetCost() + 5.0
}

func (c *MilkCoffee) GetDescription() string {
	return c.CoffeeDecorator.GetDescription() + " + milk"
}

type SugarCoffee struct {
  CoffeeDecorator
}

func (c *SugarCoffee) GetCost() float32 {
	return c.CoffeeDecorator.GetCost() + 2.0
}

func (c *SugarCoffee) GetDescription() string {
	return c.CoffeeDecorator.GetDescription() + " + sugar"
}
