package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

/**
Si plusieurs goroutines accèdent aux mêmes données, il est possible que deux goroutines accèdent à ces données simultanément et
provoquent des résultats inattendus.
Une solution à ce problème consiste à utiliser l'exclusion mutuelle, qui garantit à une goroutine un accès exclusif aux données dont elle a
besoin et empêche les autres goroutines d'y accéder. L'exclusion mutuelle est comparable à l'emprunt d'un livre à la bibliothèque :
une seule personne peut l'emprunter à la fois, et les autres doivent attendre qu'elle ait terminé.

Le package `sync` fournit l'exclusion mutuelle grâce à la structure `Mutex`.
Les méthodes définies par la structure `Mutex` :
- Lock() : cette méthode verrouille le `mutex`. Si le `Mutex` est déjà verrouillé, cette méthode bloque jusqu'à ce qu'il soit déverrouillé.
- Unlock() : cette méthode déverrouille le `Mutex`.

Un mutex est déverrouillé lors de sa création, ce qui signifie que la première goroutine appelant la méthode `Lock` ne sera pas bloquée et
pourra incrémenter la variable `counter`. On dit alors que la goroutine a acquis le verrou. Toute autre goroutine appelant la méthode `Lock`
sera bloquée jusqu'à l'appel de la méthode `Unlock`, ce qui libère le verrou. Une fois cette libération effectuée, une autre goroutine pourra
acquérir le verrou et accéder à la variable `counter`.
**/

var mutex sync.Mutex = sync.Mutex{}

func doSum(count int, val *int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	for i := range count {
		mutex.Lock()
		i++
		*val++
		mutex.Unlock()
	}
	wg.Done()
}

/**
L'utilisation de l'exclusion mutuelle requiert une grande prudence, et il est important d'en analyser les conséquences.
Dans l'exemple de la fonction `doSum` ci-dessus, le mutex est verrouillé et déverrouillé à chaque incrémentation de la variable.
L'utilisation d'un mutex a un impact, et une alternative consiste à le verrouiller avant l'exécution de la boucle `for`.
Il est préférable de commencer par verrouiller uniquement les instructions accédant aux données partagées.
**/

func doSumImproved(count int, val *int, mt *sync.Mutex, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	mt.Lock()
	for i := range count {
		i++
		*val++
	}
	mt.Unlock()
	wg.Done()
}

/**
Un Mutex traite toutes les goroutines comme étant égales et n'autorise qu'une seule goroutine à acquérir le verrou. La structure `RWMutex`
est plus flexible et prend en charge deux catégories de goroutines : les lecteurs et les rédacteurs. N'importe quel nombre de lecteurs
peut acquérir le verrou simultanément, ou un seul rédacteur peut acquérir le verrou. L'idée est que les lecteurs ne se soucient que des
conflits avec les rédacteurs et peuvent exécuter en même temps que d'autres lecteurs sans difficulté.
Les méthodes définies par la structure `RWMutex` :
- RLock() : cette méthode tente d'acquérir le verrou de lecture et bloquera jusqu'à ce qu'il soit acquis.
- RUnlock() : cette méthode libère le verrou de lecture.
- Lock() : cette méthode tente d'acquérir le verrou en écriture et bloquera jusqu'à ce qu'il soit acquis.
- Unlock() : cette méthode libère le verrou en écriture.
- RLocker() : cette méthode renvoie un pointeur vers un Locker pour acquérir et libérer le verrou de lecture

Voici les règles suivies par le `RWMutex`:
• Si le `RWMutex` est déverrouillé, alors le verrou peut être acquis par un lecteur (en appelant le
RLock) ou un rédacteur (en appelant la méthode `Lock`).
• Si le verrou est acquis par un lecteur, d'autres lecteurs peuvent également acquérir le verrou en
appelant la méthode `RLock`, qui ne bloquera pas. La méthode `Lock` bloquera jusqu'à ce que tous
des lecteurs libèrent le verrou en appelant la méthode `RUnlock`.
• Si le verrou est acquis par un rédacteur, alors les deux méthodes `RLock` et `Lock` bloqueront pour empêcher d'autres goroutines d'acquérir
le verrou jusqu'à ce que la méthode `Unlock` soit appelée.
• Si le verrou est acquis par un lecteur et qu'un rédacteur appelle la méthode `Lock`, les méthodes `Lock` et `RLock` se bloqueront jusqu'à
ce que la méthode `Unlock` soit appelée. Cela empêche le mutex d'être perpétuellement verrouillé par les lecteurs sans donner aux rédacteurs
une chance d'acquérir le verrou en écriture.

La fonction `calculateSquares` acquiert un verrou de lecture pour vérifier si une table contient une clé choisie aléatoirement. Si la table
contient la clé, la valeur associée est lue et le verrou de lecture est libéré. ​​Si la table ne contient pas la clé, un verrou d'écriture
est acquis, une valeur est ajoutée à la table pour la clé, puis le verrou d'écriture est libéré.

L'utilisation de `RWMutex` permet, lorsqu'une goroutine détient le verrou de lecture, d'autres routines peuvent également l'acquérir et
effectuer des lectures. La lecture de données ne pose aucun problème de concurrence, sauf si elles sont modifiées simultanément. Si une
goroutine appelle la méthode `Lock`, elle ne pourra pas acquérir le verrou d'écriture tant que le verrou de lecture n'aura pas été libéré
par toutes les goroutines qui l'ont acquis.

`RWMutex` ne prend pas en charge la conversion du verrou de lecture en verrou d'écriture. Il est impératif de libérer le verrou de lecture
avant d'appeler la méthode Lock afin d'éviter un blocage. Un délai peut survenir entre la libération du verrou de lecture et l'acquisition
du verrou d'écriture. Durant ce délai, d'autres goroutines peuvent acquérir le verrou d'écriture et effectuer des modifications.
Il est donc crucial de vérifier que l'état des données n'a pas changé une fois le verrou d'écriture acquis.
**/

var rwmutext sync.RWMutex = sync.RWMutex{}
var squares map[int]int = map[int]int{}

func calculateSquares(max, iterations int, wg *sync.WaitGroup) {
	for i := range iterations {
		_ = i + 1
		val := rand.Intn(max)
		rwmutext.RLock()
		square, ok := squares[val]
		rwmutext.RUnlock()
		if ok {
			Printfln("Cached value : %v = %v", val, square)
		} else {
			rwmutext.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value : %v = %v", val, squares[val])
			}
			rwmutext.Unlock()
		}
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup = sync.WaitGroup{}
	counter := 0
	numRountines := 3
	wg.Add(numRountines)
	for i := range numRountines {
		_ = i + 1
		go doSum(1000, &counter, &wg)
	}
	wg.Wait()
	Printfln("Total : %v", counter)

	var wg1 sync.WaitGroup = sync.WaitGroup{}
	var mt sync.Mutex = sync.Mutex{}
	counter1 := 0
	numRountine1s := 3
	wg1.Add(numRountine1s)
	for i := range numRountine1s {
		_ = i + 1
		go doSumImproved(1000, &counter1, &mt, &wg1)
	}
	wg1.Wait()
	Printfln("Total1 : %v", counter1)

	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var wg2 sync.WaitGroup = sync.WaitGroup{}
	numRountine2s := 3
	wg2.Add(numRountine2s)
	for i := range numRountine2s {
		_ = i + 1
		go calculateSquares(10, 5, &wg2)
	}
	wg2.Wait()
	Printfln("Cached values : %v", len(squares))
}
