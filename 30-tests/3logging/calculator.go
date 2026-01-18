package main

import (
	"log"
	"sort"
)

/**
La classe Logger est créée avec un nouveau préfixe et l'ajout de l'indicateur `Lmsgprefix`, en utilisant le Writer obtenu à partir
de la fonction Output. Le résultat est que les messages de log sont toujours écrits vers la même destination, mais avec un
préfixe supplémentaire qui désigne les messages de la fonction `SortAndTotal`.
**/

func SortAndTotal(vals []int) (sorted []int, total int) {
	var logger = log.New(log.Writer(), "sortAndTotal: ", log.Flags()|log.Lmsgprefix)
	logger.Printf("Invoked with %v values", len(vals))
	sorted = make([]int, len(vals))

	copy(sorted, vals)
	sort.Ints(sorted)

	logger.Printf("Sorted data : %v", sorted)
	for _, val := range sorted {
		total += val
	}
	logger.Printf("Total : %v", total)

	return
}
