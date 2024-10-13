package creational

type House struct {
	Foundation string
	Roof       string
	Wall       string
}

type HouseBuilder struct {
	house House
}

func (b *HouseBuilder) buildFoundation() *HouseBuilder {
	b.house.Foundation = "foundation"
	return b
}

func (b *HouseBuilder) buildRoof() *HouseBuilder {
	b.house.Roof = "roof"
	return b
}

func (b *HouseBuilder) buildWall() *HouseBuilder {
	b.house.Wall = "wall"
	return b
}

func (b *HouseBuilder) build() House {
	return b.house
}

type Director struct {
	builder HouseBuilder
}

func (d *Director) setBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *Director) getHouse() House {
	d.builder.buildFoundation().buildRoof().buildWall()
	return d.builder.build()
}
