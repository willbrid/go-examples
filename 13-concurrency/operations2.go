package main

import (
	"fmt"
	"time"
)

// Utilisation d'une goroutine sans canal : ici la variable storeTotal a pour valeur 0 car le total d'une catégorie n'est pas reçue de la goroutine

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
