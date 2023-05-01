package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64) // Creation d'un canal où envoyé et recevoir les valeurs de type float64
	var value float64

	for category, group := range data {
		// storeTotal += group.TotalPrice(category)
		// go group.TotalPrice(category) // Création d'une goroutine qui exécute de manière asynchrone un appel à une méthode
		go group.TotalPrice(category, channel)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("-- Starting to receive from channel")
	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending")
		value = <-channel
		fmt.Println("-- channel read complete", value)
		storeTotal += value // La flèche est placée devant le canal pour en recevoir une valeur provenant de ce canal.
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64

	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}

	// fmt.Println(category, "subtotal:", ToCurrency(total))
	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total // La flèche est placée devant une variable pour envoyer une valeur.
	fmt.Println(category, "channel send complete")
}
