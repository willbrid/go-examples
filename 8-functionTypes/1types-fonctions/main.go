package main

import "fmt"

/**
En Go, les fonctions possèdent un type de données, ce qui signifie qu'elles peuvent être assignées à des variables et
utilisées comme paramètres, arguments et résultats de fonctions.
**/

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	/**
	Les types fonctions sont spécifiés par le mot-clé `func`, suivi des types de paramètres entre parenthèses, puis des types de résultats.
	On parle alors de signature de la fonction. S'il y a plusieurs résultats, leurs types sont également placés entre parenthèses.
	Le type de fonction de l'exemple ci-dessous décrit une fonction qui accepte un argument de type `float64` et produit un résultat de
	type `float64`.

	La variable `calcFunc` définie dans l'exemple ci-dessous peut se voir attribuer n'importe quelle valeur correspondant à son type,
	c'est-à-dire toute fonction possédant le nombre et le type d'arguments et de résultats requis. Pour affecter une fonction spécifique
	à une variable, on utilise le nom de la fonction.

	Une fois qu'une fonction a été assignée à une variable, elle peut être appelée comme si le nom de la variable était le nom de la fonction.

	Les opérateurs de comparaison Go ne peuvent pas être utilisés pour comparer des fonctions, mais ils peuvent être utilisés
	pour déterminer si une fonction a été assignée à une variable.
	La valeur zéro pour les types de fonctions est `nil`, et les instructions de comparaison au niveau de `fmt` de
	l'exemple ci-dessous utilisent l'opérateur d'égalité pour déterminer si une fonction a été affectée à la variable `calcFunc`.
	**/

	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned :", calcFunc == nil)
		if price > 100 {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned :", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product :", product, "Price :", totalPrice)
	}

	/**
	L'important est que la fonction printPrice ignore et ne se soucie pas si elle reçoit la fonction calcWithTax ou calcWithoutTax
	via le paramètre calculator. La fonction printPrice sait seulement qu'elle pourra appeler la fonction calculator avec un argument
	de type float64 et recevoir un résultat de type float64, car il s'agit du type de la fonction du paramètre.
	**/

	for product, price := range products {
		if price > 100 {
			printPrice(product, price, calcWithTax)
		} else {
			printPrice(product, price, calcWithoutTax)
		}
	}

	for product, price := range products {
		calculator := selectCalculator(price)
		printPrice(product, price, calculator)
	}

	for product, price := range products {
		printPriceWithAlias(product, price, selectCalculatorWithAlias(price))
	}
}
