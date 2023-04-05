package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Fonction avec pour type d'un paramètre une classe anonyme
func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name : ", val.name)
}

func main() {
	fmt.Println("Hello, Structs")

	/** Définir et utiliser une structure **/
	fmt.Println("Définir et utiliser une structure")

	type Product struct {
		name, category string
		price          float64
	}
	/**
		Les noms de champ doivent être uniques avec le type struct, ce qui signifie que nous ne pouvons définir qu'un seul
		champ imbriqué pour un type spécifique.
		Si nous avons besoin de définir deux champs du même type, alors nous devrons attribuer un nom à l'un d'entre eux. Exemple avec **Alternate**
	**/
	type StockLevel struct { // Définition d'une classe imbriquant une autre classe
		Product   // Ce champ n'a pas de nom et fait référence à la classe Product
		Alternate Product
		count     int
	}

	var kayak Product = Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	fmt.Println("Produit : ", kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price : ", kayak.price)

	var kayak1 Product = Product{
		name:     "Kayak",
		category: "Watersports",
	}
	fmt.Println("Produit : ", kayak1.name, kayak1.category, kayak1.price)
	kayak1.price = 300
	fmt.Println("Changed price : ", kayak1.price)

	var lifejacket Product
	fmt.Println("Name is zero value : ", lifejacket.name == "")
	fmt.Println("Category is zero value : ", lifejacket.category == "")
	fmt.Println("Price is zero value : ", lifejacket.price == 0)

	var kayak2 Product = Product{ // On se base sur la position des paramètres sans mentionner les noms des champs pour initialiser notre struct Product
		"Kayak",
		"Watersports",
		275.00,
	}
	fmt.Println("Name : ", kayak2.name)
	fmt.Println("Category : ", kayak2.category)
	fmt.Println("Price : ", kayak2.price)

	var stockItem StockLevel = StockLevel{
		Product:   Product{"Kayak", "Watersports", 275.00}, // Les champs intégrés sont accessibles en utilisant le nom du type de champ.
		Alternate: Product{"Lifejacket", "Watersports", 48.95},
		count:     100,
	}
	fmt.Println("Name: ", stockItem.Product.name) // Les champs intégrés sont accessibles en utilisant le nom du type de champ
	fmt.Println("Count : ", stockItem.count)
	fmt.Println("Alt Name : ", stockItem.Alternate.name)

	/** La comparaison est faite sur toutes les valeurs de champ
		Les classes ne peuvent pas être comparées si le type de classe définit des champs avec des types incomparables, tels que des slices
		type Product struct {
	        name, category string
	        price float64
	        otherNames []string
	    }
		**/
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p3 := Product{name: "Kayak", category: "Boats", price: 275.00}
	fmt.Println("p1 == p2 : ", p1 == p2)
	fmt.Println("p1 == p3 : ", p1 == p3)

	type Item struct {
		name     string
		category string
		price    float64
	}
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item := Item{name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("prod == item : ", prod == Product(item)) // On concertit un Item en Product

	prod1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	item1 := Item{name: "Stadium", category: "Soccer", price: 75000}
	writeName(prod1)
	writeName(item1)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  prod1.name,
		ProductPrice: prod1.price,
	})
	fmt.Println(builder.String())
}
