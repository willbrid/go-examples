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


Reporter l'exécution d'une fonction
La fonction `AfterFunc` permet de reporter l'exécution d'une fonction pendant une période spécifiée.

Dans l'exemple ci-dessous, le premier argument de `AfterFunc` est le délai, qui est de cinq secondes. Le deuxième argument est la fonction à exécuter.
Ici, la fonction writeToChannel, mais `AfterFunc` n'accepte que les fonctions sans paramètres ni résultats; c'est pourquoi une fonction intermédiaire.
**/

func writeToChannel(channel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
	}
	close(channel)
}

func main() {
	nameChannel := make(chan string)
	time.AfterFunc(time.Second*5, func() {
		writeToChannel(nameChannel)
	})
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
