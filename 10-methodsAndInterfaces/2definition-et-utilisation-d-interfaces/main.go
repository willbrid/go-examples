package main

import "fmt"

/**
Les interfaces décrivent un ensemble de méthodes sans en préciser l'implémentation. Si un type implémente toutes les méthodes définies
par l'interface, alors une valeur de ce type peut être utilisée partout où l'interface est autorisée.
Une interface est définie à l'aide du mot-clé `type`, d'un nom, du mot-clé `interface` et d'un corps constitué de signatures
de méthodes encadrées par des accolades.

L'interface Expense décrit deux méthodes.
La première, `getName`, ne prend aucun argument et renvoie une chaîne de caractères.
La seconde, `getCost`, prend un argument booléen et renvoie un résultat de type float64.
**/

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

/**
Les types d'interface peuvent être utilisés pour les variables, les paramètres de fonction et les résultats de fonction.
Cependant il est impossible de définir des méthodes en utilisant des interfaces comme récepteurs.

La fonction `calcTotal` reçoit une tranche contenant des valeurs `Expense`, qui sont traitées à l'aide d'une boucle `for` pour
produire un `total` float64.
**/

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}

	return
}

func main() {
	kayak := Product{"Kayak", "Watersports", 275}
	insurance := Service{"Boat Cover", 12, 89.50}
	fmt.Println("Product kayak :", kayak.name, "Category :", kayak.category, "Price:", kayak.price)
	fmt.Println("Service insurance :", insurance.description, "Price:", insurance.monthlyFee*float64(insurance.durationMonths))

	/**
	Une fois que nous avons implémenté une interface, nous pouvons faire référence aux valeurs via le type d'interface.

	L'on définit une tranche Expense et l'a remplie avec des valeurs Product et Service créées à l'aide de la syntaxe littérale.
	Cette tranche est utilisée dans une boucle for, qui appelle les méthodes getName et getCost sur chaque valeur.

	Les variables dont le type est une interface possèdent deux types : le type statique et le type dynamique. Le type statique est le
	type de l'interface. Le type dynamique est le type de la valeur assignée à la variable qui implémente l'interface, comme Product ou
	Service dans ce cas. Le type statique ne change jamais (le type statique d'une variable Expense est toujours Expense, par exemple),
	mais le type dynamique peut changer en assignant une nouvelle valeur d'un type différent qui implémente l'interface.

	La boucle `for` ne traite que le type statique (Expense) et ne connaît pas (et n'a pas besoin de connaître) le type dynamique de ces valeurs.
	**/
	expenses := []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense :", expense.getName(), "Cost :", expense.getCost(true))
	}
	fmt.Println("Total :", calcTotal(expenses))

	account := Account{
		accountNumber: 12345,
		expenses: []Expense{
			Product{"Kayak", "Watersports", 275},
			Service{"Boat Cover", 12, 89.50},
		},
	}
	for _, expense := range account.expenses {
		fmt.Println("Expense account :", expense.getName(), "Cost :", expense.getCost(true))
	}
	fmt.Println("Total account :", calcTotal(account.expenses), "for account number :", account.accountNumber)

	/**
	Les méthodes définies par les types `Product` et `Service` ont des récepteurs de valeurs, ce qui signifie que les méthodes seront
	appelées avec des copies de la valeur `Product` ou `Service`.

	La valeur de Product a été copiée lors de son affectation à la variable Expense, ce qui signifie que la modification du champ prix
	n'affecte pas le résultat de la méthode `getCost`.
	**/
	product := Product{"Kayak", "Watersports", 275}
	var expense Expense = product // Copie de valeur de Product
	product.price = 100
	fmt.Println("Product field value :", product.price)
	fmt.Println("Expense method result :", expense.getCost(false))

	/**
	Un pointeur vers la valeur de la structure peut être utilisé lors de l'affectation à la variable d'interface.
	L'utilisation d'un pointeur signifie qu'une référence à la valeur de `Product` est assignée à la variable `Expense`, mais cela
	ne modifie pas le type de la variable d'interface, qui reste `Expense`. Nous constaterons l'effet de la référence dans le résultat,
	qui montre que la modification du champ prix est reflétée dans le résultat de la méthode `getCost`.
	**/
	product1 := Product{"Kayak", "Watersports", 275}
	var expense1 Expense = &product1
	product1.price = 100
	fmt.Println("Product1 field value :", product1.price)
	fmt.Println("Expense1 method result :", expense1.getCost(false))

	/**
	Nous pouvons imposer l'utilisation de références en spécifiant des récepteurs de pointeurs lors de l'implémentation des méthodes
	d'interface. Cela signifie que le type `Article` n'implémente pas l'interface `Expense` car les méthodes requises ne sont plus définies.
	C'est le type `*Article` qui implémente l'interface, ce qui implique que les pointeurs vers des valeurs `Article` peuvent être traités
	comme des valeurs `Expense` et non comme des valeurs classiques.
	**/

	article := Article{"Kayak", "Watersports", 275}
	var expense2 Expense = &article
	article.price = 100
	fmt.Println("Article1 field value :", article.price)
	fmt.Println("Expense2 method result :", expense2.getCost(false))

	/**
	Les valeurs d'interface peuvent être comparées à l'aide des opérateurs de comparaison Go. Deux valeurs d'interface sont égales
	si elles ont le même type dynamique et si tous leurs champs sont égaux.

	Les deux premières valeurs de Expense ne sont pas égales. En effet, leur type dynamique est un pointeur, et deux pointeurs ne sont égaux
	que s'ils pointent vers la même adresse mémoire. Les deux dernières valeurs de Expense sont égales car ce sont des structures simples
	ayant les mêmes valeurs de champs.
	**/
	var e1 Expense = &Article{name: "Kayak"}
	var e2 Expense = &Article{name: "Kayak"}
	var e3 Expense = Service{description: "Boat Cover"}
	var e4 Expense = Service{description: "Boat Cover"}
	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)

	/**
	Les vérifications d'égalité d'interface peuvent également provoquer des erreurs d'exécution si le type dynamique n'est pas comparable.
	Exemple avec une structure définie de la manière :

		type ServiceFeature struct {
	    	description string
	    	durationMonths int
	    	monthlyFee float64
	    	features []string
		}
	**/
}
