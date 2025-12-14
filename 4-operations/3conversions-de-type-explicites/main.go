package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	Go n’autorise pas le mélange de types dans les opérations et ne convertit pas automatiquement les types, sauf
	dans le cas des constantes non typées.
	**/

	/**
	Conversions de type explicites
	Une conversion explicite transforme une valeur pour en changer le type.
	**/
	kayak := 275
	soccerBall := 19.50
	total1 := float64(kayak) + soccerBall
	fmt.Println(total1)

	/**
		Les conversions explicites ne sont possibles que si la valeur peut être représentée dans le type cible.
		Cela signifie que nous pouvons convertir entre types numériques et entre chaînes de caractères et runes, mais d'autres combinaisons,
		comme la conversion de valeurs entières en valeurs booléennes, ne sont pas prises en charge.

	    Il convient d'être vigilant lors du choix des valeurs à convertir, car les conversions explicites peuvent entraîner une
		perte de précision pour les valeurs numériques ou provoquer des dépassements de capacité.
	**/
	total2 := kayak + int(soccerBall)
	fmt.Println("Total2 : ", total2)
	fmt.Println("Total2 : ", int8(total2))

	/**
		math.Ceil : cette fonction renvoie le plus petit entier supérieur à la valeur à virgule flottante spécifiée.
		            Par exemple, le plus petit entier supérieur à 27,1 est 28.
	    math.Floor : cette fonction renvoie le plus grand entier inférieur à la valeur à virgule flottante spécifiée.
		             Par exemple, le plus grand entier inférieur à 27,1 est 26.
		math.Round : cette fonction arrondit la valeur à virgule flottante spécifiée à l'entier le plus proche.
		math.RoundToEven : cette fonction arrondit la valeur à virgule flottante spécifiée à l'entier pair le plus proche.
	**/
	soccerBall1 := math.Ceil(soccerBall)
	soccerBall2 := math.Floor(soccerBall)
	soccerBall3 := math.Round(soccerBall)
	soccerBall4 := math.RoundToEven(soccerBall)
	fmt.Println("Ceil :", soccerBall1)
	fmt.Println("Floor :", soccerBall2)
	fmt.Println("Round :", soccerBall3)
	fmt.Println("RoundToEven :", soccerBall4)
	total3 := kayak + int(math.Round(soccerBall3))
	fmt.Println("Total3 : ", total3)
}
