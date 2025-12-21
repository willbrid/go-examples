package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product :", product, "Price :", calculator(price))
}

/**
La syntaxe littérale des fonctions permet de définir des fonctions de manière à ce qu'elles soient spécifiques à une région de code.

La syntaxe littérale omet un nom, de sorte que le mot-clé `func` est suivi des paramètres, du type de résultat et du bloc de code.

La syntaxe littérale crée une fonction qui peut être utilisée comme n'importe quelle autre valeur, notamment en l'assignant à une variable,
comme illustré dans l'exemple ci-dessous. Le type d'une fonction littérale est défini par sa signature ; le nombre et le type de ses
paramètres doivent donc correspondre au type de la variable.
Cette fonction littérale possède une signature correspondant à l'alias de type calcFunc, avec un paramètre et un résultat de type float64.
Les fonctions littérales peuvent également être utilisées avec la syntaxe courte de déclaration de variable.

Les fonctions sont traitées comme n'importe quelle autre valeur, mais la fonction qui ajoute la taxe n'est accessible que par le biais
de la variable `withTax`, elle-même accessible uniquement à l'intérieur du bloc de code de l'instruction `if`.
**/

func selectCalculator(price float64) calcFunc {
	if price > 100 {
		var withTax calcFunc = func(price float64) float64 {
			return price + price*0.2
		}
		return withTax
	}

	withoutTax := func(price float64) float64 {
		return price
	}
	return withoutTax
}

/**
Les fonctions n'ont pas besoin d'être assignées à des variables et peuvent être utilisées comme n'importe quelle autre valeur littérale.
Le mot-clé return est appliqué directement à la fonction, sans assigner la fonction à une variable.
**/

func selectCalculator1(price float64) calcFunc {
	if price > 100 {
		return func(price float64) float64 {
			return price + price*0.2
		}
	}

	return func(price float64) float64 {
		return price
	}
}
