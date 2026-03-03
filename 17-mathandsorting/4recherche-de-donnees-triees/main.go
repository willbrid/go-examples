package main

import "sort"

/**
Le package `sort` définit les fonctions permettant de rechercher une valeur spécifique dans des données triées.

SearchInts(slice, val) : Cette fonction recherche la valeur entière spécifiée dans la slice triée. Le résultat est l'indice de la valeur spécifiée
ou, si la valeur est introuvable, l'indice auquel elle peut être insérée tout en conservant l'ordre de tri.

SearchFloat64s(slice, val) : Cette fonction recherche la valeur flottante 64 bits spécifiée dans la slice triée.
Le résultat est l'indice de la valeur spécifiée ou, si la valeur est introuvable, l'indice auquel elle peut être insérée tout en conservant
l'ordre de tri.

SearchStrings(slice, val) : Cette fonction recherche la valeur de chaîne de caractères spécifiée dans la tranche triée.
Le résultat est l'indice de la valeur spécifiée ou, si la valeur est introuvable, l'indice auquel elle peut être insérée tout en conservant
l'ordre de tri.

Search(count, testFunc) : Cette fonction appelle la fonction de test pour le nombre d'éléments spécifié. Le résultat est l'indice pour lequel
la fonction renvoie vrai. S'il n'y a pas de correspondance, le résultat est l'index auquel la valeur spécifiée peut être insérée pour maintenir
l'ordre de tri.

Les fonctions décrites ci-dessus présentent une particularité : lorsqu'une valeur est trouvée, elles renvoient sa position dans la tranche.
Mais, si la valeur est introuvable, le résultat correspond à la position où elle peut être insérée tout en préservant l'ordre de tri.
**/

func main() {
	ints := []int{9, 4, 2, -1, 10}

	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints Sorted: %v", sortedInts)
	indexOf4 := sort.SearchInts(sortedInts, 4)
	indexOf3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v", indexOf4)
	Printfln("Index of 3: %v", indexOf3)

	/**
	Ces fonctions nécessitent un test supplémentaire pour vérifier si la valeur à l'emplacement renvoyé par ces fonctions est bien celle
	qui a été recherchée.
	**/
	Printfln("Index of 4: %v (present: %v)", indexOf4, sortedInts[indexOf4] == 4)
	Printfln("Index of 3: %v (present: %v)", indexOf3, sortedInts[indexOf3] == 3)
}
