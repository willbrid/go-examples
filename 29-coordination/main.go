package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

/**
Le package de synchronisation fournit la classe WaitGroup, qui peut être utilisée pour attendre qu'une ou plusieurs goroutines se terminent.
Les méthodes définies par la classe WaitGroup :
- Add(num) : cette méthode augmente le nombre de goroutines que le WaitGroup attend de l'entier spécifié.
- Done() : cette méthode diminue de un le nombre de goroutines que le WaitGroup attend.
- Wait() : cette méthode bloque jusqu'à ce que la méthode Done ait été appelée une fois pour le nombre total de goroutines spécifié par les appels
           à la méthode Add.

Le WaitGroup agit comme un compteur. Lorsque les goroutines sont créées, la méthode Add est appelée pour spécifier le nombre de goroutines
qui sont démarrées, ce qui incrémente le compteur, après quoi la méthode Wait est appelée, qui bloque. Au fur et à mesure que chaque goroutine se termine,
elle appelle la méthode Done, qui décrémente le compteur. Lorsque le compteur est à zéro, la méthode Wait arrête le blocage, terminant le processus d'attente.

Le WaitGroup paniquera si le compteur devient négatif, il est donc important d'appeler la méthode Add avant de démarrer la goroutine pour éviter que
la méthode Done ne soit appelée plus tôt. Il est également important de s'assurer que le total des valeurs transmises à la méthode Add est égal au
nombre d'appels de la méthode Done. S'il y a trop peu d'appels à Done, la méthode Wait sera bloquée pour toujours, mais si la méthode Done est appelée
trop de fois, alors WaitGroup paniquera.

Il est important de ne pas copier les valeurs WaitGroup car cela signifie que les goroutines appelleront Done et Wait sur des valeurs différentes,
ce qui signifie généralement que l'application se bloque. Si nous voulons passer un WaitGroup comme argument de fonction, cela signifie que
nous devons utiliser un pointeur.

Les méthodes définies par la classe Mutex :
- Lock() : cette méthode verrouille le mutex. Si le Mutex est déjà verrouillé, cette méthode bloque jusqu'à ce qu'il soit déverrouillé.
- Unlock() : cette méthode déverrouille le Mutex.

Un Mutex traite toutes les goroutines comme étant égales et n'autorise qu'une seule goroutine à acquérir le verrou. La classe RWMutex est plus flexible
et prend en charge deux catégories de goroutines : les lecteurs et les rédacteurs. N'importe quel nombre de lecteurs peut acquérir le verrou simultanément,
ou un seul rédacteur peut acquérir le verrou. L'idée est que les lecteurs ne se soucient que des conflits avec les rédacteur et peuvent exécuter
en même temps que d'autres lecteurs sans difficulté.
Les méthodes définies par la classe RWMutex :
- RLock() : cette méthode tente d'acquérir le verrou de lecture et bloquera jusqu'à ce qu'il soit acquis.
- RUnlock() : cette méthode libère le verrou de lecture.
- Lock() : cette méthode tente d'acquérir le verrou en écriture et bloquera jusqu'à ce qu'il soit acquis.
- Unlock() : cette méthode libère le verrou en écriture.
- RLocker() : cette méthode renvoie un pointeur vers un Locker pour acquérir et libérer le verrou de lecture

Voici les règles suivies par le RWMutex :
• Si le RWMutex est déverrouillé, alors le verrou peut être acquis par un lecteur (en appelant le
RLock) ou un rédacteur (en appelant la méthode Lock).
• Si le verrou est acquis par un lecteur, d'autres lecteurs peuvent également acquérir le verrou en
appelant la méthode RLock, qui ne bloquera pas. La méthode Lock bloquera jusqu'à ce que tous
des lecteurs libèrent le verrou en appelant la méthode RUnlock.
• Si le verrou est acquis par un rédacteur, alors les deux méthodes RLock et Lock bloqueront pour empêcher d'autres goroutines d'acquérir le verrou
jusqu'à ce que la méthode Unlock soit appelée.
• Si le verrou est acquis par un lecteur et qu'un rédacteur appelle la méthode Lock, les méthodes Lock et RLock se bloqueront jusqu'à ce que
la méthode Unlock soit appelée. Cela empêche le mutex d'être perpétuellement verrouillé par les lecteurs sans donner aux rédacteurs une chance d'acquérir
le verrou en écriture.
**/

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}
var rwmutex = sync.RWMutex{}

