package main

import (
	"fmt"

	"composition/store"
)

/**
Go ne prend pas en charge les classes ni l'héritage et privilégie la composition. Cependant, malgré ces différences, la composition
permet de créer des hiérarchies de types, même si elle est différente. Le point de départ consiste à définir un type de structure
et une méthode.
**/

func main() {
	product1 := store.NewProduct("Kayak", "Watersports", 275)
	product2 := &store.Product{Name: "Lifejacket", Category: "Watersports"}
	for _, p := range []*store.Product{product1, product2} {
		fmt.Println("Name :", p.Name, "Category :", p.Category, "Price :", p.Price(0.2))
	}
}
