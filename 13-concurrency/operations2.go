package main

import (
	"fmt"
	"time"
)

/**
Utilisation d'une goroutine sans canal : ici la variable storeTotal a pour valeur 0 car le total d'une catégorie n'est pas reçue de la goroutine

Une goroutine est créée à l'aide du mot-clé go suivi de la fonction ou de la méthode qui doit être exécutée de manière asynchrone.
Lorsque le runtime Go rencontre le mot-clé go, il crée une nouvelle goroutine et l'utilise pour exécuter la fonction ou la méthode spécifiée.
Cela modifie l'exécution du programme car, à tout moment, il existe plusieurs goroutines, chacune exécutant son propre ensemble d'instructions.
Ces instructions sont exécutées simultanément, ce qui signifie simplement qu'elles sont exécutées en même temps.
**/

func CalcStoreTotal2(data ProductData) {
	var storeTotal float64

	for category, group := range data {
		go group.TotalPrice2(category) // Création d'une goroutine qui exécute de manière asynchrone un appel à une méthode
	}

	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice2(category string) float64 {
	var total float64

	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println(category, "subtotal:", ToCurrency(total))
	return total
}
