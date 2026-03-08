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


Réception de notifications programmées
La fonction `After` attend une durée spécifiée, puis envoie une valeur Time à un canal. C'est un moyen pratique d'utiliser un canal pour recevoir
une notification à une heure ultérieure donnée.

La fonction `After` génère un canal véhiculant des valeurs de type `Time`. Ce canal se bloque pendant la durée spécifiée lorsqu'une
valeur de type `Time` est envoyée, indiquant ainsi que la durée est écoulée. Dans cet exemple ci-dessous, la valeur transmise sur le canal sert
de signal et n'est pas utilisée directement ; c'est pourquoi elle est assignée à l'identificateur vide.
==writeToChannel==
L'effet dans cet exemple ci-dessous est le même qu'avec la fonction `Sleep`, mais la différence est que la fonction `After` renvoie un canal qui ne
se bloque pas tant qu'une valeur n'est pas lue, ce qui signifie qu'une direction peut être spécifiée, un travail supplémentaire peut être effectué,
puis une lecture du canal peut être effectuée, avec pour résultat que le canal ne sera bloqué que pour la partie restante de la durée.


Utilisation des notifications comme délais d'expiration dans les instructions `select`
La fonction `After` peut être utilisée avec les instructions `select` pour définir un délai d'expiration.
==writeToChannel1==
L'instruction `select` bloque l'exécution jusqu'à ce qu'un des canaux soit prêt ou jusqu'à l'expiration du délai. Cela fonctionne car
l'instruction `select` bloque l'exécution jusqu'à ce qu'un de ses canaux soit prêt et car la fonction `After` crée un canal qui bloque pendant
une durée spécifiée.


Arrêt et réinitialisation des minuteurs
La fonction `After` est utile si nous sommes certain d'avoir toujours besoin de la notification programmée. Si nous souhaitons annuler la notification,
nous pouvons utiliser cette fonction ci-dessous.
NewTimer(duration) : Cette fonction renvoie un `*Timer` avec la période spécifiée.

Le résultat de la fonction `NewTimer` est un pointeur vers une structure `Timer`, qui définit les méthodes ci-dessous.

CT : ce champ renvoie le canal par lequel le minuteur enverra sa valeur de temps.
Stop() : cette méthode arrête le minuteur. Le résultat est un booléen qui vaut vrai si le minuteur a été arrêté et faux s'il a déjà envoyé son message.
Reset(duration) : cette méthode arrête le minuteur et le réinitialise afin que son intervalle soit égal à la durée spécifiée.

Dans cet exemple, le minuteur est configuré pour une durée de dix minutes. Une goroutine attend deux secondes, puis réinitialise
le minuteur pour que sa durée soit de deux secondes.
**/

func writeToChannel(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 1)
	}
	close(channel)
}

func writeToChannel1(channel chan<- string) {
	Printfln("Waiting for initial duration...")
	_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
		time.Sleep(time.Second * 3)
	}
	close(channel)
}

func writeToChannel2(channel chan<- string) {
	timer := time.NewTimer(time.Minute * 10)

	go func() {
		time.Sleep(time.Second * 2)
		Printfln("Resetting timer")
		timer.Reset(time.Second)
	}()

	Printfln("Waiting for initial duration...")
	<-timer.C
	Printfln("Initial duration elapsed.")
	names := []string{"Alice", "Bob", "Charlie", "Dora"}
	for _, name := range names {
		channel <- name
	}
	close(channel)
}

func main() {
	nameChannel := make(chan string)
	go writeToChannel(nameChannel)
	for name := range nameChannel {
		Printfln("Read name: %v", name)
	}

	nameChannel1 := make(chan string)
	go writeToChannel1(nameChannel1)
	channelOpen := true
	for channelOpen {
		Printfln("Starting channel read")
		select {
		case name, ok := <-nameChannel1:
			if !ok {
				channelOpen = false
			} else {
				Printfln("#1 Read name: %v", name)
			}
		case <-time.After(time.Second * 2):
			Printfln("Timeout")
		}
	}

	nameChannel2 := make(chan string)
	go writeToChannel2(nameChannel2)
	for name := range nameChannel2 {
		Printfln("#2 Read name: %v", name)
	}
}
