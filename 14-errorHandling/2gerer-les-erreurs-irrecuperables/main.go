package main

import "fmt"

/**
Certaines erreurs sont si graves qu'elles devraient entraîner l'arrêt immédiat de l'application, un processus connu sous le nom de panique.
**/

func main() {
	/**
	Go fournit la fonction intégrée `recover`, qui permet d'empêcher une panique de remonter la pile d'appels et de mettre fin au programme.
	La fonction `recover` doit être appelée dans du code exécuté avec le mot-clé `defer`.

	Cet exemple ci-dessous utilise le mot-clé `defer` pour enregistrer une fonction qui sera exécutée une fois la fonction principale
	terminée, même en l'absence d'erreur critique. L'appel à la fonction `recover` renvoie une valeur en cas d'erreur critique,
	interrompant ainsi sa progression et permettant d'accéder à l'argument utilisé pour appeler la fonction de gestion de l'erreur critique.
	**/
	/**
	// cas1
	recoveryFunc := func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error :", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message :", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}
	defer recoveryFunc() **/

	/**
	Puisque n'importe quelle valeur peut être passée à la fonction `panic`, le type de la valeur renvoyée par la fonction `recover` est
	l'interface vide `interface{}`, ce qui nécessite une assertion de type avant utilisation. La fonction `recover` de l'exemple ci-dessous
	gère les types `error` et `string`, qui sont les deux types d'arguments les plus courants pour `panic`.
	Il peut être maladroit de définir une fonction et de l'utiliser immédiatement avec le mot-clé `defer`, c'est pourquoi la récupération
	après une erreur `panic` est généralement effectuée à l'aide d'une fonction anonyme.
	Notons l'utilisation des parenthèses après l'accolade fermante de la fonction anonyme ; elles sont nécessaires pour appeler et
	non simplement définir la fonction anonyme.
	**/
	/**
	// # cas1
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error :", err.Error())
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message :", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()**/

	/**
	Il est possible de se remettre d'une crise pour ensuite se rendre compte que la situation est finalement irrémédiable. Dans ce cas,
	nous pouveons déclencher une nouvelle instruction `panic`, soit en fournissant un nouvel argument, soit en réutilisant la valeur reçue
	lors de l'appel à la fonction `recover`.
	**/
	/**
	// cas2
	defer func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error :", err.Error())
				panic(err)
			} else if str, ok := arg.(string); ok {
				fmt.Println("Message :", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}()**/

	/**
	Au lieu d'afficher un message lorsqu'une catégorie est introuvable, la fonction principale déclenche une panique, ce qui est réalisé
	à l'aide de la fonction intégrée `panic`.

	La fonction `panic` est appelée avec un argument, qui peut prendre n'importe quelle valeur permettant d'expliquer l'erreur.
	Dans l'exemple ci-dessous, elle est appelée avec une erreur, ce qui constitue une manière efficace de combiner les fonctionnalités
	de gestion des erreurs de Go.
	Lorsque la fonction `panic` est appelée, l'exécution de la fonction englobante est interrompue et les fonctions `defer` sont exécutées.
	L'erreur se propage dans la pile d'appels, interrompant l'exécution des fonctions appelantes et invoquant leurs fonctions `defer`.
	**/
	/**
	// cas1 et cas2
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			panic(message.CategoryError)
		}
	}**/

	categorie1s := []string{"Watersports", "Chess", "Running"}
	channel1 := make(chan CategoryCountMessage)
	go processCategories(categorie1s, channel1)
	for message := range channel1 {
		fmt.Println(message.Category, "Total :", message.Count)
	}

	categorie2s := []string{"Watersports", "Chess", "Running"}
	channel2 := make(chan CategoryCountResult)
	go processXCategories(categorie2s, channel2)
	for message := range channel2 {
		if message.TerminalError == nil {
			fmt.Println(message.Category, "Total :", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}
}
