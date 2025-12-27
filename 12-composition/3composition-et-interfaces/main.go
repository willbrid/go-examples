package main

import (
	"composition/store"
	"fmt"
)

/**
Go utilise des interfaces pour décrire des méthodes pouvant être implémentées par plusieurs types. Go prend en compte les méthodes promues
pour déterminer si un type est conforme à une interface, ce qui évite de dupliquer des méthodes déjà présentes via un champ imbriqué.

Nous utilisons le type d'interface `ItemForSale` pour créer une valeur de type map, qui est remplie d'éléments conformes à l'interface.

Le type `Product` est directement conforme à l'interface `ItemForSale` car il existe une méthode `Price` qui correspond à la signature
spécifiée par l'interface et qui possède un récepteur `*Product`.
Il n'existe pas de méthode `Price` qui accepte un récepteur `*Boat`, mais Go prend en compte la méthode `Price` promue à partir du champ
imbriqué du type `Boat`, qu'il utilise pour satisfaire aux exigences de l'interface.
**/

func main() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer Ball", "Soccer", 19.50),
	}
	for key, p := range products {
		fmt.Println("Key :", key, "Price :", p.Price(0.2))
	}

	/**
	L'instruction `case` effectue une assertion de type lorsqu'un seul type est spécifié, même si cela peut entraîner une duplication
	lors du traitement de chaque type.
	**/
	for key, p := range products {
		switch item := p.(type) {
		case *store.Product:
			fmt.Println("#1 Name :", item.Name, "Category :", item.Category, "Price :", item.Price(0.2))
		case *store.Boat:
			fmt.Println("#1 Name :", item.Name, "Category :", item.Category, "Price :", item.Price(0.2))
		default:
			fmt.Println("#1 Key :", key, "Price :", p.Price(0.2))
		}
	}

	/**
	Une solution alternative pour éviter une duplication de traitement de chaque type consiste à définir des méthodes d'interface permettant
	d'accéder aux valeurs des propriétés. Cela peut se faire en ajoutant des méthodes à une interface existante ou en définissant une
	interface distincte.
	**/
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println("#2 Name :", item.GetName(), "Category :", item.GetCategory(), "Price :", item.(store.ItemForSale).Price(0.2))
		default:
			fmt.Println("#2 Key :", key, "Price :", p.Price(0.2))
		}
	}

	for key, p := range products {
		switch item := p.(type) {
		case store.DescribableItem:
			fmt.Println("#3 Name :", item.GetName(), "Category :", item.GetCategory(), "Price :", item.Price(0.2))
		default:
			fmt.Println("#3 Key :", key, "Price :", p.Price(0.2))
		}
	}
}
