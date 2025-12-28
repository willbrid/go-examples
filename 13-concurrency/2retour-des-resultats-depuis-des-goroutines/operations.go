package main

import "fmt"

/**
LA fonction TotalPrice prend en paramètre, en plus de la variable catégorie à traiter, une variable (resultChannel) de type chan.
Un résultat est transmis via le canal : le canal est spécifié, suivi d'une flèche de direction (indiquée par les caractères `<` et `-`)
puis de la valeur.
Donc cette instruction `resultChannel <- total` envoie la valeur totale via le canal resultChannel, la rendant ainsi disponible pour être
reçue ailleurs dans l'application.
Lorsqu'une valeur est envoyée via un canal, l'expéditeur n'a pas besoin de savoir comment elle sera reçue et utilisée, tout comme une
fonction synchrone classique ignore comment son résultat sera utilisé.
**/

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
Récupérer le résultat d'une fonction exécutée de manière asynchrone peut s'avérer complexe, car cela nécessite une coordination entre
la goroutine qui produit le résultat et celle qui le consomme. Pour pallier ce problème, Go propose les canaux, qui servent de conduits
pour l'envoi et la réception de données.

Les canaux sont fortement typés : ils ne peuvent transporter que des valeurs d'un `type` ou d'une `interface` spécifique.
Le type d'un canal est défini par le mot-clé `chan`, suivi du type de données qu'il transportera. Les canaux sont créés à l'aide de
la fonction intégrée `make`, en spécifiant leur type.

La syntaxe fléchée est utilisée pour recevoir une valeur d'un canal, ce qui permettra à la fonction CalcStoreTotal de recevoir les
données envoyées par la méthode TotalPrice.
La flèche est placée avant le canal pour recevoir une valeur, et cette valeur reçue peut être utilisée dans n'importe quelle expression
Go standard.
La réception depuis un canal est une opération bloquante, ce qui signifie que l'exécution ne reprendra pas tant qu'une valeur n'aura pas
été reçue.
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
