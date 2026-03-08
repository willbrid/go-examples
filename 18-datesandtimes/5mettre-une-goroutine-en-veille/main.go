package main

import "time"

/**
Utilisation des fonctionnalités temporelles pour les goroutines et les canaux
Le package `time` fournit un petit ensemble de fonctions utiles pour travailler avec les goroutines et les canaux.

Sleep(duration) : Cette fonction suspend l’exécution de la goroutine courante pendant au moins la durée spécifiée.

AfterFunc(duration, fonction) : Cette fonction exécute la fonction spécifiée dans sa propre goroutine après la durée spécifiée.
Le résultat est un *Timer, dont la méthode Stop permet d’annuler l’exécution de la fonction avant la fin de la durée.

After(duration) : Cette fonction renvoie un canal qui se bloque pendant la durée spécifiée, puis produit une valeur Time.

Tick(duration) : Cette fonction renvoie un canal qui envoie périodiquement une valeur Time, la période étant spécifiée par la durée.

Mettre une goroutine en veille
La fonction `Sleep` suspend l'exécution de la goroutine en cours pendant une durée spécifiée.
**/

/**
La durée spécifiée par la fonction `Sleep` correspond à la durée minimale de suspension de la goroutine. Il est déconseillé de se fier à des durées
précises, surtout pour les courtes durées. Notons que la fonction `Sleep` suspend la goroutine dans laquelle elle est appelée, et donc également
la goroutine principale, ce qui peut donner l'impression d'un blocage de l'application.
(Dans ce cas, l'absence d'alerte de la détection automatique de blocage indique que nous avons appelé la fonction `Sleep` par erreur.)
**/

func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
