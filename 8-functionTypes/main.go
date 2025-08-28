package main

import "fmt"

/*
*
Présentation des types fonctions
L'alias est créé avec le mot-clé "type", suivi d'un nom pour l'alias puis du type et ainsi
le nom d'alias peut être utilisé à la place du type de fonction
*
*/
type calcFunc func(float64) float64 // Définition d'un type fonction

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

// Les types de fonctions peuvent être utilisés de la même manière que n'importe quel autre type, y compris comme arguments
// pour d'autres fonctions
// On peut utiliser les fonctions comme argument à une autre fonction
func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product : ", product, " - Price : ", calculator(price))
}

// Utilisation du type fonction *calcFunc* en tant que type d'argument à une fonction
// via la création d'une alias de type de fonction
func printPriceWithTypeFunction(product string, price float64, calculator calcFunc) {
	fmt.Println("Product : ", product, " - Price : ", calculator(price))
}

// Une fonction peut être le résultat de retour d'une autre fonction
func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}

	return calcWithoutTax
}

// Utilisation du type fonction *calcFunc* en tant que type de retour à une fonction
// via la création d'une alias de type de fonction
func selectCalculatorWithTypeFunction(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}

	return calcWithoutTax
}

/*
*
Utilisation de la syntaxe littérale d'une fonction
La syntaxe littérale de fonction permet de définir des fonctions de manière à ce qu'elles soient spécifiques à une région de code
*
*/
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

// Les fonctions peuvent ne pas être affectées à des variables et peuvent être utilisées comme n'importe quelle autre valeur littérale.
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

/*
*
La fonctionnalité de closure permet à une fonction d'accéder aux variables et aux paramètres du code environnant. Dans ce cas, la fonction
calculatrice (fonction retournée) s'appuie sur les paramètres de la fonction factory (priceCalcFactory).
Lorsque la fonction de calculatrice est invoquée, les valeurs des paramètres de la fonction factory sont utilisées pour produire un résultat
*
*/
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
	/**
	Les opérateurs de comparaison Go ne peuvent pas être utilisés pour comparer des fonctions, mais ils peuvent être utilisés
	pour déterminer si une fonction a été affectée à une variable
	**/
	for product, price := range products {
		var calcFunc func(float64) float64 // Définition d'une variable de type fonction (signature de cette fonction)
		fmt.Println("Function not assigned : ", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}

		fmt.Println("Function not assigned : ", calcFunc == nil)
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

	// La syntaxe littérale de fonctions peut également être utilisée comme arguments pour d'autres fonctions
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

	// Forcer une évaluation précoce
	// La fonction calculatrice de la closure priceCalcFactoryWithEvaluator n'est pas affectée par le changement de valeur
	// de la variable prizeGiveaway
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

	// La plupart des problèmes de closure sont causés par des modifications apportées aux variables après la création d'une fonction.
	// Dans ce cas, l'utilisation d'un pointeur empêchera la copie des valeurs.
	// Le pointeur est suivi lorsque la fonction priceCalcFactoryWithEvaluatorPointer est appelée, ce qui garantit que la
	// valeur actuelle est utilisée.
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
