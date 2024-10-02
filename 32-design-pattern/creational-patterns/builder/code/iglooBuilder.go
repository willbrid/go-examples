package main

type IglooHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooHouseBuilder() *IglooHouseBuilder {
	return &IglooHouseBuilder{}
}

func (b *IglooHouseBuilder) setWindowType() {
	b.windowType = "Igloo Window"
}

func (b *IglooHouseBuilder) setDoorType() {
	b.doorType = "Igloo Door"
}

func (b *IglooHouseBuilder) setNumFloor() {
	b.floor = 3
}

func (b *IglooHouseBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
