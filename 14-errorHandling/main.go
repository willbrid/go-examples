package main

import "fmt"

// Gérer une panique dans les goroutines
type CategoryCountMessage struct {
	Category      string
	Count         int
	TerminalError interface{}
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() { // définition d'une fonction anonyme
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessage{ // Envoie de l'erreur dans le canal pour récupération dans la fonction main
				TerminalError: arg,
			}
			close(outChan)
		}
	}()
	var channel chan ChannelMessage = make(chan ChannelMessage, 10)
	go products.TotalPriceAsync(categories, channel)
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

func main() {
	/**------------- Traiter les erreurs récupérables -----------------**/
	var categories []string = []string{"Watersports", "Chess", "Running"}
	var total float64
	var err *CategoryError

	for _, cat := range categories {
		total, err = products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, " Total : ", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}

	// Reception d'erreur via un canal
	var channel chan ChannelMessage = make(chan ChannelMessage, 10)
	go products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}

	// Utilisation de l'intégration du package errors
	var channel1 chan ChannelMessage1 = make(chan ChannelMessage1, 10)
	go products.TotalPriceAsync1(categories, channel1)
	for message := range channel1 {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}

	// Utilisation de l'intégration du package errors
	var channel2 chan ChannelMessage2 = make(chan ChannelMessage2, 10)
	go products.TotalPriceAsync2(categories, channel2)
	for message := range channel2 {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}

	/**------------- Traiter les erreurs non récupérables -----------------**/
	/**
		La fonction panic est invoquée avec un argument, qui peut être n'importe quelle valeur permettant d'expliquer la panique.
		Lorsque la fonction panic est appelée, l'exécution de la fonction englobante est interrompue et toutes les fonctions de report (defer) sont exécutées.

		Go fournit la fonction intégrée recover, qui peut être appelée pour empêcher une panique de remonter la pile des appels et de mettre fin au programme.
		La fonction recover doit être appelée dans le code exécuté à l'aide du mot-clé defer.

		L'appel de la fonction recover renvoie une valeur s'il y a eu une panique (panic), stoppant la progression de la panique et
		donnant accès à l'argument utilisé pour invoquer la fonction de panique.
	    Étant donné que n'importe quelle valeur peut être transmise à la fonction panique, le type de la valeur renvoyée par la fonction de récupération
		est l'interface vide (interface{}), qui nécessite une assertion de type avant de pouvoir être utilisée.
		La fonction recover traite les types d'erreur et de chaîne, qui sont les deux types d'arguments de panique les plus courants.
		**/
	var recoveryFunc func() = func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error : ", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message : ", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}
	defer recoveryFunc()

	var channel3 chan ChannelMessage = make(chan ChannelMessage, 10)
	go products.TotalPriceAsync(categories, channel3)
	for message := range channel3 {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			panic(message.CategoryError)
		}
	}

	/**
	La fonction main utilise une goroutine pour invoquer la fonction processCategories, qui panique si la fonction TotalPriceAsync envoie une erreur.
	Le processCategories se remet de la panique, mais cela a une conséquence inattendue. Le problème est que la récupération après une panique ne reprend pas
	l'exécution de la fonction processCategories, ce qui signifie que la fonction close n'est jamais appelée sur le canal à partir duquel
	la fonction main reçoit des messages. La fonction main essaie de recevoir un message qui ne sera jamais envoyé et se bloque sur le canal,
	déclenchant la détection de blocage du runtime Go : la solution la plus simple consiste à appeler la fonction close sur le canal lors de
	la récupération (fonction appelant la fonction recover()). Aussi nous pouvons passer dans le canal le message d'échec qui sera utilisé par la fonction main.
	Le résultat est que la décision sur la façon de gérer la panique est transmise de la goroutine au code appelant, qui peut choisir de continuer
	l'exécution ou de déclencher une nouvelle panique en fonction du problème.
	**/
	var channel4 chan CategoryCountMessage = make(chan CategoryCountMessage)
	go processCategories(categories, channel4)
	for message := range channel4 {
		if message.TerminalError == nil {
			fmt.Println(message.Category, " Total : ", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}
}
