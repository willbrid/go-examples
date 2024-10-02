package main

type Director struct {
	houseBuilder IHouseBuilder
}

func newDirector(b IHouseBuilder) *Director {
	return &Director{
		houseBuilder: b,
	}
}

func (d *Director) setHouseBuilder(b IHouseBuilder) {
	d.houseBuilder = b
}

func (d *Director) buildHouse() House {
	d.houseBuilder.setDoorType()
	d.houseBuilder.setWindowType()
	d.houseBuilder.setNumFloor()

	return d.houseBuilder.getHouse()
}
