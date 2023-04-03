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

func printPriceWithParamTypeOmit(product string, price, taxRate float64) { // price et taxRate sont des variables consécutives : elles ont le même type float64
	var taxAmount float64 = price * taxRate
	fmt.Println(product, " price: ", price, " Tax: ", taxAmount)
}

func printPriceWithParamNameOmit(product string, price, _ float64) {
	var taxAmount float64 = price * 0.2
	fmt.Println(product, " price: ", price, " Tax: ", taxAmount)
}

func printPriceWithNoParam(string, float64, float64) { // c'est uniquement les types qui sont mentionnés
	fmt.Println("No parameters")
}

func printSuppliers(product string, suppliers ...string) { // fonction avec un paramètre dynamique : ce paramètre doit être en dernière position
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

	var names []string = []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Soccer Ball", names...) // Passer un tableau dynamique en paramètre à la position de l'argument dynamique d'une fonction en utilisant ...

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
}
