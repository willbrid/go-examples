package main

import (
	"fmt"
	"sort"
)

func main() {
	/** Utilisation des pointeurs **/

	var first int = 100
	var second *int = &first

	first++
	*second++

	/*
		Les chaines de caractères précédées d'une barre oblique inverse sont interprétées si la valeur est placée entre guillemets doubles
		"Bonjour\n". Les séquences d'échappement ne sont pas interprétées si la valeur est placée entre guillemets obliques inverses
		`Bonjour\n`.

		Les caractères, les glyphes et les séquences d'échappement sont entourés de guillemets
		simples (le caractère ').
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