var squares = map[int]int{}

func doSum(count int, val *int) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

func doSumWithWaitGroup(count int, val *int, waitGroup *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		*val++
	}
	waitGroup.Done()
}

/*
*
Un Mutex est déverrouillé lors de sa création, ce qui signifie que la première goroutine qui appelle la méthode Lock ne bloquera pas et pourra incrémenter
la variable compteur. On dit que la goroutine a acquis la serrure. Toute autre goroutine qui appelle la méthode Lock se bloquera jusqu'à ce que
la méthode Unlock soit appelée, connue sous le nom de libération du verrou, moment auquel une autre goroutine pourra acquérir le verrou et procéder
à son accès à la variable compteur.

La meilleure approche pour utiliser l'exclusion mutuelle est d'être prudent et conservateur. Nous devons nous assurer que tout le code qui
accède aux données partagées le fait en utilisant le même Mutex, et chaque appel à une méthode Lock doit être équilibré par un appel à la méthode Unlock.
Il peut être tentant d'essayer de créer des améliorations ou des optimisations intelligentes, mais cela peut entraîner des performances médiocres ou
des blocages d'applications.
*
*/
func doSumWithMutualExclusion(count int, val *int, waitGroup *sync.WaitGroup) {
	time.Sleep(time.Second)
	mutex.Lock()
	for i := 0; i < count; i++ {
		*val++
	}
	mutex.Unlock()
	waitGroup.Done()
}

/*
*
La fonction calculateSquares acquiert le verrou de lecture pour vérifier si une map contient une clé choisie au hasard. Si la map contient la clé,
la valeur associée est lue et le verrou de lecture est libéré. Si la map ne contient pas la clé, le verrou en écriture est acquis, une valeur est
ajoutée à la map pour la clé, puis le verrou en écriture est relâché.
L'utilisation du RWMutex signifie que lorsqu'une goroutine a le verrou de lecture, d'autres routines peuvent également acquérir le verrou et effectuer
des lectures. La lecture des données ne pose aucun problème de simultanéité à moins qu'elles ne soient modifiées en même temps. Si une goroutine appelle
la méthode Lock, elle ne pourra pas acquérir le verrou en écriture tant que le verrou en lecture n'aura pas été relâché par toutes les goroutines qui
l'ont acquis.
*
*/
func calculateSquares(max, iterations int, waitGroup *sync.WaitGroup) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Cached value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value : %v = %v", val, squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	counter := 0
	waitGroup.Add(1)
	go doSum(5000, &counter)
	waitGroup.Wait()
	Printfln("Total : %v", counter)

	counter1 := 0
	waitGroup1 := sync.WaitGroup{}
	waitGroup1.Add(1)
	go doSumWithWaitGroup(5000, &counter1, &waitGroup1)
	waitGroup1.Wait()
	Printfln("Total : %v", counter1)

	// Exclusion mutuelle
	counter2 := 0
	waitGroup2 := sync.WaitGroup{}
	numRoutines2 := 3
	waitGroup2.Add(numRoutines2)
	for i := 0; i < numRoutines2; i++ {
		go doSumWithMutualExclusion(5000, &counter2, &waitGroup2)
	}
	waitGroup2.Wait()
	Printfln("Total : %v", counter2)

	// Utilisation d'un mutex en lecture et écriture
	numRoutines3 := 3
	waitGroup3 := sync.WaitGroup{}
	waitGroup3.Add(numRoutines3)
	for i := 0; i < numRoutines3; i++ {
		go calculateSquares(10, 5, &waitGroup3)
	}
	waitGroup3.Wait()
	Printfln("Cached values : %v", len(squares))
}
