package main

import "fmt"

/**
Une panique remonte la pile d'appels jusqu'au sommet de la goroutine courante, où elle provoque l'arrêt de l'application.
Cette restriction implique que les paniques doivent être gérées au sein du code exécuté par la goroutine.
**/

type CategoryCountMessage struct {
	Category string
	Count    int
}

type CategoryCountResult struct {
	Category      string
	Count         int
	TerminalError any
}

/**
La fonction `main` utilise une goroutine pour appeler la fonction `processCategories`, qui provoque une panique si la fonction
`TotalPriceAsync` renvoie une erreur. `processCategories` se remet de la panique, mais cela a une conséquence inattendue.
Le problème est que la récupération après une panique ne reprend pas l'exécution de `processCategories`, ce qui signifie que
la fonction `close` n'est jamais appelée sur le canal par lequel la fonction `main` reçoit des messages. La fonction `main` tente de
recevoir un message qui ne sera jamais envoyé et se bloque sur le canal, déclenchant ainsi la détection d'interblocage du runtime Go.
L'approche la plus simple consiste à appeler la fonction `close` sur le canal pendant la récupération.
**/

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			close(outChan) // Correction du problème de deadlock
		}
	}()

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}

/**
Une meilleur approche permet de signaler les erreurs via la canal.
Cela évite le blocage, mais sans indiquer à la fonction principale que la fonction `processCategories` ci-dessus n'a pas pu terminer
son exécution, ce qui peut avoir des conséquences. Une meilleure approche consiste à signaler ce résultat via le canal avant de le fermer.

Il en résulte que la décision concernant la gestion de la panique est transmise de la goroutine au code appelant, qui peut choisir de
poursuivre l'exécution ou de déclencher une nouvelle panique en fonction du problème.
**/

func processXCategories(categories []string, outChan chan<- CategoryCountResult) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountResult{
				TerminalError: arg,
			}
			close(outChan) // Correction du problème de deadlock
		}
	}()

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountResult{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}
