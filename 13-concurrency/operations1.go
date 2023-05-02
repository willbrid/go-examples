package main

import (
	"fmt"
)

func CalcStoreTotal1(data ProductData) {
	var storeTotal float64

	for category, group := range data {
		storeTotal += group.TotalPrice1(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice1(category string) float64 {
	var total float64

	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}

	fmt.Println(category, "subtotal:", ToCurrency(total))
	return total
}
