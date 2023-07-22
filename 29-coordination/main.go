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

Le package sync fournit la fonction ci-après pour créer des valeurs de classe Cond.
- NewCond(*locker) : cette fonction crée un Cond en utilisant le pointeur vers le Locker spécifié.
L'argument de la fonction NewCond est un Locker, qui est une interface qui définit les méthodes :
--- Lock() : cette méthode acquiert le verrou géré par le Locker.
--- Unlock() : cette méthode libère le verrou géré par le Locker.
Les classes Mutex et RWMutex définissent la méthode requise par l'interface Locker. Dans le cas du RWMutex, les méthodes Lock et Unlock fonctionnent sur
le verrou en écriture, et la méthode RLocker peut être utilisée pour obtenir un Locker qui fonctionne sur le verrou en lecture.
Le champ et les méthodes définis par la classe Cond :
- L : ce champ renvoie le Locker qui a été transmis à la fonction NewCond et qui est utilisé pour acquérir le verrou.
- Wait() : cette méthode libère le verrou et suspend la goroutine.
- Signal() : cette méthode réveille une goroutine en attente.
- Broadcast() : cette méthode réveille toutes les goroutines en attente.
**/

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}
var rwmutex = sync.RWMutex{}
var readyCond = sync.NewCond(rwmutex.RLocker())
var once = sync.Once{}

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

/*
*
Cet exemple nécessite une coordination entre les goroutines qui serait difficile à réaliser sans Cond. Une goroutine est chargée de remplir une map
avec des valeurs de données, qui sont ensuite lues par d'autres goroutines. Les lecteurs doivent être avertis que la génération des données est terminée
avant de s'exécuter. Les lecteurs attendent en acquérant le verrou Cond et en appelant la méthode Wait.
L'appel de la méthode Wait suspend la goroutine et libère le verrou afin qu'il puisse être acquis. L'appel à la méthode Wait est généralement
effectué à l'intérieur d'une boucle for qui vérifie que la condition pour laquelle la goroutine attend s'est produite, juste pour s'assurer que
les données sont dans l'état attendu. Il n'est pas nécessaire d'acquérir à nouveau le verrou lorsque la méthode Wait se débloque et une goroutine
peut soit appeler à nouveau la méthode Wait, soit accéder aux données partagées. Lorsque nous avons terminé avec les données partagées, le verrou doit
être libéré.
La goroutine qui génère les données acquiert le verrou en écriture à l'aide du RWMutex, modifie les données, libère le verrou en écriture, puis appelle
la méthode Cond.Broadcast, qui réveille toutes les goroutines en attente.
L'appel à la fonction time.Sleep dans la fonction readSquares ralentit le processus de lecture des données de sorte que les deux goroutines de lecteur
traitent les données en même temps, ce que nous pouvons voir dans l'entrelacement du premier nombre dans les lignes de sortie. Étant donné que
ces goroutines acquièrent un verrou de lecture RWMutex, les deux acquièrent le verrou et peuvent lire les données simultanément.

La classe Once (sync.Once) définit une méthode :
- Do(func) : cette méthode exécute la fonction spécifiée, mais seulement si elle n'a pas déjà été exécutée.
*
*/
func generateSquares(max int, waitGroup *sync.WaitGroup) {
	rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroup.Done()
}

func readSquares(id, max, iterations int, waitGroup *sync.WaitGroup) {
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value : %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	waitGroup.Done()
}

/*
*
L'utilisation de la classe Once simplifie l'exemple précédente car la méthode Do se bloque jusqu'à ce que la fonction qu'elle reçoit ait été exécutée,
après quoi elle revient sans exécuter à nouveau la fonction. Étant donné que les seules modifications apportées aux données partagées dans cet exemple
sont apportées par la fonction generateSquaresByOnce, l'utilisation de la méthode Do pour exécuter cette fonction garantit que les modifications
sont apportées en toute sécurité.
*
*/
func generateSquaresByOnce(max int) {
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
}

func readSquaresWithOnce(id, max, iterations int, waitGroup *sync.WaitGroup) {
	once.Do(func() {
		generateSquaresByOnce(max)
	})
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value : %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
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

	// Utilisation conditionnelle de go routine
	numRoutines4 := 2
	waitGroup4 := sync.WaitGroup{}
	waitGroup4.Add(numRoutines4)
	for i := 0; i < numRoutines4; i++ {
		go readSquares(i, 10, 5, &waitGroup4)
	}
	waitGroup4.Add(1)
	go generateSquares(10, &waitGroup4)
	waitGroup4.Wait()
	Printfln("Cached values : %v", len(squares))

	// Exécuter une fonction une fois
	numRoutines5 := 2
	waitGroup5 := sync.WaitGroup{}
	waitGroup5.Add(numRoutines5)
	for i := 0; i < numRoutines5; i++ {
		go readSquaresWithOnce(i, 10, 5, &waitGroup5)
	}
	waitGroup5.Wait()
	Printfln("Cached values : %v", len(squares))
}
