package main

import (
	"fmt"
	"math"
)

func main() {
	/** Utilisation des opérateurs arithmétiques **/

	var price, tax float32 = 275.00, 27.40
	var sum, difference, product, quotient float32

	sum = price + tax
	difference = price - tax
	product = price * tax
	quotient = price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)

	/**
	Go autorise le dépassement de capacité des valeurs entières par un retour à la ligne, plutôt que de signaler une erreur.
	Les valeurs à virgule flottante peuvent dépasser la valeur infinie positive ou négative.
	**/
	var intVal = math.MaxFloat64
	var floatVal = math.MaxFloat64
	fmt.Println("overflow1 :", intVal*2)
	fmt.Println("overflow2 :", floatVal*2)
	fmt.Println("la variable floatVal est-elle infinie :", math.IsInf((floatVal*2), 0))

	// Go propose l'opérateur % qui renvoie le reste de la division d'un entier par un autre et peut renvoyer des valeurs négatives.
	var posResult int = 3 % 2
	var negResult int = -3 % 2
	var absResult float64 = math.Abs(float64(negResult))
	fmt.Println(posResult)
	fmt.Println(negResult)
	fmt.Println(absResult)

	/**
	Go propose un ensemble d'opérateurs pour incrémenter et décrémenter des valeurs numériques.
	Ces opérateurs peuvent être appliqués aux nombres entiers et à virgule flottante.
	Les opérateurs ++ et -- incrémentent ou décrémentent une valeur de un. Les opérateurs += et -= incrémentent ou décrémentent
	une valeur d'une valeur spécifiée.
	**/
	var value float32 = 10.2
	value++
	fmt.Println(value)
	value += 2
	fmt.Println(value)
	value -= 2
	fmt.Println(value)
	value--
	fmt.Println(value)

	// L'opérateur + peut être utilisé pour concaténer des chaînes de caractères afin de produire des chaînes plus longues.
	var greeting string = "Hello"
	var language string = "Go"
	var combinedString = greeting + ", " + language
	fmt.Println(combinedString)
}
