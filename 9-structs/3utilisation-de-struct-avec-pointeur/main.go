package main

import "fmt"

/**
L'affectation d'une structure à une nouvelle variable ou l'utilisation d'une structure comme paramètre de fonction crée une nouvelle valeur
qui copie les valeurs des champs.
Comme pour les autres types de données, il est possible de créer des références aux valeurs de structures à l'aide de pointeurs.
**/

func main() {
	/**
	L'on utilise une esperluette pour créer un pointeur vers la variable product1 et l'on affecte l'adresse à product2, dont le type
	devient *Product, c'est-à-dire un pointeur vers une valeur de type Product. Notons qu'on doit utiliser des parenthèses pour faire suivre
	le pointeur vers la valeur de la structure, puis lire la valeur du champ `name`.
	**/
	product1 := Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	product2 := &product1
	product1.name = "Original Kayak"
	fmt.Println("product1 :", product1.name)
	fmt.Println("product2 :", (*product2).name)

	calcTax(&product1)
	fmt.Println("product1 name :", product1.name, "category :", product1.category, "price :", product1.price)

	/**
	L'exemple ci-dessus utilise un pointeur en deux étapes.
	La première étape consiste à créer une valeur et à l'affecter à une variable.
	La deuxième étape consiste à utiliser l'opérateur d'adresse pour créer un pointeur.

	Il n'est pas nécessaire d'assigner une valeur de structure à une variable avant de créer un pointeur, et l'opérateur d'adresse
	peut être utilisé directement avec la syntaxe littérale des structures.
	**/
	product3 := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTax(product3)
	fmt.Println("product3 name :", product3.name, "category :", product3.category, "price :", product3.price)

	product4 := &Product{
		name:     "Kayak",
		category: "Watersports",
		price:    275,
	}
	calcTaxWithResult(product4)
	fmt.Println("product4 name :", product4.name, "category :", product4.category, "price :", product4.price)

	products := [2]*Product{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Hat", "Skiing", 42.50),
	}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}

	/**
	Les champs définis par la structure `Supplier` sont accessibles via le champ défini par la structure Article.
	**/
	acme := &Supplier{"Acme Co", "New York"}
	articles := [2]*Article{
		newArticle("Kayak", "Watersports", 275, acme),
		newArticle("Hat", "Skiing", 42.50, acme),
	}
	for _, a := range articles {
		fmt.Println("Name :", a.name, "Supplier :", a.Supplier.name, a.Supplier.city)
	}
	a1 := newArticle("Kayak", "Watersports", 275, acme)
	a2 := copyArticle(a1)
	a1.name = "Original Kayak"
	a1.Supplier.name = "BoatCo"
	for _, a := range []Article{*a1, a2} {
		fmt.Println("Name :", a.name, "Supplier :", a.Supplier.name, a.Supplier.city)
	}

	/**
	La valeur nulle pour un type struct est une valeur struct dont les champs sont initialisés à leur type nul.
	La valeur nulle d'un pointeur vers une struct est nil.
	Le problème provient d’une tentative d’accès au champ `name` de la structure imbriquée pourtant elle est nulle.
	Dans ce cas, comme la valeur zéro de la structure imbriquée est assimilée à une valeur nulle, ce qui entraîne l’erreur d’exécution observée.
	Pour éviter cette erreur, on initialise le type imbriqué sans produire de valeur de champs (`&Supplier{}`).
	**/
	var art Article = Article{Supplier: &Supplier{}}
	var artPtr *Article
	fmt.Println("Value :", art.name, art.category, art.price, art.Supplier.name)
	fmt.Println("Pointer :", artPtr)
}
