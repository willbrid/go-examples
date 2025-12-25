package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

/**
Le mot-clé `type` permet de créer un alias pour le type `[]Product`, nommé `ProductList`. Ce type peut être utilisé pour définir des méthodes,
soit directement pour des récepteurs de type valeur, soit via un pointeur.
**/

type ProductList []Product

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

type Supplier struct {
	name, city string
}

func newSupplier(name, city string) *Supplier {
	return &Supplier{name, city}
}

/**
Les méthodes sont des fonctions qui peuvent être appelées via une valeur et constituent un moyen pratique d'exprimer des fonctions
qui opèrent sur un type spécifique.

Les méthodes sont définies comme des fonctions, utilisant le même mot-clé `func`, mais avec l'ajout d'un récepteur, qui désigne un
paramètre spécial, lequel est le type sur lequel la méthode opère.
Le type du récepteur pour cette méthode est `*Product` et porte le nom `p`, qui peut être utilisé dans la méthode comme
n'importe quel paramètre de fonction normal.

Lors de l'appel de la méthode `calcTax`, les arguments sont fournis comme pour une fonction classique.
**/

func (p *Product) printDetails() {
	fmt.Println("Product name :", p.name, "category :", p.category, "price :", p.calcTax(0.2, 100))
}

func (p Product) printCalculatedPrice() {
	fmt.Println("Product calculated price :", p.calcTax(0.2, 100))
}

/**
Les méthodes peuvent définir des paramètres et des résultats, tout comme les fonctions classiques.
La méthode calcTax définit les paramètres de taux et de seuil et renvoie un résultat de type float64.
**/

func (p *Product) calcTax(rate, threshold float64) float64 {
	if p.price > threshold {
		return p.price + p.price*rate
	}

	return p.price
}

/**
Chaque combinaison de nom de méthode et de type de récepteur doit être unique, quels que soient les autres paramètres définis.
Donc deux méthodes peuvent avoir le même nom mais leur combinaison nom méthode et récepteur est différente.
**/

func (s *Supplier) printDetails() {
	fmt.Println("Supplier :", s.name, "City:", s.city)
}

/**
Les méthodes ne sont pas limitées aux structures, car le mot-clé `type` peut être utilisé pour créer des alias vers n'importe quel `type`,
et des méthodes peuvent être définies pour l'alias.
**/

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}

	return totals
}

/**
Le résultat de la fonction getProducts est []Product, qui doit être converti en ProductList avec une conversion explicite,
permettant ainsi d'utiliser la méthode définie sur l'alias.
**/

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}
