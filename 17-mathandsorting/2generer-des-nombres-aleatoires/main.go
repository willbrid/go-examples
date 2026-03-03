package main

import (
	"math/rand"
	"time"
)

/**
Le package math/rand permet de générer des nombres aléatoires.

Float32() : Cette fonction génère une valeur float32 aléatoire comprise entre 0 et 1.

Float64() : Cette fonction génère une valeur float64 aléatoire comprise entre 0 et 1.

Int() : Cette fonction génère une valeur int aléatoire.

Intn(max) : Cette fonction génère un int aléatoire inférieur à une valeur spécifiée.

UInt32() : Cette fonction génère une valeur uint32 aléatoire.

UInt64() : Cette fonction génère une valeur uint64 aléatoire.

Shuffle(count, func) : Cette fonction permet de mélanger aléatoirement l'ordre des éléments.
**/

/**
Il n'existe pas de fonction permettant de spécifier une valeur minimale, mais il est facile de ramener les valeurs générées par la
fonction Intn dans une plage spécifique.
**/
// La fonction IntRange renvoie un nombre aléatoire dans une plage spécifique.
func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	for i := range 5 {
		Printfln("rand.Int - Value %v : %v", i, rand.Int())
	}

	/**
	Utilisation d'un générateur local avec rand.New
	**/
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range 5 {
		Printfln("r.Int - Value %v : %v", i, r.Int())
	}

	for i := range 5 {
		Printfln("rand.Intn - Value %v : %v", i, rand.Intn(10))
	}

	for i := range 5 {
		Printfln("IntRange - Value %v : %v", i, IntRange(10, 20))
	}

	/**
	La fonction Shuffle est utilisée pour réorganiser aléatoirement les éléments, ce qu'elle fait à l'aide d'une fonction personnalisée.
	La fonction Shuffle prend comme arguments le nombre d'éléments et une fonction qui échange deux éléments, identifiés par leur index.
	Cette fonction (2ème argument) est appelée pour permuter les éléments de manière aléatoire.
	**/
	names := []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}
	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Shuffle - Index %v : Name: %v", i, name)
	}
}
