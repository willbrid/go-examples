package main

import "fmt"

const (
	SHOE  = "shoe"
	SHIRT = "shirt"
)

func getProduct(productType string) (IProduct, error) {
	if productType == SHOE {
		return newShoe(), nil
	}
	if productType == SHIRT {
		return newShirt(), nil
	}
	return nil, fmt.Errorf("wrong product type")
}
