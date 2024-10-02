package main

type IHouseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

const (
	NORMAL_HOUSE = "normal"
	IGLOO_HOUSE  = "igloo"
)

func getHouseBuilder(houseBuilderType string) IHouseBuilder {
	if houseBuilderType == NORMAL_HOUSE {
		return newNormalHouseBuilder()
	}

	if houseBuilderType == IGLOO_HOUSE {
		return newIglooHouseBuilder()
	}

	return nil
}
