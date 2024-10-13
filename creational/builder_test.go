package creational

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	director := new(Director)
	builder := HouseBuilder{}
	director.setBuilder(builder)
	house := director.getHouse()
	if house.Foundation != "foundation" {
		t.Error("expected foundation")
	}
	if house.Roof != "roof" {
		t.Error("expected roof")
	}
	if house.Wall != "wall" {
		t.Error("expected wall")
	}
}
