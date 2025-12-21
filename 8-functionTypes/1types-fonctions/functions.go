package main

import "fmt"

/**
L'utilisation des types fonctions peut s'avérer verbeuse et répétitive, ce qui rend le code difficile à lire et à maintenir.
Go prend en charge les alias de types, qui permettent d'attribuer un nom à la signature d'une fonction afin de ne pas avoir à
spécifier systématiquement les types des paramètres et du résultat.

L'alias est créé avec le mot-clé type, suivi d'un nom pour l'alias, puis du type.
**/

type calcFunc func(float64) float64

func calcWithTax(price float64) float64 {
	return price + price*0.2
}

func calcWithoutTax(price float64) float64 {
	return price
}

/**
Les types fonctions peuvent être utilisés de la même manière que n'importe quel autre type, comme arguments pour d'autres fonctions.

La fonction `printPrice` définit trois paramètres : les deux premiers acceptent des chaînes de caractères et des nombres à
virgule flottante 64 bits (float64). Le troisième paramètre, nommé calculator, reçoit une fonction qui prend une valeur de
type float64 et produit un résultat de type float64.
**/

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Product :", product, "Price :", calculator(price))
}

/**
Les fonctions peuvent aussi être des résultats, ce qui signifie que la valeur renvoyée par une fonction est une autre fonction.
La fonction `selectCalculator` reçoit une valeur float64 et renvoie une fonction de type `func(float64) float64`.
**/

func selectCalculator(price float64) func(float64) float64 {
	if price > 100 {
		return calcWithTax
	}

	return calcWithoutTax
}

func printPriceWithAlias(product string, price float64, calculator calcFunc) {
	fmt.Println("Product :", product, "Price :", calculator(price))
}

func selectCalculatorWithAlias(price float64) calcFunc {
	if price > 100 {
		return calcWithTax
	}

	return calcWithoutTax
}
