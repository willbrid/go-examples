package main

import (
	"fmt"
)

func main() {
	/** Utilisation des variables **/

	var price2 float32 = 275.00
	var tax2 float32 = 27.50
	var quantity2 float32 = 2
	var inStock2 bool = true

	fmt.Println(20 + 20)
	fmt.Println(price2 + tax2)
	fmt.Println("Total : ", 2*quantity2*(price2+tax2))
	fmt.Println("En stock : ", inStock2)

	fmt.Println(price2 + tax2)
	price2 = 300
	fmt.Println(price2 + tax2)

	/**
	Les variables définies sans assignation de valeur se voient attribuer la valeur zéro pour le type spécifié.
	Les valeurs zéro pour les types de base :
	int : 0
	unit : 0
	byte : 0
	float64 : 0
	bool : false
	string : "" (chaîne vide)
	rune : 0
	*/
	var price3 float32
	fmt.Println(price3)
	price3 = 275.00
	fmt.Println(price3)

	var price4, tax4 = 275.00, 27.50
	fmt.Println(price4 + tax4)

	/** La déclaration de variable abrégée offre une syntaxe concise pour déclarer des variables. **/
	price5 := 275.00
	fmt.Println(price5)
}
