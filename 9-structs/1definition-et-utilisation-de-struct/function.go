package main

import "fmt"

/**
Les types de structures anonymes sont définis sans utiliser de nom.
La fonction writeName utilise un type structure anonyme comme paramètre, ce qui signifie qu'elle peut accepter n'importe quel type structure définissant l'ensemble de champs spécifié.
**/

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name :", val.name, "Category :", val.category, "Price :", val.price)
}
