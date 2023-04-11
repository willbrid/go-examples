package main

import "fmt"

// définition d'une interface avec les signatures de méthode
// Les signatures de méthode consistent en un nom, des paramètres et des types de résultats.
type Expense interface {
	getName() string
	getCost(annual bool) float64
}

// Définition d'un type alias à []Product
type ProductList []Product

/** Cette fonction est une méthode à la classe Product car nous avons défini un récepteur (receiver) (récepteur sur Product : (product *Product))
avant le nom de la fonction. Donc un récepteur est le type sur lequel la méthode opère.
**/
func (product *Product) printDetails() {
	fmt.Println("Name : ", product.name, " - Category : ", product.category, " - Price : ", product.price)
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}

	return product.price
}

func (product *Product) printDetailsWithTax() {
	// Ici on accède à la méthode calcTax depuis la méthode printDetailsWithTax
	fmt.Println("Name : ", product.name, " - Category : ", product.category, " - Price : ", product.calcTax(0.2, 100))
}

/**
Chaque combinaison de nom de méthode et de type de récepteur doit être unique, quels que soient les autres paramètres définis.
Nous avons donc deux méthodes avec le même nom printDetails mais de récepteur différent.
**/
func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier : ", supplier.name, " - City : ", supplier.city)
}

func (product Product) printDetailsWithTax1() {
	fmt.Println("Name : ", product.name, " - Category : ", product.category, " - Price : ", product.calcTax(0.2, 100))
}

// Définition d'une méthode à l'alias ProductList
func (products *ProductList) calcCategoryTotals() map[string]float64 {
	var totals map[string]float64 = make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}

	return totals
}

func getProducts() []Product {
	return []Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
}

// Utilisation d'une interface dans une fonction
func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total = item.getCost(true)
	}

	return
}

func main() {
	fmt.Println("Hello Methods and interfaces of struct")

	var products []*Product = []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for _, p := range products {
		fmt.Println("Name : ", p.name, " - Category : ", p.category, " - Price : ", p.price)
	}

	/** Définir et utiliser des méthodes de classe
		Les méthodes sont des fonctions qui sont appelées sur une classe et ont accès à tous les champs définis par le type de la valeur.
	**/
	fmt.Println("Définir et utiliser des méthodes de classe")
	for _, p := range products {
		p.printDetails() // On invoque la méthode printDetails depuis un objet de la classe Product
	}

	for _, p := range products {
		p.printDetailsWithTax()
	}

	// Chaque combinaison de nom de méthode et de type de récepteur doit être unique, quels que soient les autres paramètres définis
	var suppliers []*Supplier = []*Supplier{
		{"Acme Co", "New York City"},
		{"BoatCo", "Chicago"},
	}
	for _, s := range suppliers {
		s.printDetails()
	}

	/** Une méthode dont le récepteur est un type pointeur peut également être appelée via une valeur normale du type sous-jacent,
	// ce qui signifie qu'une méthode dont le type est *Product, par exemple, peut être utilisée avec un objet Product
	**/
	var kayak Product = Product{"Kayak", "Watersports", 275}
	kayak.printDetails()

	/**
	Le processus inverse est également vrai : une méthode qui reçoit une valeur peut être invoquée à l'aide d'un pointeur.
	Cette fonctionnalité signifie que nous pouvons écrire des méthodes en fonction de la façon dont nous souhaitons qu'elles se comportent,
	en utilisant des pointeurs pour éviter la copie de valeur ou pour permettre au récepteur d'être modifié par une méthode.

	Une méthode nommée printDetails dont le type de récepteur est Product entrera en conflit avec une méthode printDetails
	dont le type de récepteur est *Product.
	**/
	var kayak1 *Product = &Product{"Kayak", "Watersports", 275}
	kayak1.printDetailsWithTax1()

	var products1 ProductList = ProductList{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.50},
	}
	for category, total := range products1.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}
	var products2 ProductList = ProductList(getProducts()) // Conversion en type []Product en type ProductList
	for category, total := range products2.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

	// Création du type Service dans un fichier séparé
	var insurance Service = Service{"Boat Cover", 12, 89.50}
	fmt.Println("Service : ", insurance.description, " - Price : ", insurance.monthlyFee*float64(insurance.durationMonths))

	/** Définition et utilisation des interfaces **/
	/**
	Les interfaces décrivent un ensemble de méthodes sans spécifier l'implémentation de ces méthodes.
	Si un type implémente toutes les méthodes définies par l'interface, alors une valeur de ce type peut être utilisée partout où l'interface est autorisée.
	**/
	// Utilisation d'une interface
	var expenses []Expense = []Expense{
		Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense : ", expense.getName(), " - Cost : ", expense.getCost(true))
	}

	/**
	Les variables dont le type est une interface ont deux types : le type statique et le type dynamique. Le type statique est le type d'interface.
	Le type dynamique est le type de valeur attribuée à la variable qui implémente l'interface, telle que Product ou Service dans ce cas.
	Le type statique ne change jamais—le type statique d'une variable Expense est toujours Expense, par exemple—mais le type dynamique peut changer
	en affectant une nouvelle valeur d'un type différent qui implémente l'interface.
	La boucle for ne traite que du type statique—Expense—et ne connaît pas (et n'a pas besoin de connaître) le type dynamique de ces valeurs.
	L'utilisation de l'interface nous permet de regrouper des types dynamiques disparates et d'utiliser les méthodes communes spécifiées par le
	type d'interface statique.
	**/
	// Utilisation d'une interface dans une méthode
	fmt.Println("Total : ", calcTotal(expenses))
}
