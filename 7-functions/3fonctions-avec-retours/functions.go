package main

/**
Les fonctions peuvent définir des résultats, ce qui permet aux fonctions de fournir à leurs appelants le résultat des opérations.
La fonction déclare son résultat en utilisant un type de données qui suit le paramètre de la fonction.
La fonction calcTax produit un résultat float64, qui est produit par l'instruction `return`.
**/

func calcTax(price float64) float64 {
	return price + price*0.2
}

/**
Une caractéristique des fonctions Go est leur capacité à produire plusieurs résultats.
Les types de résultats produits par la fonction sont regroupés entre parenthèses.
Lorsqu'une fonction définit plusieurs résultats, les valeurs de chaque résultat sont fournies avec le mot-clé `return`,
séparées par des virgules.
**/

func swapValues(first, second int) (int, int) {
	return second, first
}

/**
Le résultat supplémentaire renvoyé par la méthode calcTax1 est une valeur booléenne indiquant si une taxe est due,
ce qui distingue cette information de l'autre résultat.
**/

func calcTax1(price float64) (float64, bool) {
	if price > 100 {
		return price + price*0.2, true
	}

	return 0, false
}

/**
Les résultats d'une fonction peuvent être nommés et se voir attribuer des valeurs pendant l'exécution de la fonction.
Lorsque l'exécution atteint le mot-clé `return`, les valeurs actuellement attribuées aux résultats sont renvoyées.
Les résultats nommés sont définis comme une combinaison d'un nom et d'un type de résultat.

La fonction calcTotalPrice définit des résultats nommés total et tax. Les deux sont des valeurs float64, ce qui signifie qu'on peut
omettre le type de données du premier nom. À l'intérieur de la fonction, les résultats peuvent être utilisés comme des variables classiques.
Le mot-clé `return` est utilisé seul, permettant de renvoyer les valeurs actuelles attribuées aux résultats nommés.
**/

func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, taxDue := calcTax1(price); taxDue {
			total += taxAmount
			tax += taxAmount
		} else {
			total += price
		}
	}

	return
}

func calcTotalPrice1(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}

	return
}
