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


Réception de notifications récurrentes
La fonction `Tick` renvoie un canal par lequel les valeurs de type `Time` sont envoyées à un intervalle spécifié.
L'utilité du canal créé par la fonction `Tick` ne réside pas dans les valeurs de type `Time` qui y transitent, mais dans leur périodicité d'envoi.
Dans cet exemple, la fonction `Tick` sert à créer un canal qui envoie des valeurs chaque seconde. Le canal se bloque lorsqu'aucune valeur n'est
disponible, ce qui permet au canal créé avec la fonction `Tick` de contrôler la fréquence de génération des valeurs par la fonction writeToChannel.

La fonction `Tick` est utile lorsqu'une séquence indéfinie de signaux est requise. Si une série fixe de valeurs est requise, alors
on utilise la fonction ci-dessous :

NewTicker(duration) : Cette fonction renvoie un *Ticker avec la période spécifiée. Le résultat de la fonction `NewTicker` est un pointeur vers
une structure Ticker, qui définit le champ et les méthodes :
- C : Ce champ indique le canal par lequel le `Ticker` transmettra ses valeurs de type `Time`.
- Stop() : Cette méthode arrête le `Ticker` (sans fermer le canal indiqué par le champ `C`).
- Reset(duration) : Cette méthode arrête le `Ticker` et le réinitialise afin que son intervalle corresponde à la durée spécifiée.
**/

func writeToChannel(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	tickChannel := time.Tick(time.Second)
	index := 0

	for {
		<-tickChannel
		nameChannel <- names[index]
		index++
		if index == len(names) {
			index = 0
		}
	}
}

/**
Cette approche est utile lorsqu'une application doit créer plusieurs flux de données sans laisser ceux qui ne sont plus nécessaires
envoyer des messages.
**/

func writeToChannel1(nameChannel chan<- string) {
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	ticker := time.NewTicker(time.Second / 10)
	index := 0
	for {
		<-ticker.C
		nameChannel <- names[index]
		index++
		if index == len(names) {
			ticker.Stop()
			close(nameChannel)
			break
		}
	}
}

func main() {
	nameChannel := make(chan string)
	// go writeToChannel(nameChannel)
	go writeToChannel1(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}
}
