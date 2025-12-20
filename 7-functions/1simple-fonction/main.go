package main

import "fmt"

/*
*
Les fonctions sont des groupes d'instructions qui peuvent être utilisées et réutilisées comme une seule action.
Le fichier main.go contient désormais deux fonctions. La nouvelle fonction, `printPrice`, définit deux variables et appelle la
fonction `Printf` du package fmt. La fonction `main` est le point d'entrée de l'application : son exécution commence et se termine ici.
En Go, les fonctions doivent être définies entre accolades, l'accolade ouvrante devant figurer sur la même ligne que le mot-clé `func`
et le nom de la fonction.
**/

/*
*
Les fonctions sont définies par le mot-clé `func`, suivi du nom de la fonction, de parenthèses et d'un bloc de code entre accolades.
*
*/
func printPrice() {
	kayakPrice := 275.00
	kayakTax := kayakPrice * 0.2
	fmt.Printf("price %.2f, taxe de %.2f.\n", kayakPrice, kayakTax)
}

func main() {
	/**
	La fonction `main` appelle la fonction `printPrice`, ce qui se fait au moyen d'une instruction qui spécifie le nom de la fonction,
	suivi de parenthèses.
	**/
	fmt.Println("About to call function")
	printPrice()
	fmt.Println("Function complete")
}
