package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

/**
Go permet l'utilisation d'une interface vide (c'est-à-dire une interface ne définissant aucune méthode) de représenter n'importe quel type,
ce qui peut s'avérer utile pour regrouper des types disparates ne partageant aucune caractéristique commune.
**/

func main() {
	/**
	L'interface vide est utilisée dans une syntaxe littérale, définie avec le mot-clé `interface` et des accolades vides.
	L'interface vide (`interface{}`) a pour alias `any` à partir de Go 1.18.

	L'interface vide représente tous les types, y compris les types intégrés et toutes les structures et interfaces qui ont été définies.
	**/

	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}
	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product :", value.name, "Price :", value.price)
		case *Product:
			fmt.Println("Product Pointer :", value.name, "Price :", value.price)
		case Service:
			fmt.Println("Service :", value.description, "Price :", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person :", value.name, "City :", value.city)
		case *Person:
			fmt.Println("Person Pointer :", value.name, "City :", value.city)
		case string, bool, int:
			fmt.Println("Built-in type :", value)
		default:
			fmt.Println("Default :", value)
		}
	}

	/**
	Appel à la fonction processItem qui a pour type de paramètre une interface vide.
	**/
	for _, item := range data {
		processItem(item)
	}

	// Appel à la fonction processItems qui a pour type de paramètre un paramètre variadique sur l'interface.
	processItems(data...)
}
