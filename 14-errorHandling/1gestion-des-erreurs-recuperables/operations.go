package main

/**
Les fonctions et les méthodes peuvent exprimer des résultats exceptionnels ou inattendus en produisant des réponses d'erreur.
**/

type CategoryError struct {
	requestedCategory string
}

func (e *CategoryError) Error() string {
	return "Category : " + e.requestedCategory + " does not exist"
}

/**
Si une fonction est exécutée à l'aide d'une goroutine, la communication se fait exclusivement par le biais du canal dédié. Cela signifie que
les détails de tout problème doivent être communiqués en même temps que les opérations réussies. Il est important de simplifier au maximum
la gestion des erreurs ; il est préférable d'éviter d'utiliser des canaux supplémentaires ou de créer des mécanismes complexes pour signaler
les erreurs en dehors du canal dédié. Une approche consiste à créer un type personnalisé qui regroupe les deux types de résultats.
**/

type ChannelMessage struct {
	Category string
	Total    float64
	*CategoryError
}

func (slice ProductSlice) TotalPrice(category string) (total float64) {
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
		}
	}

	return
}

func (slice ProductSlice) TotalPriceWithCategoryError(category string) (total float64, err *CategoryError) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = &CategoryError{requestedCategory: category}
	}

	return
}

/**
Le type `ChannelMessage` permet de communiquer la paire de résultats nécessaires pour refléter fidèlement le résultat de la méthode
`TotalPriceWithError`, exécutée de manière asynchrone par la nouvelle méthode `TotalPriceAsync`. Le résultat est similaire à la manière
dont les résultats des méthodes synchrones peuvent exprimer des erreurs.
**/

func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPriceWithCategoryError(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
