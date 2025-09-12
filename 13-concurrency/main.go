package main

import (
	"fmt"
	"time"
)

/**
- Pour exécuter une fonction de manière asynchrone, créer une goroutine

- Pour produire un résultat à partir d'une fonction exécutée de manière asynchrone, utiliser un canal

- Pour envoyer et recevoir des valeurs via un canal, utiliser des expressions fléchées

- Pour indiquer qu'aucune autre valeur ne sera envoyée, utiliser la fonction close sur un canal

- Pour énumérer les valeurs reçues d'un canal, utiliser une boucle for avec le mot-clé range

- Pour envoyer ou recevoir des valeurs via plusieurs canaux, utiliser une instruction select
**/

/*
*
Fonction permettant de recevoir les valeurs depuis un canal
*
*/
func receiveDispatches(channel <-chan DispatchNotification1) {
	var details DispatchNotification1

	for details = range channel {
		fmt.Println("Dispatch to ", details.Customer, " : ", details.Quantity, " x ", details.Product.Name)
	}
	fmt.Println("Channel has been closed")

}

func enumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}

// Envoyé dans un canal sans blocage
func enumerateProducts1(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		select {
		case channel <- p:
			fmt.Println("Sent product:", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}

// Envoyé vers plusieurs canaux sans blocage
func enumerateProducts2(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList[:3] {
		select {
		case channel1 <- p:
			fmt.Println("Send via channel 1 : ", p.Name)
		case channel2 <- p:
			fmt.Println("Send via channel 2 : ", p.Name)
		default:
			fmt.Println("Discarding product:", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel1)
	close(channel2)
}

/**
Le bloc de construction clé pour l'exécution d'un programme Go est la goroutine, qui est un thread léger créé par le runtime Go.
Tous les programmes Go utilisent au moins une goroutine car c'est ainsi que Go exécute le code dans la fonction main.
Lorsque le code Go compilé est exécuté, le runtime crée une goroutine qui commence à exécuter les instructions dans le point d'entrée,
qui est la fonction main du package main. Chaque instruction de la fonction main est exécutée dans l'ordre dans lequel elle est définie.
La goroutine continue d'exécuter des instructions jusqu'à ce qu'elle atteigne la fin de la fonction main, moment auquel l'application se termine.
La goroutine exécute chaque instruction de la fonction main de manière synchrone, ce qui signifie qu'elle attend que l'instruction se
termine avant de passer à l'instruction suivante.

Une goroutine est créée à l'aide du mot-clé "go" suivi de la fonction ou de la méthode à exécuter de manière asynchrone.
Lorsque l'environnement d'exécution Go rencontre le mot-clé go, il crée une nouvelle goroutine et l'utilise pour exécuter
la fonction ou la méthode spécifiée. Cela modifie l'exécution du programme, car à tout moment, plusieurs goroutines sont présentes,
chacune exécutant son propre ensemble d'instructions. Ces instructions sont exécutées simultanément, ce qui signifie simplement
qu'elles sont exécutées en même temps.
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

	// Forcer la direction d'un canal : le canal dispatchChannel1 est utilisé pour envoyer des valeurs
	var dispatchChannel1 chan DispatchNotification1 = make(chan DispatchNotification1)
	var details1 DispatchNotification1

	go DispatchOrders1(dispatchChannel1)
	for details1 = range dispatchChannel1 {
		fmt.Println("Dispatch to ", details1.Customer, " : ", details1.Quantity, " x ", details1.Product.Name)
	}
	fmt.Println("Channel has been closed")

	var dispatchChannel2 chan DispatchNotification1 = make(chan DispatchNotification1, 100)
	var sendOnlyChannel chan<- DispatchNotification1 = dispatchChannel2
	var receiveOnlyChannel <-chan DispatchNotification1 = dispatchChannel2
	go DispatchOrders1(sendOnlyChannel)
	receiveDispatches(receiveOnlyChannel)

	/**
	Utilisation de l'instruction select
	L'instruction select évalue ses instructions case une seule fois, c'est pourquoi on peut utiliser également utilisé une boucle for.
	La boucle continue d'exécuter l'instruction select, qui recevra les valeurs du canal dès qu'elles seront disponibles.
	Si aucune valeur n'est disponible, la clause default est exécutée.

	L'utilisation la plus simple des instructions select est de recevoir d'un canal sans bloquer,
	garantissant qu'une goroutine n'aura pas à attendre lorsque le canal est vide.
	Une instruction select a une structure similaire à une instruction switch, sauf que les instructions case sont des opérations de canal.
	Lorsque l'instruction select est exécutée, chaque opération de canal est évaluée jusqu'à ce qu'une opération pouvant être effectuée sans
	blocage soit atteinte. L'opération de canal est effectuée et les instructions incluses dans l'instruction case sont exécutées.

	l'instruction select est utilisée pour recevoir des valeurs de deux canaux, l'un qui porte les valeurs DispatchNofitication2 et l'autre qui
	porte les valeurs Product. Chaque fois que l'instruction select est exécutée, elle se fraye un chemin à travers les instructions case,
	constituant une liste de celles à partir desquelles une valeur peut être lue sans blocage. L'une des déclarations de cas est sélectionnée au hasard
	dans la liste et exécutée. Si aucune des instructions case ne peut être exécutée, la clause par défaut est exécutée. Des précautions doivent être
	prises pour gérer les canaux fermés car ils fourniront une valeur nulle pour chaque opération de réception qui se produit après la fermeture du canal,
	en s'appuyant sur l'indicateur fermé pour montrer que le canal est fermé. Malheureusement, cela signifie que les instructions case pour les canaux
	fermés seront toujours choisies par les instructions select car elles sont toujours prêtes à fournir une valeur sans blocage, même si cette valeur
	n'est pas utile.
	**/
	var dispatchChannel3 chan DispatchNotification2 = make(chan DispatchNotification2, 100)
	var productChannel chan *Product = make(chan *Product)
	go DispatchOrders2(dispatchChannel3)
	go enumerateProducts(productChannel)
	var openChannels int = 2

	for {
		select {
		case details2, ok := <-dispatchChannel3:
			if ok {
				fmt.Println("Dispatch to ", details2.Customer, ":", details2.Quantity, "x", details2.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				dispatchChannel3 = nil
				openChannels--
			}
		case product, ok := <-productChannel:
			if ok {
				fmt.Println("Product : ", product.Name)
			} else {
				fmt.Println("Product channel has been closed")
				productChannel = nil
				openChannels--
			}
		default:
			if openChannels == 0 {
				goto alldone
			}
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	alldone:
		fmt.Println("All values received")
		break
	}

	/**
	Envoyé dans un canal sans blocage avec l'instruction select.
	la fonction enumerateProducts peut envoyer des valeurs via le canal sans bloquer jusqu'à ce que la mémoire tampon soit pleine.
	La clause par défaut de l'instruction select ignore les valeurs qui ne peuvent pas être envoyées.
	**/
	var productChannel1 chan *Product = make(chan *Product, 5)
	go enumerateProducts1(productChannel1)
	time.Sleep(time.Second)
	for p := range productChannel1 {
		fmt.Println("Received product : ", p.Name)
	}

	/**
	Envoyé vers plusieurs canaux sans blocage
	S'il y a plusieurs canaux disponibles, une instruction select peut être utilisée pour trouver un canal pour lequel l'envoi ne bloquera pas.
	**/
	var c1 chan *Product = make(chan *Product, 2)
	var c2 chan *Product = make(chan *Product, 2)
	go enumerateProducts2(c1, c2)
	time.Sleep(time.Second)
	for p := range c1 {
		fmt.Println("Channel 1 received product : ", p.Name)
	}
	for p := range c2 {
		fmt.Println("Channel 2 received product : ", p.Name)
	}

	fmt.Println("main function complete")
}
