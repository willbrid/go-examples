package main

import (
	"fmt"
	"sort"
)

func main() {
	/** Utilisation des pointeurs
	Un pointeur est une variable dont sa valeur est une adresse mémoire.
	**/

	/**
	Le type d'un pointeur dépend du type de la variable dont il est issu, précédé d'un astérisque (*).
	La variable nommée « second » est de type *int, car elle a été créée en appliquant l'opérateur d'adresse (&) à la variable nommée « first »,
	dont la valeur est un entier (int). Le type *int indique qu'il s'agit d'une variable dont la valeur est une adresse mémoire
	stockant un entier.

	L'expression "suivre un pointeur" signifie lire la valeur à l'adresse mémoire pointée par le pointeur, et elle est effectuée à
	l'aide d'un astérisque (le caractère *).
	**/
	var first int = 100
	var second *int = &first

	first++
	*second++
	fmt.Println("First : ", first)
	fmt.Println("Second : ", second)
	fmt.Println("Second- : ", *second)

	var myNewPointer *int = second
	*myNewPointer++

	fmt.Println("First : ", first)
	fmt.Println("Second : ", *second)

	var first1 int = 100
	var second1 *int

	// Les pointeurs définis mais auxquels aucune valeur n'a été attribuée ont la valeur zéro nil
	fmt.Println(second1) // nil
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
