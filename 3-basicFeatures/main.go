package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	fmt.Println("Types, valeurs et pointeurs de base")

	/** Types de données de base **/
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
	fmt.Println(price + tax)

	var price3 float32
	fmt.Println(price3)
	price3 = 275.00
	fmt.Println(price3)

	var price4, tax4 = 275.00, 27.50
	fmt.Println(price4 + tax4)

	/** Utilisation des pointeurs **/
	var first int = 100
	var second *int = &first

	first++
	*second++
	/*
		Les chaines de caractères précédées d'une barre oblique inverse sont interprétées si la valeur est placée entre guillemets doubles
		"Bonjour\n". Les séquences d'échappement ne sont pas interprétées si la valeur est placée entre guillemets obliques inverses
		`Bonjour\n`.
	*/
	fmt.Println("First : ", first)
	fmt.Println("Second : ", second)
	fmt.Println("Second- : ", *second)

	var myNewPointer *int
	myNewPointer = second
	*myNewPointer++

	fmt.Println("First : ", first)
	fmt.Println("Second : ", *second)

	var first1 int = 100
	var second1 *int

	fmt.Println(second1)
	second1 = &first1
	fmt.Println(second1)

	var third **int = &second
	fmt.Println(third)
	fmt.Println(*third)
	fmt.Println(**third)

	var names [3]string = [3]string{"Alice", "Charlie", "Bob"}
	var secondName string = names[1]
	var secondPosition *string = &names[1]

	fmt.Println(secondName)
	fmt.Println(*secondPosition)
	sort.Strings(names[:])
	fmt.Println(secondName)
	fmt.Println(*secondPosition)
}
