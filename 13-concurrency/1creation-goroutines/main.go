package main

import (
	"fmt"
	"time"
)

/**
L'élément fondamental pour l'exécution d'un programme Go est la goroutine, un thread léger créé par l'environnement d'exécution Go.
Tous les programmes Go utilisent au moins une goroutine, car c'est ainsi que Go exécute le code de la fonction `main`. Lors de l'exécution
d'un code Go compilé, l'environnement d'exécution crée une goroutine qui commence à exécuter les instructions du point d'entrée,
c'est-à-dire la fonction `main` du package `main`. Chaque instruction de la fonction `main` est exécutée dans l'ordre de sa définition.
La goroutine continue d'exécuter les instructions jusqu'à atteindre la fin de la fonction `main`, moment auquel l'application se termine.

La goroutine exécute chaque instruction de la fonction `main` de manière synchrone, c'est-à-dire qu'elle attend la fin de l'exécution de
chaque instruction avant de passer à la suivante. Les instructions de la fonction `main` peuvent appeler d'autres fonctions, utiliser des
boucles `for`, créer des valeurs, etc. La goroutine `main` parcourt le code, suivant son chemin d'exécution en exécutant une instruction à
la fois.

Tous les produits d'une catégorie soient traités avant de passer à la catégorie suivante.

L'exécution synchrone présente l'avantage d'être simple et cohérente : son comportement est facile à comprendre et prévisible.
Son inconvénient réside dans son inefficacité potentielle. Le traitement séquentiel de neuf éléments de données, comme dans l'exemple,
ne pose aucun problème, mais la plupart des projets réels gèrent des volumes de données plus importants ou doivent effectuer d'autres
tâches, ce qui rend l'exécution séquentielle trop lente et peu performante.

Dans le cas3, l'exécution de la goroutine principale sera suspendue, ce qui permettra aux goroutines créées d'exécuter la méthode TotalPrice.
Une fois la période de suspension écoulée, la goroutine principale reprendra son exécution, atteindra la fin de la fonction et le programme
se terminera.
**/

func main() {
	fmt.Println("main function started")
	// Cas1: Exécution sans goroutine
	// CalcStoreTotal(Products)

	// Cas2: Exécution avec goroutine
	// CalcStoreTotal1Async(Products)

	// Cas3: Exécution avec goroutine plus temporisation au niveau de la fonction main
	CalcStoreTotal1Async(Products)
	time.Sleep(time.Second * 5)
	fmt.Println("main function complete")
}
