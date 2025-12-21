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

/**
Les fonctions définies à l'aide de la syntaxe littérale peuvent faire référence à des variables du code environnant,
une fonctionnalité connue sous le nom de `closure`.

La fonction `priceCalcFactory` a pour rôle de créer des fonctions de calcul pour une combinaison spécifique de `threshold` et de `rate`.
Ses entrées sont le seuil et le taux d'une catégorie, et sa sortie est une fonction qui calcule les prix appliqués à cette catégorie.
Le code de la fonction `priceCalcFactory` utilise la syntaxe littérale pour définir une fonction de calcul contenant le code commun
nécessaire au calcul.

La fonctionnalité closure assure la liaison entre la fonction `priceCalcFactory` et la fonction de calcul.
Cette dernière utilise deux variables pour produire un résultat.
La fonctionnalité closure permet à une fonction d'accéder aux variables et paramètres du code environnant.
Ici, la fonction de calcul utilise les paramètres de la fonction `priceCalcFactory`. Lorsqu'elle est appelée, les valeurs de
ces paramètres servent à calculer le résultat.
**/

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + price*rate
		}

		return price
	}
}

/**
Les variables sur lesquelles la fonction de calcul est appelée sont évaluées à chaque appel de la fonction,
ce qui signifie que des modifications effectuées en dehors de la fonction peuvent affecter les résultats qu'elle produit.

Lors de l'appel à la fonction de calcul, la valeur de zeroPrices est copiée et il ne s'agit pas toujours de sa valeur courante.
**/

func priceCalcFactoryWithZeroPrices(threshold, rate float64, zeroPrices bool) calcFunc {
	return func(price float64) float64 {
		if zeroPrices {
			return 0
		} else if price > threshold {
			return price + price*rate
		} else {
			return price
		}
	}
}

/**
La plupart des problèmes liés au closure sont dus à des modifications apportées aux variables après la création d'une fonction.
Ces problèmes peuvent être résolus à l'aide de l'utilisation de pointeur qui empêche la copie des valeurs.

Lors de l'appel à la fonction de calcul, c'est la valeur courante du pointeur zeroPrices qui sera utilisée.
**/

func priceCalcFactoryWithZeroPricesPointer(threshold, rate float64, zeroPrices *bool) calcFunc {
	return func(price float64) float64 {
		if *zeroPrices {
			return 0
		} else if price > threshold {
			return price + price*rate
		} else {
			return price
		}
	}
}
