package main

import "fmt"

/**
L'interface vide peut être utilisée comme type pour un paramètre de fonction, permettant ainsi d'appeler une fonction avec
n'importe quelle valeur.
**/

func processItem(item any) {
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
L'interface vide peut également être utilisée pour les paramètres variadiques, ce qui permet d'appeler une fonction avec un nombre
quelconque d'arguments, chacun pouvant être de n'importe quel type.
**/

func processItems(items ...any) {
	for _, item := range items {
		processItem(item)
	}
}
