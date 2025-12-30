package main

import "fmt"

func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("#n Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

func main() {
	fmt.Println("main function started")
	// Cas1 : sans buffer
	// CalcStoreTotal1Async(Products)
	// Cas2 : avec buffer
	// CalcStoreTotal2Async(Products)
	// Cas3 : avec buffer et calcul de la longueur et la capacité
	// CalcStoreTotal3Async(Products)

	/**
	L'important est que la goroutine se termine après avoir envoyé ses valeurs, laissant la goroutine principale bloquée en attente
	d'une nouvelle valeur. L'environnement d'exécution Go détecte l'absence de goroutines actives et met fin à l'application.
	La solution à ce problème consiste pour l'émetteur à indiquer, en fermant le canal, qu'aucune nouvelle valeur n'est reçue par celui-ci.

	L'opérateur de réception permet d'obtenir deux valeurs. La première correspond à la valeur reçue du canal, et la seconde indique
	si le canal est fermé.
	Si le canal est ouvert, l'indicateur de fermeture sera faux et la valeur reçue du canal sera affectée à l'autre variable. Si le canal
	est fermé, l'indicateur de fermeture sera vrai et la valeur zéro correspondant au type de canal sera affectée à l'autre variable.
	**/
	dispatchChannel := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel)
	for {
		if details, open := <-dispatchChannel; open {
			fmt.Println("#1 Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
		} else {
			fmt.Println("Channel has been closed")
			break
		}
	}

	/**
	Une boucle for peut être utilisée avec le mot-clé `range` pour énumérer les valeurs transmises via un canal, facilitant ainsi
	leur réception et interrompant la boucle lorsque le canal est fermé.
	L'expression `range` produit une valeur par itération, correspondant à la valeur reçue du canal. La boucle « for » continuera de recevoir
	des valeurs jusqu'à la fermeture du canal. (Nous pouvons utiliser une boucle « `for...range` » sur un canal non fermé ;
	dans ce cas, la boucle ne s'arrêtera jamais.)
	**/
	dispatchChannel1 := make(chan DispatchNotification, 100)
	go DispatchOrders(dispatchChannel1)
	for details := range dispatchChannel1 {
		fmt.Println("#2 Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}

	dispatchChannel2 := make(chan DispatchNotification, 100)
	go DispatchOrder1s(dispatchChannel2)
	for details := range dispatchChannel2 {
		fmt.Println("#3 Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}

	/**
	Go permet d'assigner des canaux bidirectionnels à des variables de canaux unidirectionnels, ce qui permet d'appliquer des restrictions.
	L'on utilise la syntaxe complète des variables pour définir des variables de canal d'envoi uniquement et de réception uniquement,
	qui sont ensuite utilisées comme arguments de fonction. Cela garantit que le destinataire du canal d'envoi uniquement peut uniquement
	envoyer des valeurs ou fermer le canal, et que le destinataire du canal de réception uniquement peut uniquement recevoir des valeurs.
	Ces restrictions s'appliquent au même canal sous-jacent, de sorte que les messages envoyés via `sendOnlyChannel` seront reçus
	via `receiveOnlyChannel`.
	**/
	dispatchChannel3 := make(chan DispatchNotification, 100)
	var sendOnlyChannel chan<- DispatchNotification = dispatchChannel3
	var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel3
	go DispatchOrder1s(sendOnlyChannel)
	receiveDispatches(receiveOnlyChannel)

	/**
	Des restrictions sur la direction du canal peuvent également être créées par conversion explicite.
	**/
	dispatchChannel4 := make(chan DispatchNotification, 100)
	go DispatchOrder1s((chan<- DispatchNotification)(dispatchChannel4))
	receiveDispatches((<-chan DispatchNotification)(dispatchChannel4))

	fmt.Println("main function complete")
}
