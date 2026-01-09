package main

import (
	"context"
	"sync"
	"time"
)

/**
Go facilite la création d'applications serveur qui reçoivent des requêtes pour le compte des clients et les traitent dans leur propre
goroutine. Le package context fournit l'interface `Context`, qui simplifie la gestion des requêtes à l'aide des méthodes :
- Value(key) : cette méthode renvoie la valeur associée à la clé spécifiée.
- Done() : cette méthode renvoie un canal qui peut être utilisé pour recevoir une notification d'annulation.
- Deadline() : cette méthode renvoie le time.Time qui représente le délai de la requête et une valeur booléenne qui sera fausse si aucun délai
			   n'a été spécifié.
- Err() : cette méthode renvoie une erreur qui indique pourquoi le canal Done a reçu un signal. Le package `context` définit deux variables
          qui peuvent être utilisées pour comparer l'erreur : `Canceled` indique que la demande a été annulée et `DeadlineExeeded` indique
		  que le délai est passé.

Le package `context` fournit les fonctions de création de valeurs `context` :
- Background() : cette méthode renvoie le contexte par défaut, à partir duquel d'autres contextes sont dérivés.
- WithCancel(ctx) : cette méthode renvoie un contexte et une fonction d'annulation.
- WithDeadline(ctx, time) : cette méthode renvoie un contexte avec une échéance, qui est exprimée à l'aide d'une valeur time.Time.
- WithTimeout(ctx, duration) : cette méthode renvoie un contexte avec une échéance, qui est exprimée à l'aide d'une valeur time.Duration.
- WithValue(ctx, key, val) : cette méthode renvoie un contexte contenant la paire clé-valeur spécifiée.

La fonction `processRequest` simule le traitement d'une requête en incrémentant un compteur, et appelle la fonction `time.Sleep`
pour ralentir l'exécution. La fonction principale utilise une goroutine pour invoquer `processRequest`, en se substituant à une requête
provenant d'un client.

La première utilité d'un contexte est d'informer le code traitant la requête lorsque celle-ci est annulée.
**/

func processRequest(wg *sync.WaitGroup, count int) {
	total := 0
	for i := range count {
		_ = i + 1
		Printfln("Processing request : %v", total)
		total++
		time.Sleep(time.Millisecond * 250)
	}
	Printfln("Request processed...%v", total)
	wg.Done()
}

func processRequestWithCTX(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := range count {
		_ = i + 1
		select {
		case <-ctx.Done():
			Printfln("Stopping processing - request cancelled")
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

/**
Il est possible de créer des contextes avec une date limite, après laquelle un signal est envoyé sur le canal `Done`, comme lors
de l'annulation d'une requête. Une durée absolue peut être spécifiée à l'aide de la fonction `WithDeadline`, qui accepte une
valeur `time.Time`. La fonction `WithTimeout` accepte une durée `time.Duration`, qui spécifie une date limite relative à l'heure actuelle.
La méthode `Context.Deadline` permet de vérifier la date limite pendant le traitement d'une requête.
**/

func processRequestWithCTXDeadline(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := range count {
		_ = i + 1
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

/**
La fonction WithValue crée un contexte dérivé avec une paire clé-valeur qui peut être lue lors du traitement de la requête.
**/

const (
	countKey = iota
	sleepPeriodKey
)

func processRequestWithCTXValue(ctx context.Context, wg *sync.WaitGroup) {
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
	for i := range count {
		_ = i + 1
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(sleepPeriod)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	go processRequest(&waitGroup, 10)
	waitGroup.Wait()

	/**
	La fonction Background renvoie le contexte par défaut, qui ne fait rien d'utile mais fournit un point de départ pour dériver de
	nouvelles valeurs de contexte avec les autres fonctions.

	La fonction `WithCancel` renvoie un contexte annulable et la fonction appelée pour effectuer l'annulation.
	Le contexte ainsi obtenu est transmis à la fonction processRequest. La fonction principale appelle la fonction `time.Sleep` pour laisser
	à processRequestWithCTX le temps d'effectuer certaines opérations, puis appelle la fonction d'annulation.
	L'appel de la fonction d'annulation envoie un message au canal renvoyé par la méthode `Done` du contexte, lequel est surveillé par
	une instruction `select`.
	Le canal `Done` est bloqué si la requête n'a pas été annulée ; la clause `default` est alors exécutée, permettant ainsi le traitement de
	la requête. Le canal est vérifié après chaque unité de travail, et une instruction `goto` est utilisée pour sortir de la boucle de
	traitement afin que le `WaitGroup` puisse être signalé et que la fonction se termine.
	**/
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(1)
	Printfln("Request dispatched...")
	ctx1, cancel := context.WithCancel(context.Background())
	go processRequestWithCTX(ctx1, &waitGroup1, 10)
	time.Sleep(time.Second)
	Printfln("Canceling request")
	cancel()
	waitGroup1.Wait()

	/**
	Les fonctions WithDeadline et WithTimeout renvoient le contexte dérivé et une fonction d'annulation, permettant d'annuler la requête
	avant l'expiration du délai. Dans cet exemple, le temps d'exécution de la fonction processRequest dépasse le délai imparti,
	ce qui entraîne l'arrêt du traitement par le canal Done.
	**/
	waitGroup2 := sync.WaitGroup{}
	waitGroup2.Add(1)
	Printfln("Request dispatched...")
	ctx2, _ := context.WithTimeout(context.Background(), time.Second*2)
	go processRequestWithCTXDeadline(ctx2, &waitGroup2, 10)
	waitGroup2.Wait()

	/**
	La fonction `WithValue` crée un contexte dérivé avec une paire clé-valeur qui peut être lue lors du traitement de la requête.
	La fonction WithValue n'accepte qu'une seule paire clé-valeur, mais les fonctions du package `context` peuvent être appelées de manière
	répétée pour créer la combinaison de caractéristiques requise. Dans l'exemple ci-dessous, la fonction WithTimeout est utilisée pour
	créer un contexte avec une échéance, et ce contexte est ensuite utilisé comme argument de la fonction WithValue afin d'y ajouter
	deux paires clé-valeur. Ces données sont accessibles via la méthode Value, ce qui signifie que les fonctions de traitement des
	requêtes n'ont pas besoin de définir de paramètres pour toutes les valeurs de données dont elles ont besoin.
	**/
	waitGroup3 := sync.WaitGroup{}
	waitGroup3.Add(1)
	Printfln("Request dispatched...")
	ctx3, _ := context.WithTimeout(context.Background(), time.Second*2)
	ctx3 = context.WithValue(ctx3, countKey, 4)
	ctx3 = context.WithValue(ctx3, sleepPeriodKey, time.Millisecond*250)
	go processRequestWithCTXValue(ctx3, &waitGroup3)
	waitGroup3.Wait()
}
