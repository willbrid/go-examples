package main

import "fmt"

/**
Les valeurs des paramètres sont fournies comme arguments lors de l'appel de la fonction, ce qui permet de fournir des valeurs différentes
à chaque appel. Les arguments sont placés entre parenthèses après le nom de la fonction, séparés par des virgules et dans l'ordre où
les paramètres ont été définis.
Les valeurs utilisées comme arguments doivent correspondre aux types des paramètres définis par la fonction.
**/

func main() {
	printPrice("Kayak", 275, 0.2)
	printPrice1("Lifejacket", 48.95, 0.2)
	printPrice2("Soccer Ball", 19.50, 0.15)
	printPrice3("Soccer Ball", 19.50, 0.15)

	printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	printSuppliers("Lifejacket", "Sail Safe Co")
	// Go permet d'omettre entièrement les arguments des paramètres variadiques.
	printSuppliers("Soccer Ball")
	/**
	On peut utiliser de tranches comme valeurs pour les paramètres variadiques en le faisant suivre par des points de suspension (...).
	**/
	suppliers := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Kayak", suppliers...)

	val1, val2 := 10, 20
	fmt.Println("Before calling function ", val1, val2)
	swapValues(&val1, &val2)
	fmt.Println("After calling function ", val1, val2)
}
