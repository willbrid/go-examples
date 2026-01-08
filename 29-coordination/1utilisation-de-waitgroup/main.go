package main

import "sync"

func doSum(count int, val *int) {
	for i := range count {
		i++
		*val++
	}
}

var waitGroup sync.WaitGroup = sync.WaitGroup{}

func doSumWG(count int, val *int) {
	for i := range count {
		i++
		*val++
	}
	waitGroup.Done()
}

/**
Il est important de ne pas copier les valeurs `WaitGroup` car cela signifie que les goroutines appelleront Done et Wait sur des
valeurs différentes, ce qui signifie généralement que l'application se bloque. Si nous voulons passer un `WaitGroup` comme argument de
fonction, cela signifie que nous devons utiliser un pointeur.
**/

func doSumWG2(count int, val *int, wg *sync.WaitGroup) {
	for i := range count {
		i++
		*val++
	}
	wg.Done()
}

func main() {
	counter := 0
	doSum(5000, &counter)
	Printfln("Total : %v", counter)

	/**
	Un problème courant consiste à s'assurer que la fonction principale ne se termine pas avant que les goroutines qu'elle lance ne soient
	terminées, auquel cas le programme s'arrête.
	Les goroutines sont si simples à créer qu'il est facile d'oublier leur impact. Dans ce cas précis, l'exécution de la fonction principale
	se poursuit en parallèle avec la goroutine, ce qui signifie que la dernière instruction de la fonction principale est exécutée
	avant que la goroutine n'ait terminé l'exécution de la fonction `doSum`.
	**/
	counter1 := 0
	go doSum(5000, &counter1)
	Printfln("Total1 : %v", counter1)

	/**
	Le package `sync` fournit la structure `WaitGroup`, qui peut être utilisée pour attendre la fin d'une ou plusieurs goroutines.
	Les méthodes définies par la structure `WaitGroup` :
	- Add(num) : cette méthode augmente le nombre de goroutines que le `WaitGroup` attend de l'entier spécifié.
	- Done() : cette méthode diminue de un le nombre de goroutines que le `WaitGroup` attend.
	- Wait() : cette méthode bloque jusqu'à ce que la méthode Done ait été appelée une fois pour le nombre total de goroutines spécifié
	    par les appels à la méthode `Add`.

	Le `WaitGroup` agit comme un compteur. Lorsque les goroutines sont créées, la méthode `Add` est appelée pour spécifier le nombre de goroutines
	qui sont démarrées, ce qui incrémente le compteur, après quoi la méthode `Wait` est appelée, qui bloque. Au fur et à mesure que chaque
	goroutine se termine, elle appelle la méthode Done, qui décrémente le compteur. Lorsque le compteur est à zéro, la méthode `Wait`
	arrête le blocage, terminant le processus d'attente.

	Le `WaitGroup` paniquera si le compteur devient négatif, il est donc important d'appeler la méthode `Add` avant de démarrer la goroutine
	pour éviter que la méthode `Done` ne soit appelée plus tôt. Il est également important de s'assurer que le total des valeurs transmises à
	la méthode `Add` est égal au nombre d'appels de la méthode Done. S'il y a trop peu d'appels à `Done`, la méthode `Wait` sera bloquée pour
	toujours, mais si la méthode `Done` est appelée trop de fois, alors `WaitGroup` paniquera.
	**/
	counter2 := 0
	waitGroup.Add(1)
	go doSumWG(5000, &counter2)
	waitGroup.Wait()
	Printfln("Total2 : %v", counter2)

	/**
	Si nous voulons passer un `WaitGroup` comme argument de fonction, cela signifie que nous devons utiliser un pointeur.
	**/
	waitGroup2 := sync.WaitGroup{}
	counter3 := 0
	waitGroup2.Add(1)
	go doSumWG2(5000, &counter3, &waitGroup2)
	waitGroup2.Wait()
	Printfln("Total3 : %v", counter3)
}
