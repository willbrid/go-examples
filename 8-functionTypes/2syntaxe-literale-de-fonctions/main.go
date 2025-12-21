package main

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator1(price))
	}

	/**
	Les fonctions littérales peuvent également être utilisées comme arguments d'autres fonctions.
	Le dernier argument de la fonction printPrice est exprimé en utilisant la syntaxe littérale et sans assigner la fonction à une variable.
	**/
	for product, price := range products {
		printPrice(product, price, func(price float64) float64 {
			return price + price*0.2
		})
	}
}
