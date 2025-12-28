package main

import "fmt"

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		storeTotal += group.TotalPrice(category)
	}
	fmt.Println("Total :", ToCurrency(storeTotal))
}

func (group ProductGroup) TotalPrice(category string) (total float64) {
	for _, p := range group {
		fmt.Println(category, "product :", p.Name)
		total += p.Price
	}
	fmt.Println(category, "subtotal :", ToCurrency(total))
	return
}

/**
Go permet au développeur de créer des goroutines supplémentaires, qui exécutent du code simultanément à la goroutine principale (main).
Go facilite la création de nouvelles goroutines.

Une goroutine est créée à l'aide du mot-clé `go` suivi de la 'fonction' ou de la 'méthode' qui doit être exécutée de manière asynchrone.

Lorsque l'environnement d'exécution Go rencontre le mot-clé `go`, il crée une nouvelle goroutine et l'utilise pour exécuter la fonction
ou la méthode spécifiée.

Cela modifie l'exécution du programme car, à tout instant, plusieurs goroutines sont actives, chacune exécutant son propre ensemble
d'instructions. Ces instructions sont exécutées simultanément.

Dans l'exemple ci-dessous, une goroutine est créée pour chaque appel à la méthode `TotalPrice`, ce qui signifie que les catégories
sont traitées simultanément.
Cette instruction indique au runtime d'exécuter les instructions de la méthode `TotalPrice` à l'aide d'une nouvelle goroutine.
Le runtime n'attend pas la fin de l'exécution de la méthode par la goroutine et passe immédiatement à l'instruction suivante.
C'est le principe même des goroutines : la méthode `TotalPrice` est appelée de manière asynchrone, ce qui signifie que ses instructions
sont évaluées par une goroutine en même temps que la goroutine d'origine exécute les instructions de la fonction principale.
Cependant, le programme s'arrête lorsque la goroutine principale a exécuté toutes les instructions de la fonction principale.
Par conséquent, le programme s'arrête avant que les goroutines nécessaires à l'exécution complète de la méthode `TotalPrice` ne soient créées,
ce qui explique l'absence d'affichage de `subtotal`.
**/

func CalcStoreTotal1Async(data ProductData) {
	var storeTotal float64
	for category, group := range data {
		go group.TotalPrice(category)
	}
	fmt.Println("Total :", ToCurrency(storeTotal))
}
