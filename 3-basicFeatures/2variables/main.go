package main

import (
	"fmt"
)

func main() {
	/** Utilisation des variables **/

	var price float32 = 275.00
	var tax float32 = 27.50
	var quantity float32 = 2
	var inStock bool = true

	fmt.Println(20 + 20)
	fmt.Println(price + tax)
	fmt.Println("Total : ", 2*quantity*(price+tax))
	fmt.Println("En stock : ", inStock)

	fmt.Println(price + tax)
	price = 300
	fmt.Println(price + tax)

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
	var price1 float32
	fmt.Println(price1)
	price1 = 275.00
	fmt.Println(price1)

	var price3, tax3 = 275.00, 27.50
	fmt.Println(price3 + tax3)

	/** La déclaration de variable abrégée offre une syntaxe concise pour déclarer des variables. **/
	price5 := 275.00
	fmt.Println(price5)

	/**
	En Go, il est illégal de définir une variable et de ne pas l'utiliser.
	Dans ces situations, Go fournit l'identifiant vide, utilisé pour désigner une valeur qui ne sera pas utilisée.
	**/
	price4, tax4, inStock4, _ := 275.00, 27.50, true, true
	var _ = "Alice"
	fmt.Println("Total :", price4+tax4)
	fmt.Println("In Stock :", inStock4)

	/*
		Les chaines de caractères précédées d'une barre oblique inverse sont interprétées si la valeur est placée entre guillemets doubles
		"Bonjour\n". Les caractères d'échappement ne sont pas interprétées si la valeur est placée entre guillemets obliques inverses
		`Bonjour\n`.

		Les caractères, les glyphes et les caractères d'échappement sont entourés de guillemets simples (le caractère ').
	*/
}
