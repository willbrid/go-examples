package main

import "fmt"

/**
Le programme démarre et exécute les instructions de la fonction principale. Cela entraîne l'appel de la fonction CalcStoreTotal, qui crée
un canal et lance plusieurs goroutines. Ces goroutines exécutent les instructions de la méthode TotalPrice, qui transmet son résultat via
le canal.
La goroutine principale poursuit l'exécution des instructions de la fonction CalcStoreTotal, qui reçoit les résultats par le canal.
Ces résultats servent à calculer un total, qui est ensuite affiché. Enfin, les instructions restantes de la fonction principale sont
exécutées et le programme se termine.
**/

func main() {
	fmt.Println("main function started")
	CalcStoreTotal1Async(Products)
	fmt.Println("main function complete")
}
