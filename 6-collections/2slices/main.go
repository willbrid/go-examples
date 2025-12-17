package main

import "fmt"

func main() {
	/**
	La meilleure façon de concevoir les tranches est de les considérer comme un tableau de longueur variable, car elles sont utiles
	lorsque nous ne connaissons pas le nombre de valeurs à stocker ou lorsque ce nombre évolue dans le temps.
	Une façon de définir une tranche est d'utiliser la fonction intégrée `make`.
	**/

	names := make([]string, 3)
	names[0] = "kayak"
	names[1] = "lifejacket"
	names[2] = "paddle"
	fmt.Println("names :", names)
}
