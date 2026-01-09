package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

/**
Une autre approche possible par rapport à l'exemple précédent consiste à s'assurer que la fonction `generateSquares` est exécutée une seule fois,
en utilisant la structure `sync.Once`. La structure `Once` définit une méthode
- Do(func) : cette méthode exécute la fonction spécifiée, mais seulement si elle n'a pas déjà été exécutée.

L'utilisation de la structure `Once` simplifie l'exemple précédent, car la méthode `Do` bloque l'exécution jusqu'à ce que la fonction
qu'elle reçoit soit exécutée, puis retourne sans réexécuter la fonction. Étant donné que les seules modifications apportées aux données
partagées dans cet exemple sont effectuées par la fonction `generateSquares`, l'utilisation de la méthode `Do` pour exécuter cette fonction
garantit que les modifications sont effectuées en toute sécurité.
**/

var waitGroup = sync.WaitGroup{}
var once sync.Once = sync.Once{}

var squares map[int]int = map[int]int{}

func generateSquares(max int) {
	Printfln("Generating data...")
	for val := range max {
		squares[val] = int(math.Pow(float64(val), 2))
	}
}

func readSquares(id, max, iterations int) {
	once.Do(func() {
		generateSquares(max)
	})
	for i := range iterations {
		_ = i + 1
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	waitGroup.Done()
}

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := range numRoutines {
		go readSquares(i, 10, 5)
	}
	waitGroup.Wait()
	Printfln("Cached values : %v", len(squares))
}
