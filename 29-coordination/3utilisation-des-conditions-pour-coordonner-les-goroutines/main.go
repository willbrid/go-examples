package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

/**
package sync
Lorsque les goroutines nécessitent une coordination, par exemple pour attendre un événement, la structure `Cond` peut être utilisée.
Le package sync fournit la fonction décrite ci-dessous pour créer des valeurs de structure `Cond`.
- NewCond(*locker) : cette fonction crée un `Cond` en utilisant le pointeur vers le Locker spécifié.

L'argument de la fonction `NewCond` est un `Locker`, qui est une interface qui définit les méthodes :
--- Lock() : cette méthode acquiert le verrou géré par le `Locker`.
--- Unlock() : cette méthode libère le verrou géré par le `Locker`.

Les structures `Mutex` et `RWMutex` définissent la méthode requise par l'interface `Locker`.
Dans le cas du `RWMutex`, les méthodes `Lock` et `Unlock` fonctionnent sur le verrou en écriture, et la méthode `RLocker` peut être utilisée
pour obtenir un `Locker` qui fonctionne sur le verrou en lecture.
Le champ et les méthodes définis par la structure `Cond` :
- L : ce champ renvoie le Locker qui a été transmis à la fonction `NewCond` et qui est utilisé pour acquérir le verrou.
- Wait() : cette méthode libère le verrou et suspend la goroutine.
- Signal() : cette méthode réveille une goroutine en attente.
- Broadcast() : cette méthode réveille toutes les goroutines en attente.
**/

var waitGroup sync.WaitGroup = sync.WaitGroup{}
var rwmutex sync.RWMutex = sync.RWMutex{}
var readyCond *sync.Cond = sync.NewCond(rwmutex.RLocker())

var squares map[int]int = map[int]int{}

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Generating data...")
	for val := range max {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroup.Done()
}

func readSquares(id, max, iterations int) {
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	for i := range iterations {
		_ = i + 1
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	waitGroup.Done()
}

func main() {
	/**
	Cet exemple nécessite une coordination entre goroutines qui serait difficile à réaliser sans un verrou `Cond`. Une goroutine est chargée
	de remplir une map avec des valeurs de données, lesquelles sont ensuite lues par d'autres goroutines. Ces dernières doivent être
	notifiées de la fin de la génération des données avant de s'exécuter. Elles attendent en acquérant le verrou `Cond` et en appelant
	la méthode `Wait`.

	L'appel à la méthode Wait suspend la goroutine et libère le verrou afin qu'il puisse être acquis. Cet appel est généralement effectué
	dans une boucle for qui vérifie que la condition d'attente de la goroutine est remplie, afin de s'assurer que les données sont dans
	l'état attendu.
	Il n'est pas nécessaire d'acquérir à nouveau le verrou lorsque la méthode Wait débloque la goroutine ; celle-ci peut alors soit appeler
	à nouveau la méthode Wait, soit accéder aux données partagées. Une fois l'accès aux données partagées terminé, le verrou doit être libéré.
	La goroutine qui génère les données acquiert le verrou d'écriture à l'aide de RWMutex, modifie les données, libère le verrou d'écriture,
	puis appelle la méthode Cond.Broadcast, ce qui réveille toutes les goroutines en attente.
	L'appel à la fonction `time.Sleep` dans la fonction `readSquares` ralentit la lecture des données afin que les deux goroutines de lecture
	traitent les données simultanément, ce qui se traduit par l'entrelacement du premier nombre sur les lignes de sortie. Ces goroutines
	acquérant un verrou de lecture RWMutex, elles peuvent toutes deux lire les données simultanément.

	var readyCond = sync.NewCond(&rwmutex) : ce changement signifie que toutes les goroutines utilisent le verrou d'écriture, ce qui
	signifie qu'une seule goroutine pourra acquérir ce verrou.
	**/

	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := range numRoutines {
		go readSquares(i, 10, 5)
	}

	waitGroup.Add(1)
	go generateSquares(10)
	waitGroup.Wait()
	Printfln("Cached values : %v", len(squares))
}
