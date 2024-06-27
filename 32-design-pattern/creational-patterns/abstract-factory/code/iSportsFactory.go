package main

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

const (
	AdidasBrand string = "adidas"
	NikeBrand   string = "nike"
)

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case AdidasBrand:
		return &Adidas{}, nil
	case NikeBrand:
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}
