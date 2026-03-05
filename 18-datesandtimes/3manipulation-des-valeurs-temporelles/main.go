package main

import (
	"fmt"
	"time"
)

/**
Le package `time` définit des méthodes pour manipuler les valeurs de type `Time`. Certaines de ces méthodes reposent sur le type `Duration`.

Add(duration) : Cette méthode ajoute la durée spécifiée à la durée et renvoie le résultat.

Sub(time) : Cette méthode renvoie une durée exprimant la différence entre l'heure à laquelle la méthode a été appelée et l'heure fournie en argument.

AddDate(y, m, d) : Cette méthode ajoute le nombre spécifié d'années, de mois et de jours à la durée et renvoie le résultat.

After(time) : Cette méthode renvoie vrai si l'heure à laquelle la méthode a été appelée est postérieure à l'heure fournie en argument.

Before(time) : Cette méthode renvoie vrai si l'heure à laquelle la méthode a été appelée est antérieure à l'heure fournie en argument.

Equal(time) : Cette méthode renvoie vrai si l'heure à laquelle la méthode a été appelée est égale à l'heure fournie en argument.

IsZero() : Cette méthode renvoie vrai si l'heure à laquelle la méthode a été appelée correspond à l'instant zéro, soit le 1er janvier de
l'an 1 à 00:00:00 UTC.

In(loc) : Cette méthode renvoie la valeur de l’heure, exprimée dans le fuseau horaire spécifié.

Location() : Cette méthode renvoie le fuseau horaire associé à l’heure, permettant ainsi d’exprimer une heure dans un fuseau horaire différent.

Round(duration) : Cette méthode arrondit l’heure à l’intervalle le plus proche représenté par une valeur de durée.

Truncate(duration) : Cette méthode arrondit l’heure à l’intervalle inférieur le plus proche représenté par une valeur de durée.
**/

func main() {
	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if err == nil {
		Printfln("After : %v", t.After(time.Now()))
		Printfln("Before : %v", t.Before(time.Now()))
		Printfln("Round : %v", t.Round(time.Hour))
		Printfln("Truncate : %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}

	/**
	Dans cet exemple, les valeurs de type `Time` expriment le même instant dans différents fuseaux horaires.
	La fonction `Equal` tient compte de l'effet des fuseaux horaires, contrairement à l'opérateur d'égalité standard.
	**/

	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	Printfln("Equal Method: %v", t1.Equal(t2))
	// Printfln("Equality Operator: %v", t1 == t2)
}
