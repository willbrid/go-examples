package main

import "fmt"

/** Fonction simple **/
func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Println("Price : ", kayakPrice, " - Tax : ", kayakTax)
}

/** Définition et utilisation des paramètres de fonction **/
func printPriceWithParam(product string, price float64, taxRate float64) {
	var taxAmount float64 = price * taxRate
	fmt.Println(product, " price: ", price, " Tax: ", taxAmount)
}

// price et taxRate sont des variables consécutives : elles ont le même type float64
func printPriceWithParamTypeOmit(product string, price, taxRate float64) {
	var taxAmount float64 = price * taxRate
	fmt.Println(product, " price: ", price, " Tax: ", taxAmount)
}

func printPriceWithParamNameOmit(product string, price, _ float64) {
	var taxAmount float64 = price * 0.2
	fmt.Println(product, " price: ", price, " Tax: ", taxAmount)
}

// fonction où c'est uniquement les types qui sont mentionnés
func printPriceWithNoParam(string, float64, float64) {
	fmt.Println("No parameters")
}

/*
*
fonction avec un paramètre dynamique : ce paramètre doit être en dernière position.
Les paramètres dynamiques permettent à une fonction de recevoir un nombre variable d'arguments.
Le paramètre dynamiques est défini par une ellipse (trois points), suivie d'un type.
*
*/
func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product : ", product, " - Supplier : ", supplier)
	}
}

func printSuppliersWithCond(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product : ", product, " - Supplier : ", supplier)
		}
	}
}

func swapValues(first, second *int) {
	fmt.Println("Before swap: ", *first, *second)
	temp := *first
	*first = *second
	*second = temp
	fmt.Println("After swap: ", *first, *second)
}

/** Définition et utilisation des résultats de fonction **/
func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

func swapValuesWithTwoResults(first int, second int) (int, int) { // Fonction avec plusieurs valeurs retournées
	return second, first
}

/*
*
Donner plusieurs significations à un même résultat peut devenir problématique à mesure que les projets évoluent.
La valeur de -1 peut être ambiguë, car elle pourrait indiquer qu'aucune taxe n'est due ou qu'un remboursement de 1$
devrait être effectué.
*
*/
func calcTaxWithCond(price float64) float64 {
	if price > 100 {
		return price * 0.2
	}
	return -1
}

/*
*
Utiliser plusieurs résultats de fonctions est une solution élégante qui résoud le problème issu par le retour
de plusieurs significations à un même résultat.
*
*/
func calcTaxWithTwoResults(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

/*
*
Les résultats d'une fonction peuvent être nommés et recevoir des valeurs lors de son exécution. Lorsque l'exécution atteint
le mot-clé return, les valeurs actuelles attribuées aux résultats sont renvoyées.

La fonction définit les résultats nommés comme variable : total et taxe. Les deux sont des valeurs float64,
ce qui signifie que l'on peut omettre le type de données du premier nom.

Au sein de la fonction, les résultats nommés comme variable peuvent être utilisés comme des variables normales.
*
*/
func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTaxWithTwoResults(price); due {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

func calcTotalPriceWithOneParam(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}

	return
}

/*
*
L'utilisation principale du mot-clé defer est d'appeler des fonctions qui libèrent des ressources, telles que la fermeture de fichiers
ouverts ou de connexions HTTP. Sans le mot-clé defer, l'instruction qui libère la ressource doit apparaître à la fin d'une fonction,
qui peut être composée de plusieurs instructions après la création et l'utilisation de la ressource. Le mot clé defer nous permet
de regrouper les instructions qui créent, utilisent et libèrent la ressource ensemble.

Le mot-clé defer peut être utilisé avec n'importe quel appel de fonction, et une seule fonction peut utiliser le mot-clé defer plusieurs fois.
Juste avant le retour de la fonction, Go effectuera les appels programmés avec le mot-clé defer dans l'ordre dans lequel ils ont été définis.
*
*/
func calcTotalPriceWithDefer(products map[string]float64) (count int, total float64) {
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	fmt.Println("Hello, Functions")

	/** Fonction simple **/
	printPrice()

	/** Définition et utilisation des paramètres de fonction **/
	printPriceWithParam("Kayak", 275, 0.2)
	printPriceWithParam("Lifejacket", 48.95, 0.2)
	printPriceWithParam("Soccer Ball", 19.50, 0.15)

	printPriceWithParamTypeOmit("Kayak", 275, 0.2)
	printPriceWithParamNameOmit("Lifejacket", 48.95, 0.2)
	printPriceWithNoParam("Soccer Ball", 19.50, 0.15)

	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Lifejacket", "Sail Safe Co")

	// Utilisation de tranches comme valeurs pour les paramètres dynamiques
	// Passer un tableau dynamique en paramètre à la position de l'argument dynamique d'une fonction en utilisant ...
	var names []string = []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Soccer Ball", names...)
	printSuppliersWithCond("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliersWithCond("Lifejacket", "Sail Safe Co")
	printSuppliersWithCond("Soccer Ball")

	var val1, val2 int = 10, 20
	fmt.Println("Before calling function : ", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("After calling function : ", val1, val2)

	/** Définition et utilisation des résultats de fonction **/
	var products map[string]float64 = map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var priceWithTax float64 = calcTax(price)
		fmt.Println("Product: ", product, "Price:", priceWithTax)
	}
	for product, price := range products {
		fmt.Println("Product: ", product, "Price:", calcTax(price))
	}

	var val3, val4 int = 10, 20
	fmt.Println("Before calling function : ", val3, val4)
	valResult1, valResult2 := swapValuesWithTwoResults(val3, val4)
	fmt.Println("After calling function : ", valResult1, valResult2)

	// Eviter les fonctions avec plusieurs significations pour un même résultat
	for product, price := range products {
		tax := calcTaxWithCond(price)
		if tax != -1 {
			fmt.Println("Product: ", product, " - Tax:", tax)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}

	// Opter pour les fonctions avec plusieurs retours pour éviter les problèmes issus
	// des fonctions avec plusieurs significations pour un même résultat
	for product, price := range products {
		taxAmount, taxDue := calcTaxWithTwoResults(price)
		if taxDue {
			fmt.Println("Product: ", product, " - Tax:", taxAmount)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}

	for product, price := range products {
		if taxAmount, taxDue := calcTaxWithTwoResults(price); taxDue {
			fmt.Println("Product: ", product, " - Tax:", taxAmount)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}

	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1:", total1, "Tax 1:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2:", total2, "Tax 2:", tax2)

	/**
	Go requiert l'utilisation de toutes les variables déclarées, ce qui peut être gênant lorsqu'une fonction renvoie des valeurs inutiles.
	Pour éviter les erreurs de compilation, l'identifiant vide (le caractère underscore _) peut être utilisé pour indiquer les résultats
	qui ne seront pas utilisés.

	La fonction calcTotalPriceWithOneParam renvoie deux résultats, dont un seul est utilisé. L'identifiant vide est utilisé pour la valeur
	non souhaitée, évitant ainsi une erreur de compilation.
	**/
	_, total3 := calcTotalPriceWithOneParam(products)
	fmt.Println("Total : ", total3)

	_, total4 := calcTotalPriceWithDefer(products)
	fmt.Println("Total : ", total4)
}
