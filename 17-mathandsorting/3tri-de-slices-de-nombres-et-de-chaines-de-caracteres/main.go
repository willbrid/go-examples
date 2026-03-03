package main

import "sort"

/**
Pour organiser les éléments selon une séquence plus prévisible, on utilise les fonctions fournies par le package de `sort`.

Ces fonctions ci-dessous servent à trier des slices contenant des valeurs de type int, float64 ou string.
Chaque type de données possède son propre ensemble de fonctions permettant de trier les données ou de déterminer si elles sont déjà triées.

Float64s(slice) : Cette fonction trie une slice de valeurs float64. Les éléments sont triés sur place.

Float64sAreSorted(slice) : Cette fonction renvoie vrai si les éléments de la slice float64 spécifiée sont triés.

Ints(slice) : Cette fonction trie une slice de valeurs int. Les éléments sont triés sur place.

IntsAreSorted(slice) : Cette fonction renvoie vrai si les éléments de la slice int spécifiée sont triés.

Strings(slice) : Cette fonction trie une slice de valeurs string. Les éléments sont triés sur place.

StringsAreSorted(slice) : Cette fonction renvoie vrai si les éléments de la slice string spécifiée sont triés.
**/

func main() {
	ints := []int{9, 4, 2, -1, 10}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints Sorted: %v", ints)

	floats := []float64{279, 48.95, 19.50}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats Sorted: %v", floats)

	strings := []string{"Kayak", "Lifejacket", "Stadium"}
	Printfln("Strings: %v", strings)
	if !sort.StringsAreSorted(strings) {
		sort.Strings(strings)
		Printfln("Strings Sorted: %v", strings)
	} else {
		Printfln("Strings Already Sorted: %v", strings)
	}

	/**
	Ces fonctions ci-dessus trient les éléments sur place, au lieu de créer une nouvelle tranche. Si nous souhaitons créer une nouvelle slice triée,
	nous devons utiliser les fonctions intégrées `make` et `copy`, comme dans l'exemple ci-dessous.
	**/
	int1s := []int{9, 4, 2, -1, 10}
	sortedInt1s := make([]int, len(int1s))
	copy(sortedInt1s, int1s)
	sort.Ints(sortedInt1s)
	Printfln("Int1s: %v", int1s)
	Printfln("Int1s Sorted: %v", sortedInt1s)
}
