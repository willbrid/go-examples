package main

import "fmt"

/**
Par défaut, l'envoi et la réception via un canal sont des opérations bloquantes. Cela signifie qu'une goroutine qui envoie une valeur
ne pourra exécuter aucune autre instruction tant qu'une autre goroutine n'aura pas reçu cette valeur du canal. Si une seconde goroutine
envoie une valeur, elle sera bloquée jusqu'à ce que le canal soit libéré, ce qui créera une file d'attente de goroutines attendant
la réception de valeurs. Ce comportement est également réversible : les goroutines qui reçoivent des valeurs seront bloquées jusqu'à
ce qu'une autre goroutine en envoie une.

Aucun récepteur n'étant disponible, les goroutines sont contraintes d'attendre, formant une file d'attente d'expéditeurs jusqu'à ce que
le récepteur prenne le relais. Dès réception d'une valeur, la goroutine émettrice est débloquée et peut poursuivre l'exécution des
instructions de la méthode TotalPrice.
**/

func CalcStoreTotal1Async(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		storeTotal += <-channel
	}
	fmt.Println("Total :", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		fmt.Println(category, "product :", p.Name)
		total += p.Price
	}
	fmt.Println(category, "subtotal :", ToCurrency(total))
	resultChannel <- total
}

/**
Le comportement par défaut d'un canal peut entraîner des pics d'activité lorsque les goroutines s'exécutent, suivis d'une longue période
d'inactivité en attente de la réception de messages. Dans un projet réel, les goroutines effectuent souvent des tâches répétitives, et
l'attente d'un destinataire peut engendrer un goulot d'étranglement au niveau des performances.

Une autre approche consiste à créer un canal avec une mémoire tampon, qui reçoit les valeurs d'un expéditeur et les stocke jusqu'à ce
qu'un destinataire soit disponible. L'envoi d'un message devient ainsi une opération non bloquante, permettant à l'expéditeur de
transmettre sa valeur au canal et de poursuivre son travail sans attendre de destinataire. On peut comparer cela à la boîte de réception
d'Alice sur son bureau. Les expéditeurs viennent au bureau d'Alice et y déposent leur message, qu'elle lira quand elle sera disponible.
Cependant, si la boîte de réception est pleine, ils devront attendre qu'elle ait traité une partie des messages en attente avant
d'en envoyer un nouveau.

Dans l'exemple ci-dessous, la taille du tampon a été configuré à 2, ce qui signifie que deux expéditeurs pourront envoyer des valeurs via
le canal sans avoir à attendre leur réception. Tout expéditeur suivant devra attendre la réception d'un des messages mis en mémoire tampon.

Dans les projets réels, on utilise un tampon plus grand, choisi de manière à ce que les goroutines aient une capacité suffisante
pour envoyer des messages sans avoir à attendre. (L'on peut spécifier généralement une taille de tampon de 100,
ce qui est généralement suffisant pour la plupart des projets sans pour autant nécessiter une quantité de mémoire importante.)
**/

func CalcStoreTotal2Async(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2) // Création d'un canal avec buffer
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		storeTotal += <-channel
	}
	fmt.Println("Total :", ToCurrency(storeTotal))
}

/**
L'on peut déterminer la taille du tampon d'un canal à l'aide de la fonction intégrée `cap` et déterminer le nombre de valeurs contenues
dans le tampon à l'aide de la fonction `len`.
**/

func CalcStoreTotal3Async(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2) // Création d'un canal avec buffer
	for category, group := range data {
		go group.TotalPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending", len(channel), "items in buffer, size", cap(channel))
		storeTotal += <-channel
	}
	fmt.Println("Total :", ToCurrency(storeTotal))
}
