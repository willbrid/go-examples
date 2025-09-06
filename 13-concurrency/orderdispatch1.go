package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification1 struct {
	Customer string
	*Product
	Quantity int
}

var Customers1 []string = []string{"Alice", "Bob", "Charlie", "Dora"}

/*
*
La fonction intégrée close accepte un canal comme argument et est utilisée pour indiquer qu'aucune autre valeur
ne sera envoyée via le canal. Les récepteurs peuvent vérifier si un canal est fermé lorsqu'ils demandent une valeur
L'emplacement de la flèche spécifie la direction du canal. Lorsque la flèche suit le mot-clé chan (chan<-), le canal ne peut être utilisé que pour envoyer.
Le canal peut être utilisé pour recevoir uniquement si la flèche précède le mot-clé chan (<-chan).
*
*/
func DispatchOrders1(channel chan<- DispatchNotification1) {
	var orderCount int
	// var notification DispatchNotification1

	rand.NewSource(time.Now().UTC().UnixNano())
	orderCount = rand.Intn(3) + 2
	fmt.Println("Order count : ", orderCount)

	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification1{
			Customer: Customers1[rand.Intn(len(Customers1)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		/**
		L'instruction de reception n'est pas possibles car la mention chan<- force la canal a envoyé uniquement.
		if i == 1 {
			notification = <-channel
			fmt.Println("Read : ", notification.Customer)
		}**/
	}
	close(channel)
}
