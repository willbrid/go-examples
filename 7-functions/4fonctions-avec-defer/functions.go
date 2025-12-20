package main

import "fmt"

/**
Le mot-clé `defer` est utilisé pour planifier un appel de fonction qui sera exécuté immédiatement avant la fin de cette fonction (avant
l'instruction `return` de la fonction).

Le mot-clé defer est utilisé avant l'appel de fonction.

Le principal usage du mot-clé `defer` est d'appeler des fonctions qui libèrent des ressources, comme la fermeture de fichiers ouverts
ou de connexions HTTP. Sans `defer`, l'instruction de libération de la ressource doit figurer à la fin de la fonction, ce qui peut
représenter plusieurs instructions après la création et l'utilisation de la ressource. Le mot-clé `defer` permet de regrouper les
instructions de création, d'utilisation et de libération de la ressource.

Le mot-clé `defer` peut être utilisé avec n'importe quel appel de fonction.
**/

func printPrice(product string, price float64, taxRate float64) {
	defer fmt.Println("final report") // Pas d'instruction return alors defer s'exécutera à la fin de la fonction
	taxAmount := price * taxRate
	fmt.Println(product, "price :", price, "tax :", taxAmount)
}

func calcTotalPrice(products map[string]float64) (count int, total float64) {
	count = len(products)
	defer fmt.Println("total products :", count) // Avec instruction return alors defer s'exécutera avant l'instruction return

	for _, price := range products {
		total += price
	}

	return
}
