package main

import "fmt"

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	/**
	Lorsque la fonction est appelée, le résultat peut être affecté à une variable.
	**/
	for product, price := range products {
		priceWithTax := calcTax(price)
		fmt.Println("#1 Product :", product, "Price :", priceWithTax)
	}

	/**
	Les résultats des fonctions peuvent être utilisés directement dans les expressions. Dans l'exemple ci-dessous,
	la variable est omise et la fonction calcTax est appelée directement pour produire un argument pour la fonction fmt.PrintLn.
	**/
	for product, price := range products {
		fmt.Println("#2 Product :", product, "Price :", calcTax(price))
	}

	val1, val2 := 10, 20
	fmt.Println("Before calling function ", val1, val2)
	val1, val2 = swapValues(val1, val2)
	fmt.Println("After calling function ", val1, val2)

	/**
	Les deux résultats sont obtenus dans une instruction distincte, mais l'obtention de plusieurs résultats est parfaitement
	adaptée à la prise en charge d'une instruction d'initialisation par l'instruction if.
	**/
	for product, price := range products {
		taxAmount, taxDue := calcTax1(price)
		if taxDue {
			fmt.Println("#3 Product :", product, "Price :", taxAmount)
		} else {
			fmt.Println("#3 Product :", product, "No tax due")
		}
	}

	/**
	Les deux résultats sont obtenus en appelant la fonction calcTax1 dans l'instruction d'initialisation, et le résultat booléen
	est ensuite utilisé comme expression de l'instruction.
	**/
	for product, price := range products {
		if taxAmount, taxDue := calcTax1(price); taxDue {
			fmt.Println("#4 Product :", product, "Price :", taxAmount)
		} else {
			fmt.Println("#4 Product :", product, "No tax due")
		}
	}

	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1 :", total1, "Tax 1 :", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2 :", total2, "Tax 2 :", tax2)

	/**
	La fonction calcTotalPrice1 renvoie deux résultats, dont un seul est utilisé. L'identificateur vide est utilisé pour la valeur
	indésirable, évitant ainsi une erreur de compilation.
	**/
	_, total3 := calcTotalPrice1(products)
	fmt.Println("Total 3 :", total3)
}
