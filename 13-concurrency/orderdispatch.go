package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers []string = []string{"Alice", "Bob", "Charlie", "Dora"}

/**
La fonction intégrée close accepte un canal comme argument et est utilisée pour indiquer qu'aucune autre valeur
ne sera envoyée via le canal. Les récepteurs peuvent vérifier si un canal est fermé lorsqu'ils demandent une valeur
**/

func DispatchOrders(channel chan DispatchNotification) {
	var orderCount int

	rand.NewSource(time.Now().UTC().UnixNano())
	orderCount = rand.Intn(3) + 2
	fmt.Println("Order count : ", orderCount)

	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
	}
	close(channel)
}
