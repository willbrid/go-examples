package main

import (
	"math"
	"math/rand"
	"time"
)

// Générer un entier aléatoire entre deux valeurs min et max
func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

var names []string = []string{"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {
	Printfln("Hello, Math and Sorting")

	var val1 float64 = 279.00
	var val2 float64 = 48.95

	/**
	Cette fonction Abs renvoie la valeur absolue d'une valeur float64.
	**/
	Printfln("Abs : %v", math.Abs(val1))
	/**
	Cette fonction Ceil renvoie le plus petit entier égal ou supérieur à la valeur float64 spécifiée.
	Le résultat est également une valeur float64, même s'il représente un nombre entier.
	**/
	Printfln("Ceil : %v", math.Ceil(val2))
	// Cette fonction Copysign renvoie une valeur float64, qui est la valeur absolue de val1 avec le signe de -5.
	Printfln("Copysign : %v", math.Copysign(val1, -5))
	/**
	Cette fonction renvoie le plus grand entier inférieur ou égal à la valeur float64 spécifiée.
	Le résultat est également une valeur float64, même s'il représente un nombre entier.
	**/
	Printfln("Floor : %v", math.Floor(val2))
	// Cette fonction Max renvoie le maximun entre val1 et val2.
	Printfln("Max : %v", math.Max(val1, val2))
	// Cette fonction Max renvoie le minimun entre val1 et val2.
	Printfln("Min: %v", math.Min(val1, val2))
	// Cette fonction Mod renvoie le reste de val1/val2.
	Printfln("Mod: %v", math.Mod(val1, val2))
	// Cette fonction renvoie le résultat de val1 élevé à l'exposant 2.
	Printfln("Pow: %v", math.Pow(val1, 2))
	/**
	Cette fonction arrondit la valeur spécifiée à l'entier le plus proche, en arrondissant les demi-valeurs supérieures.
	Le résultat est une valeur float64, même si elle représente un entier.
	**/
	Printfln("Round: %v", math.Round(val2))
	/**
	Cette fonction arrondit la valeur spécifiée à l'entier le plus proche, en arrondissant les demi-valeurs au nombre pair le plus proche.
	Le résultat est une valeur float64, même si elle représente un entier.
	**/
	Printfln("RoundToEven: %v", math.RoundToEven(val2))

	// Cette fonction rand.Seed définit la valeur de départ à l'aide de la valeur int64 spécifiée.
	// rand.Seed utilise la valeur de départ fournie pour initialiser la source par défaut à un état déterministe.
	rand.Seed(time.Now().UnixNano()) // obsolète depuis 1.20
	for i := 0; i < 5; i++ {
		// Cette fonction rand.Int génère une valeur int aléatoire.
		Printfln("Value rand.Int %v : %v", i, rand.Int())
		// Cette fonction rand.Intn génère un entier aléatoire inférieur à une valeur spécifiée.
		Printfln("Value rand.Intn %v : %v", i, rand.Intn(10))
		Printfln("Value IntRange %v : %v", i, IntRange(10, 20))
	}

	/**
	Les arguments de la fonction Shuffle sont le nombre d'éléments et une fonction qui échangera deux éléments, qui sont identifiés par index.
	La fonction est appelée pour échanger des éléments au hasard.
	**/
	rand.Shuffle(len(names), func(first, second int) {
		names[first], names[second] = names[second], names[first]
	})
	for i, name := range names {
		Printfln("Index %v : - Name: %v", i, name)
	}
}
