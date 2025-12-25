package main

import "fmt"

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		fmt.Println("Product name :", p.name, "category :", p.category, "price :", p.price)
	}

	/**
	Ce qui distingue les méthodes des fonctions classiques, c'est la manière dont elles sont appelées.
	Les méthodes sont appelées via une valeur dont le type correspond à celui du récepteur. Dans l'exemple ci-dessous l'on a
	la valeur `*Product` générée par la boucle `for` pour appeler la méthode `printDetails` pour chaque valeur d'une tranche.
	**/
	for _, p := range products {
		p.printDetails()
	}

	someProducts := []*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}
	for _, p := range someProducts {
		p.printDetails()
	}

	suppliers := []*Supplier{
		newSupplier("Acme Co", "New York City"),
		newSupplier("BoatCo", "Chicago"),
	}
	for _, s := range suppliers {
		s.printDetails()
	}

	/**
	Une méthode dont le récepteur est un type pointeur peut également être appelée via une valeur classique du type sous-jacent ;
	ainsi, une méthode de type *Product, par exemple, peut être utilisée avec une valeur Product.
	La variable `kayak` se voit attribuer la valeur `Product`, mais elle est utilisée avec la méthode `printDetails`,
	dont le récepteur est `*Product`.

	Le processus inverse est également vrai : une méthode recevant une valeur peut donc être appelée à l’aide d’un pointeur.

	Cette fonctionnalité nous permet d'écrire des méthodes en fonction du comportement souhaité, en utilisant des pointeurs pour
	éviter la copie de valeurs ou pour permettre à la méthode de modifier le récepteur.

	L'un des effets de cette fonctionnalité est que les types valeur et pointeur sont considérés comme identiques en matière de
	surcharge de méthodes, ce qui signifie qu'une méthode nommée printDetails dont le type de récepteur est Product entrera en conflit
	avec une méthode printDetails dont le type de récepteur est *Product.
	**/
	kayak := Product{"Kayak", "Watersports", 275}
	kayak.printDetails()
	kayakPtr := &Product{"Kayak", "Watersports", 275}
	kayakPtr.printCalculatedPrice()

	otherProducts := ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for category, total := range otherProducts.calcCategoryTotals() {
		fmt.Println("Category :", category, "Total :", total)
	}

	/**
	Nous ne pourrons pas toujours recevoir des données du type requis pour appeler une méthode définie pour un alias,
	par exemple lors du traitement des résultats d'une fonction. Dans ce cas, nous pouvons effectuer une conversion de type.
	**/
	someOtherProducts := ProductList(getProducts())
	for category, total := range someOtherProducts.calcCategoryTotals() {
		fmt.Println("Category :", category, "Total :", total)
	}
}
