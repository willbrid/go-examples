package main

import "fmt"

/**
Les paramètres permettent à une fonction de recevoir des valeurs de données lors de son appel, ce qui permet de modifier son comportement.
Les paramètres sont définis par un nom suivi d'un type. Plusieurs paramètres sont séparés par des virgules.
**/

func printPrice(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price :", price, "tax :", taxAmount)
}

/**
Le type d'un paramètre peut être omis lorsque les paramètres adjacents ont le même type.
Dans l'exemple suivant, les paramètres `price` et `taxRate` ont tous deux le type `float64`,
donc le type n'est pas répété pour le second paramètre.
**/

func printPrice1(product string, price, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price :", price, "tax :", taxAmount)
}

/**
Le caractère de soulignement (_) peut être utilisé pour les paramètres définis par une fonction mais non utilisés dans
les instructions de code de cette fonction.

Le caractère de soulignement est appelé identificateur vide. Il s'agit d'un paramètre pour lequel une valeur doit être fournie lors de
l'appel de la fonction, mais dont la valeur est inaccessible à l'intérieur du bloc de code de la fonction.
C'est un moyen utile d'indiquer qu'un paramètre n'est pas utilisé dans une fonction,
ce qui peut se produire lors de l'implémentation des méthodes requises par une interface.
**/

func printPrice2(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price :", price, "tax :", taxAmount)
}

/**
Les fonctions peuvent également omettre les noms de tous leurs paramètres.
Les paramètres sans nom ne sont pas accessibles au sein de la fonction, et cette fonctionnalité est principalement
utilisée en conjonction avec des interfaces.
**/

func printPrice3(string, float64, float64) {
	fmt.Println("No parameters")
}

/**
Un paramètre variadique accepte un nombre variable de valeurs, ce qui peut faciliter l'utilisation des fonctions.
Les paramètres variadiques permettent à une fonction de recevoir un nombre variable d'arguments.
Le paramètre variadique est défini par des points de suspension (...), suivie d'un type.
Le paramètre variadique doit être le dernier paramètre défini par la fonction, et un seul type peut être utilisé,
comme le type chaîne de caractères dans cet exemple ci-dessous.

La fonction intégrée `len` est utilisée pour identifier si la tranche `suppliers` est vide. Cependant l'on pouvait aussi vérifier
avec la valeur `nil`.
**/

func printSuppliers(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product :", product, "Supplier :", supplier)
		}
	}
}

/**
Par défaut, Go copie les valeurs utilisées comme arguments afin que les modifications soient limitées à l'intérieur de la fonction.
Go permet aux fonctions de recevoir des pointeurs, ce qui modifie ce comportement par défaut.
Cela permet aux fonctions de modifier les valeurs des variables situées en dehors de leur portée.
**/

func swapValues(first, second *int) {
	temp := *first
	*first = *second
	*second = temp
}
