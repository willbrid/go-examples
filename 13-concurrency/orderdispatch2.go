package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification2 struct {
	Customer string
	*Product
	Quantity int
}

var Customers2 []string = []string{"Alice", "Bob", "Charlie", "Dora"}

/*
*
La fonction intégrée close accepte un canal comme argument et est utilisée pour indiquer qu'aucune autre valeur
ne sera envoyée via le canal. Les récepteurs peuvent vérifier si un canal est fermé lorsqu'ils demandent une valeur
L'emplacement de la flèche spécifie la direction du canal. Lorsque la flèche suit le mot-clé chan (chan<-), le canal ne peut être utilisé que pour envoyer.
Le canal peut être utilisé pour recevoir uniquement si la flèche précède le mot-clé chan (<-chan).
*
*/
func DispatchOrders2(channel chan<- DispatchNotification2) {
	var orderCount int

	rand.NewSource(time.Now().UTC().UnixNano())
	orderCount = rand.Intn(5) + 5
	fmt.Println("Order count : ", orderCount)

	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification2{
			Customer: Customers2[rand.Intn(len(Customers2)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
		time.Sleep(time.Millisecond * 750)
	}
	close(channel)
}
