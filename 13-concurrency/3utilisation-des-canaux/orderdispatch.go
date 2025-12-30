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
Envoie de plusieurs valeurs dans le canal.
La fonction DispatchOrders crée un nombre aléatoire de valeurs DispatchNotification et les envoie via le canal reçu par le paramètre channel.
Il est impossible de savoir à l'avance combien de valeurs DispatchNotification la fonction DispatchOrders va créer, ce qui représente
un défi lors de l'écriture du code qui reçoit les données du canal.

La fonction intégrée `close` accepte un canal comme argument et sert à indiquer qu'aucune autre valeur ne sera envoyée par ce canal.
Les récepteurs peuvent vérifier si un canal est fermé lors d'une requête de valeur.
**/

func DispatchOrders(channel chan DispatchNotification) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	orderCount := r.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := range orderCount {
		fmt.Println("i :", i)
		channel <- DispatchNotification{
			Customer: Customers[r.Intn(len(Customers)-1)],
			Quantity: r.Intn(10),
			Product:  ProductList[r.Intn(len(ProductList)-1)],
		}
	}
	close(channel) // On ferme le canal, pour indiquer qu'aucune nouvelle valeur ne sera envoyée.
}

/**
Par défaut, les canaux permettent d'envoyer et de recevoir des données, mais cette fonctionnalité peut être restreinte lorsqu'ils sont
utilisés comme arguments, limitant ainsi les opérations à l'envoi ou à la réception. Cette fonctionnalité est utile pour éviter les erreurs
de réception.
La direction du canal est spécifiée à côté du mot-clé `chan` : `channel chan<- DispatchNotification`.
La position de la flèche indique le sens du canal. Lorsque la flèche suit le mot-clé `chan` : `chan<-`, le canal ne peut être utilisé que
pour l'envoi (comme une affectation). Le canal ne peut être utilisé que pour la réception si la flèche précède le mot-clé `chan` (`<-chan`).
Tenter de recevoir des données d'un canal en mode envoi uniquement (et vice versa) provoque une erreur de compilation.
**/

func DispatchOrder1s(channel chan<- DispatchNotification) {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	orderCount := r.Intn(3) + 2
	fmt.Println("Order count:", orderCount)
	for i := range orderCount {
		fmt.Println("i :", i)
		channel <- DispatchNotification{
			Customer: Customers[r.Intn(len(Customers)-1)],
			Quantity: r.Intn(10),
			Product:  ProductList[r.Intn(len(ProductList)-1)],
		}
	}
	close(channel)
}
