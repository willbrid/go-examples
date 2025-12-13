package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/**
	Avec le mot clé const, il existe deux façons de définir des constantes en Go : les constantes typées et les constantes non typées.
	**/
	fmt.Println("Value: ", rand.Int())
	fmt.Println("Hello Will")
	const price float32 = 275.00
	const tax float32 = 27.50
	const quantity = 2
	fmt.Println("Constante : prix : ", price, " Tax : ", tax, " quantite : ", quantity)

	/** Utilisation des constances **/

	const price1, tax1 float32 = 275, 27.50
	const quantity1, inStock1 = 2, true
	fmt.Println("Constante : prix : ", price1, " Tax : ", tax1, " quantite : ", quantity1, " en stock : ", inStock1)

	/**
	Le mot-clé "iota" peut être utilisé pour créer une série de constantes entières non typées successives sans avoir besoin
	de leur attribuer des valeurs individuelles.
	**/
	const (
		foot1 = iota
		foot2
		foot3
	)
	fmt.Println("Foot1 :", foot1, " - Fooot2", foot2, " - Foot3", foot3)
}
