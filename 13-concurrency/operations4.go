package main

import (
	"fmt"
	"time"
)

/** Utilisation d'un canal avec tampon et avec les actions d'envoie et de reception de valeur depuis ce canal
Cela fait de l'envoi d'un message une opération non bloquante, permettant à un expéditeur de transmettre sa valeur au canal
et de continuer à travailler sans avoir à attendre un destinataire.
Dans les projets réels, un tampon plus grand (taille = 100) est utilisé, choisi de manière à ce qu'il y ait une capacité
suffisante pour que les goroutines envoient des messages sans avoir à attendre.
**/

func CalcStoreTotal4(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2) // Creation d'un canal avec tampon où envoyé et recevoir les valeurs de type float64
	var value float64

	for category, group := range data {
		go group.TotalPrice4(category, channel)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("-- Starting to receive from channel")
	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending ", len(channel), " items in buffer, size ", cap(channel))
		value = <-channel
		fmt.Println("-- channel read complete", value)
		storeTotal += value
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice4(category string, resultChannel chan float64) {
	var total float64

	for _, p := range group {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println(category, "channel sending", ToCurrency(total))
	resultChannel <- total // La flèche est placée devant une variable pour envoyer une valeur.
	fmt.Println(category, "channel send complete")
}
