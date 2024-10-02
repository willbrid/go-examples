package main

type NormalHouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalHouseBuilder() *NormalHouseBuilder {
	return &NormalHouseBuilder{}
}

func (b *NormalHouseBuilder) setWindowType() {
	b.windowType = "Normal Window"
}

func (b *NormalHouseBuilder) setDoorType() {
	b.doorType = "Normal Door"
}

func (b *NormalHouseBuilder) setNumFloor() {
	b.floor = 1
}

func (b *NormalHouseBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
