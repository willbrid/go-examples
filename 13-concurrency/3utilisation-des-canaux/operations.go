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
