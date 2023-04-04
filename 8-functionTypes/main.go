package main

import "fmt"

/** Présentation des types fonctions **/
type calcFunc func(float64) float64 // Définition d'un type fonction

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product : ", product, " - Price : ", calculator(price))
}

// Utilisation du type fonction *calcFunc* en tant que type d'argument à une fonction
func printPriceWithTypeFunction(product string, price float64, calculator calcFunc) {
	fmt.Println("Product : ", product, " - Price : ", calculator(price))
}

func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}

	return calcWithoutTax
}

// Utilisation du type fonction *calcFunc* en tant que type de retour à une fonction
func selectCalculatorWithTypeFunction(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}

	return calcWithoutTax
}

/** Utilisation de la syntaxe de la fonction littérale **/
func selectCalculatorWithLitteral(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 {
			return price + (price * 0.2)
		}

		return withTax
	}

	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

func selectCalculatorWithDirectLitteral(price float64) calcFunc {
	if price > 100 {
		return func(price float64) float64 {
			return price + (price * 0.2)
		}
	}

	return func(price float64) float64 {
		return price
	}
}

/**
La fonctionnalité de closure permet à une fonction d'accéder aux variables et aux paramètres du code environnant. Dans ce cas, la fonction
calculatrice (fonction retournée) s'appuie sur les paramètres de la fonction factory (priceCalcFactory).
Lorsque la fonction de calculatrice est invoquée, les valeurs des paramètres de la fonction factory sont utilisées pour produire un résultat
**/
func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}

		return price
	}
}

var prizeGiveaway = false

func priceCalcFactoryWithGlobal(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if prizeGiveaway {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}

		return price
	}
}

func priceCalcFactoryWithEvaluator(threshold, rate float64, zeroPrices bool) calcFunc {
	return func(price float64) float64 {
		if zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}

		return price
	}
}

func priceCalcFactoryWithEvaluatorPointer(threshold, rate float64, zeroPrices *bool) calcFunc {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else if price > threshold {
			return price + (price * rate)
		}

		return price
	}
}

func main() {
	fmt.Println("Hello, Function Types")

	/** Présentation des types fonctions **/
	fmt.Println("Présentation des types fonctions")
	var products map[string]float64 = map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}
	for product, price := range products {
		var calcFunc func(float64) float64 // Définition d'une variable de type fonction (signature de cette fonction)
		fmt.Println("Function assigned : ", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}

		fmt.Println("Function assigned : ", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product : ", product, " - Price : ", totalPrice)
	}

	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}

	for product, price := range products {
		printPriceWithTypeFunction(product, price, selectCalculatorWithTypeFunction(price))
	}

	/** Utilisation de la syntaxe de la fonction littérale **/
	fmt.Println("Utilisation de la syntaxe de la fonction littérale")
	for product, price := range products {
		printPrice(product, price, selectCalculatorWithLitteral(price))
	}

	for product, price := range products {
		printPrice(product, price, selectCalculatorWithDirectLitteral(price))
	}

	for product, price := range products {
		printPrice(product, price, func(price float64) float64 {
			return price + (price * 0.2)
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

	prizeGiveaway = false
	waterCalc1 := priceCalcFactoryWithGlobal(100, 0.2)
	prizeGiveaway = true
	soccerCalc1 := priceCalcFactoryWithGlobal(50, 0.1)
	// Ici les fonctions waterCalc1 et soccerCalc1 ne sont pas encore évaluées
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc1)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc1)
	}

	// La fonction calculatrice de la closure priceCalcFactoryWithEvaluator n'est pas affectée par le changement de valeur de la variable prizeGiveaway
	prizeGiveaway = false
	waterCalc2 := priceCalcFactoryWithEvaluator(100, 0.2, prizeGiveaway)
	prizeGiveaway = true
	soccerCalc2 := priceCalcFactoryWithEvaluator(50, 0.1, prizeGiveaway)
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc2)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc2)
	}

	// Le pointeur est suivi lorsque la fonction calculatrice est appelée, ce qui garantit que la valeur actuelle est utilisée.
	prizeGiveaway = false
	waterCalc3 := priceCalcFactoryWithEvaluatorPointer(100, 0.2, &prizeGiveaway)
	prizeGiveaway = true
	soccerCalc3 := priceCalcFactoryWithEvaluatorPointer(50, 0.1, &prizeGiveaway)
	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc3)
	}
	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc3)
	}
}
