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

func DispatchOrders(channel chan<- DispatchNotification) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	orderCount := r.Intn(5) + 5
	fmt.Println("Order count:", orderCount)
	for i := range orderCount {
		fmt.Println("i :", i)
		channel <- DispatchNotification{
			Customer: Customers[r.Intn(len(Customers)-1)],
			Quantity: r.Intn(10),
			Product:  ProductList[r.Intn(len(ProductList)-1)],
		}
		time.Sleep(time.Millisecond * 750)
	}
	close(channel)
}
