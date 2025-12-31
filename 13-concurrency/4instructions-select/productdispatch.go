package main

import (
	"fmt"
	"time"
)

func EnumerateProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		channel <- p
		time.Sleep(time.Millisecond * 800)
	}
	close(channel)
}

/**
Une instruction `SELECT` peut également être utilisée pour envoyer des données à un canal sans bloquer le flux.
La clause `default` de l'instruction `select` ignore les valeurs qui ne peuvent pas être envoyées.
**/

func EnumerateWithSelectProducts(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		select {
		case channel <- p:
			fmt.Println("#3 Sent product :", p.Name) // Cette ligne peut être commentée ou enlevée (exemple ci-dessous)
		default:
			fmt.Println("#3 Discarding product :", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}

/**
Dans l'exemple ci-dessus, l'instruction `CASE` contient une instruction qui écrit un message ; cependant, cela n'est pas obligatoire et
l'instruction `CASE` peut spécifier des `opérations d'envoi sans instructions supplémentaires`.
**/

func EnumerateWithSelectProduct1s(channel chan<- *Product) {
	for _, p := range ProductList[:3] {
		select {
		case channel <- p:
		default:
			fmt.Println("#3 Discarding product :", p.Name)
			time.Sleep(time.Second)
		}
	}
	close(channel)
}

/**
S'il existe plusieurs canaux disponibles, une instruction `select` permet de trouver celui sur lequel l'envoi ne sera pas bloquant.
**/

func EnumerateWithSelectCasesProducts(channel1, channel2 chan<- *Product) {
	for _, p := range ProductList {
		select {
		case channel1 <- p:
			fmt.Println("#4 Send via channel 1")
		case channel2 <- p:
			fmt.Println("#4 Send via channel 2")
		}
	}
	close(channel1)
	close(channel2)
}
