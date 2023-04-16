package main

import (
	"fmt"
	"packages/store" // On précise le module puis le nom du package
)

func main() {
	fmt.Println("Hello, Packages and Modules")

	// Création d'un custom package
	fmt.Println("Création d'un custom package")

	var product1 store.Product = store.Product{
		Name:     "Kayak",
		Category: "Watersports",
	}
	fmt.Println("Name : ", product1.Name)
	fmt.Println("Category:", product1.Category)

	var product2 *store.Product = store.NewProduct("Kayak", "Watersports", 279)
	fmt.Println("Name : ", product2.Name)
	fmt.Println("Category : ", product2.Category)
	fmt.Println("Price : ", product2.Price())
	fmt.Println("Price standard : ", product2.PriceStandardTax())
}
