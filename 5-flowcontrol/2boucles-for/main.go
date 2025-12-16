package main

import "fmt"

func main() {
	/**
	Le mot-clé `for` est utilisé pour créer des boucles qui exécutent des instructions de manière répétée. Les boucles `for`
	les plus simples se répètent indéfiniment, sauf si elles sont interrompues par le mot-clé `break`.
	**/
	var counter int = 0
	for {
		fmt.Println("Counter: ", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	// La condition peut être intégrée à la syntaxe de la boucle for
	// La condition est spécifiée entre le mot-clé « for » et l’accolade ouvrante qui entoure les instructions de la boucle.
	var counter1 int = 0
	for counter1 <= 3 {
		fmt.Println("Counter1: ", counter1)
		counter1++
	}

	/**
	Les boucles peuvent être définies avec des instructions supplémentaires qui sont exécutées avant la
	première itération de la boucle (appelée instruction d'initialisation) et après
	chaque itération (l'instruction d'incrémentation ou de décrémentation).
	**/
	for counter2 := 0; counter2 <= 3; counter2++ {
		fmt.Println("Counter2: ", counter2)
	}

	for counter3 := 0; true; counter3++ {
		fmt.Println("Counter3: ", counter3)
		if counter3 > 3 {
			break
		}
	}

	// Le mot-clé `continue` peut être utilisé pour interrompre l'exécution des instructions de la boucle for pour la valeur actuelle
	// et passer à l'itération suivante.
	for counter4 := 0; counter4 <= 3; counter4++ {
		if counter4 == 1 {
			continue
		}
		fmt.Println("Counter4: ", counter4)
	}

	// Utilisation de `for` avec l'instruction `range` pour parcourir les chaines, les entiers, les arrays, les slices et les maps
	/**
	Cet exemple ci-dessous parcourt une chaîne de caractères que la boucle `for` traite comme une séquence de valeurs de type `rune`,
	chacune représentant un caractère. À chaque itération de la boucle, des valeurs sont assignées à deux variables :
	l’indice courant dans la séquence et la valeur (de type rune) à cet indice.
	**/
	var product string = "Kayak"
	for index, character := range product {
		fmt.Println("Index: ", index, " Character: ", string(character))
	}

	// for avec range peut également être utilisé pour parcourir les entiers.
	fmt.Println("Range over integer")
	for value := range 30 {
		fmt.Println("Integer element Value: ", value)
	}

	fmt.Println("Range over slices")
	var products []string = []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}

	// Nous pouvons omettre la variable de valeur de l'instruction `for` range si nous n'avons besoin que des valeurs d'index.
	for index := range product {
		fmt.Println("Index: ", index)
	}

	// L'identifiant vide (`_`) peut être utilisé lorsque nous avons uniquement besoin des valeurs de la séquence et non des indices.
	for _, character := range product {
		fmt.Println("Character: ", string(character))
	}
}
