package main

import (
	"fmt"
	"time"
)

/**
Le bloc de construction clé pour l'exécution d'un programme Go est la goroutine, qui est un thread léger créé par le runtime Go.
Tous les programmes Go utilisent au moins une goroutine car c'est ainsi que Go exécute le code dans la fonction main.
Lorsque le code Go compilé est exécuté, le runtime crée une goroutine qui commence à exécuter les instructions dans le point d'entrée,
qui est la fonction main du package main. Chaque instruction de la fonction main est exécutée dans l'ordre dans lequel elle est définie.
La goroutine continue d'exécuter des instructions jusqu'à ce qu'elle atteigne la fin de la fonction main, moment auquel l'application se termine.
La goroutine exécute chaque instruction de la fonction main de manière synchrone, ce qui signifie qu'elle attend que l'instruction se
termine avant de passer à l'instruction suivante.
**/

func main() {
	fmt.Println("main function started")
	fmt.Println("Day : ", time.Now().Day(), "-", time.Now().Month(), "-", time.Now().Year())

	// Appel à la méthode CalcStoreTotal1 sans temporisation
	fmt.Println("Méthode CalcStoreTotal1")
	CalcStoreTotal1(Products)

	// Appel à la méthode CalcStoreTotal2 avec temporisation
	fmt.Println("Méthode CalcStoreTotal2")
	CalcStoreTotal2(Products)
	time.Sleep(time.Second * 5) // On retarde la goroutine principale pour permettre aux autres goroutines de terminer leur exécution

	// Appel à la méthode CalcStoreTotal3 sans temporisation
	fmt.Println("Méthode CalcStoreTotal3")
	CalcStoreTotal3(Products)

	// Appel à la méthode CalcStoreTotal3 sans temporisation
	fmt.Println("Méthode CalcStoreTotal4")
	CalcStoreTotal4(Products)

	// Envoi et réception d'un nombre inconnu de valeurs
	/**
	L'opérateur de réception peut être utilisé pour obtenir deux valeurs.
	La première valeur est affectée à la valeur reçue du canal et la deuxième valeur indique si le canal est fermé.
	**/
	var dispatchChannel chan DispatchNotification = make(chan DispatchNotification)
	var details DispatchNotification
	var open bool

	go DispatchOrders(dispatchChannel)
	for {
		// On utilise cette instruction parce que la fonction close est appelée au niveau de la méthode DispatchOrders
		if details, open = <-dispatchChannel; open {
			fmt.Println("Dispatch to ", details.Customer, " : ", details.Quantity, " x ", details.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}

	/**
	Une boucle for peut être utilisée avec le mot-clé range pour énumérer les valeurs envoyées via un canal,
	permettant aux valeurs d'être reçues plus facilement et mettant fin à la boucle lorsque le canal est fermé
	**/
	for details = range dispatchChannel {
		fmt.Println("Dispatch to ", details.Customer, " : ", details.Quantity, " x ", details.Product.Name)
	}
	fmt.Println("Channel has been closed")

	fmt.Println("main function complete")
}
