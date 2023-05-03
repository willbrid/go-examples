package main

// Utilisation du package errors pour créer un commentaire d'erreur
import "errors"

type ChannelMessage1 struct {
	Category      string
	Total         float64
	CategoryError error
}

func (slice ProductSlice) TotalPrice1(category string) (total float64, err error) {
	var productCount int = 0

	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = errors.New("cannot find category")
	}
	return
}

// Signalement d'erreurs via les canaux
func (slice ProductSlice) TotalPriceAsync1(categories []string, channel chan<- ChannelMessage1) {
	for _, c := range categories {
		total, err := slice.TotalPrice1(c)
		channel <- ChannelMessage1{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
