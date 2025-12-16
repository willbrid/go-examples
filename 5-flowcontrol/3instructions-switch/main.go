package main

import (
	"fmt"
)

func main() {
	/**
	Utilisation de l'instruction `switch`

	L'instruction `switch` offre une autre façon de contrôler le flux d'exécution, en faisant correspondre le résultat d'une
	expression à une valeur spécifique.
	Le mot-clé `switch` est suivi d'une valeur ou d'une expression qui produit un résultat utilisé pour la comparaison.
	Les comparaisons sont effectuées par rapport à une série d'instructions `case`, chacune spécifiant une valeur.

	Le mot-clé `case` est suivi d'une valeur, de deux points et d'une ou plusieurs instructions à exécuter lorsque
	la valeur de comparaison correspond à la valeur de l'instruction `case`.

	Avec l'instruction switch, plusieurs valeurs peuvent être spécifiées à l'aide d'une liste séparée par des virgules.
	**/
	var character rune = 'a'
	switch character {
	case 'a':
		fmt.Println("Character is a")
	case 'b':
		fmt.Println("Character is b")
	}

	var product1 string = "Kayak"
	for index, character := range product1 {
		switch character {
		case 'K':
			fmt.Println("K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		case 'a', 'k':
			fmt.Println("a or k at position", index)
		}
	}

	/**
	    switch : on peut utiliser l'instruction break parmi les instructions d'un cas pour sortir immédiatement du switch
	**/
	var product2 string = "Kayak"
	for index, character := range product2 {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}

	/**
	  	En Go un `case` ne “tombe pas” automatiquement dans le `case` suivant. Par défaut : un seul case est exécuté et
		dès qu’un `case` correspond, le switch s’arrête.
	    L'instruction `fallthrough` exécutée dans un cas oblige l'exécution des instructions du prochain cas sans re-tester sa condition.
	**/
	var product3 string = "Kayak"
	for index, character := range product3 {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}

	/**
	    Switch : on peut avoir plusieurs valeurs de cas; elles doivent être séparées par une virgule
		L'instruction `default` est exécutée lorsqu'aucune des instructions `case` ne correspond à la valeur de l'instruction switch.
	**/
	var product4 string = "Kayak"
	for index, character := range product4 {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("k at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}

	/**
	  Une instruction switch peut être définie avec une instruction d'initialisation, ce qui peut s'avérer utile pour préparer
	  la valeur de comparaison afin qu'elle puisse être référencée dans les instructions case.

	  Switch : il y'a duplication de l'expression -> counter5 / 2
	**/
	for counter5 := 0; counter5 < 20; counter5++ {
		switch counter5 / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime value: ", counter5/2)
		default:
			fmt.Println("Non-prime value: ", counter5/2)
		}
	}

	/**
	  Switch : on utilise une variable d'initialisaiton (déclaration courte de variable) pour éviter les duplications d'expression.
	**/
	for counter5 := 0; counter5 < 20; counter5++ {
		switch val := counter5 / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value: ", val)
		default:
			fmt.Println("Non-prime value: ", val)
		}
	}

	/**
		Switch : lorsque la valeur (ou expression) de comparaison est omise au niveau du switch, on utilise les expressions
		de comparaison dans les cas
	**/
	for counter6 := 0; counter6 < 10; counter6++ {
		switch {
		case counter6 == 0:
			fmt.Println("Zero value")
		case counter6 < 3:
			fmt.Println(counter6, "is < 3")
		case counter6 >= 3 && counter6 < 7:
			fmt.Println(counter6, "is >= 3 && < 7")
		default:
			fmt.Println(counter6, "is >= 7")
		}
	}
}
