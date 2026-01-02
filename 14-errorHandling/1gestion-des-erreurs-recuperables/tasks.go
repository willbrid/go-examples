package main

import (
	"errors"
	"fmt"
)

/**
Il peut être fastidieux de définir des types de données pour chaque type d'erreur qu'une application peut rencontrer. Le package `errors`,
qui fait partie de la bibliothèque `standard`, fournit une fonction `New` qui renvoie une erreur dont le contenu est une chaîne de caractères.
L'inconvénient de cette approche est qu'elle génère des erreurs simples, mais elle a l'avantage de la simplicité.
**/

type ChannelResult struct {
	Category      string
	Total         float64
	CategoryError error
}

func (slice ProductSlice) TotalPriceWithError(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = errors.New("Cannot find category")
	}

	return
}

func (slice ProductSlice) TotalPriceWithFormattedError(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = fmt.Errorf("Cannot find category : %v", category)
	}

	return
}

func (slice ProductSlice) TotalPriceResultAsync(categories []string, channel chan<- ChannelResult) {
	for _, c := range categories {
		total, err := slice.TotalPriceWithError(c)
		channel <- ChannelResult{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
