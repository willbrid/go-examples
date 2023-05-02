package main

import (
	"fmt"
	"time"
)

// Utilisation d'un canal avec l'action d'envoie et de reception de valeur depuis ce canal
/**
Renvoi des résultats des goroutines
Go fournit des canaux, qui sont des conduits par lesquels les données peuvent être envoyées et reçues.
Les canaux sont fortement typés, ce qui signifie qu'ils transporteront des valeurs d'un type ou d'une interface spécifiés.
Le type d'un canal est le mot-clé chan, suivi du type que le canal portera.
Pour envoyer une valeur dans un canel, le canal est spécifié, suivi d'une flèche de direction exprimée par les caractères < et - puis par la valeur.
La flèche est placée devant le canal pour en recevoir une valeur.
Lorsqu'une valeur est envoyée via un canal, l'expéditeur n'a pas besoin de savoir comment la valeur sera reçue et utilisée.
**/

func CalcStoreTotal3(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64) // Creation d'un canal où envoyé et recevoir les valeurs de type float64
	var value float64

	for category, group := range data {
		go group.TotalPrice3(category, channel)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("-- Starting to receive from channel")
	for i := 0; i < len(data); i++ {
		fmt.Println("-- channel read pending")
		value = <-channel
		fmt.Println("-- channel read complete", value)
		storeTotal += value
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice3(category string, resultChannel chan float64) {
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
