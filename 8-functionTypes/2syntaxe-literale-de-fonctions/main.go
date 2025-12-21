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

	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}
	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}

	var prizeGiveaway bool = false
	waterCalc1 := priceCalcFactoryWithZeroPrices(100, 0.2, prizeGiveaway)
	prizeGiveaway = true
	soccerCalc1 := priceCalcFactoryWithZeroPrices(50, 0.1, prizeGiveaway)
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc1)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc1)
	}

	var prizeGiveaway1 bool = false
	waterCalc2 := priceCalcFactoryWithZeroPricesPointer(100, 0.2, &prizeGiveaway1)
	prizeGiveaway1 = true
	soccerCalc2 := priceCalcFactoryWithZeroPricesPointer(50, 0.1, &prizeGiveaway1)
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc2)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc2)
	}
}
