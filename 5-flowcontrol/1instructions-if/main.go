package main

import (
	"fmt"
	"strconv"
)

func main() {
	/**
	Une instruction if permet d'exécuter un groupe d'instructions uniquement lorsqu'une expression spécifiée produit
	la valeur booléenne « vrai » lors de son évaluation.

	Le mot-clé `else` permet de créer des clauses supplémentaires dans une instruction `if`.
	Lorsque `else` est combiné avec `if`, les instructions entre accolades ne sont exécutées que si l'expression est vraie et
	que l'expression de la clause précédente est fausse.

	La combinaison else/if peut être répétée pour créer une séquence de clauses, chacune d'elles ne sera exécutée que
	lorsque toutes les expressions précédentes auront produit un résultat faux.

	Le mot-clé else peut également être utilisé pour créer une clause de repli, dont les instructions ne seront exécutées
	que si toutes les expressions `if` et `else/if` de l'instruction produisent des résultats faux. La clause de repli doit être définie
	à la fin de l'instruction et est spécifiée avec le mot-clé `else` sans expression.
	**/

	var kayakPrice float32 = 275.00
	fmt.Println("Price : ", kayakPrice)

	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}

	var kayakPrice1 float32 = 275.00
	if kayakPrice1 > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice1 < 300 {
		fmt.Println("Price is less than 300")
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}

	var kayakPrice2 float32 = 275.00
	if kayakPrice2 > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice2 < 100 {
		fmt.Println("Price is less than 100")
	} else if kayakPrice2 > 200 && kayakPrice2 < 300 {
		fmt.Println("Price is between 200 and 300")
	} else {
		fmt.Println("Price is between 300 and 500")
	}

	/**
	Chaque clause d'une instruction if possède sa propre portée, ce qui signifie que les variables ne sont accessibles qu'à
	l'intérieur de la clause où elles sont définies. Cela signifie également que nous pouvons utiliser le même nom de variable
	à des fins différentes dans des clauses distinctes.
	**/
	var kayakPrice3 float32 = 275.00
	if kayakPrice3 > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice3 < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}

	/**
	Go permet à une instruction if d'utiliser une instruction d'initialisation, qui est exécutée avant que l'expression de
	l'instruction if ne soit évaluée.
	**/
	priceString := "275"
	if kayakPrice4, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price: ", kayakPrice4)
	} else {
		fmt.Println("Error: ", err)
	}
}
