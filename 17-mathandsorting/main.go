package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// Générer un entier aléatoire entre deux valeurs min et max
func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// Fonction permettant de déviner la saisie d'un utilisateur
func GuessingGame() {
	var s string
	Printfln("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d ? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
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

	// Cette fonction rand.NewSource définit la valeur de départ à l'aide de la valeur int64 spécifiée.
	// rand.NewSource utilise la valeur de départ fournie pour initialiser la source par défaut à un état déterministe.
	rand.NewSource(time.Now().UnixNano()) // obsolète depuis 1.20
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

	var ints []int = []int{9, 4, 2, -1, 10}
	Printfln("Ints : %v", ints)
	// Cette fonction sort.Ints permet de trier un tableau dont les éléments sont de type int
	sort.Ints(ints)
	Printfln("Ints : %v", ints)

	var floats []float64 = []float64{279, 48.95, -100, 19.50}
	Printfln("Floats : %v", floats)
	// Cette fonction sort.Float64s permet de trier un tableau dont les éléments sont de type float64
	sort.Float64s(floats)
	Printfln("Floats : %v", floats)

	var strings []string = []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		// Cette fonction sort.Strings permet de trier un tableau dont les éléments sont de type string
		sort.Strings(strings)
		Printfln("Strings Sorted : %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}

	var int1s []int = []int{9, 4, 2, -1, 10}
	sortedInts := make([]int, len(int1s))
	// cette fonction copy permet de copier un tableau dans un autre
	copy(sortedInts, int1s)
	sort.Ints(sortedInts)
	Printfln("Ints : %v", int1s)
	Printfln("Ints Sorted : %v", sortedInts)

	/**
	Cette fonction sort.SearchInts recherche la tranche (slice) triée pour la valeur int spécifiée. Le résultat est l'index de la valeur spécifiée ou,
	si la valeur n'est pas trouvée, l'index auquel la valeur peut être insérée tout en conservant l'ordre de tri.
	**/
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4 : %v - (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
	Printfln("Index of 3 : %v - (present: %v)", indexOf3, sortedInts[indexOf3] == 3)

	// Lancement de la fonction GuessingGame
	GuessingGame()

	var product1s []Product = []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlices(product1s)
	for _, p := range product1s {
		Printfln("Name : %v, Price: %.2f", p.Name, p.Price)
	}

	var product2s []Product = []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	ProductSlicesByName(product2s)
	for _, p := range product2s {
		Printfln("Name : %v, Price: %.2f", p.Name, p.Price)
	}

	var product3s []Product = []Product{
		{"Kayak", 279},
		{"Lifejacket", 49.95},
		{"Soccer Ball", 19.50},
	}
	SortWith(product3s, func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	})
	for _, p := range product3s {
		Printfln("Name : %v, Price: %.2f", p.Name, p.Price)
	}
}
