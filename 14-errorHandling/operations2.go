package main

// Utilisation du package fmt pour formater les erreurs
import "fmt"

type ChannelMessage2 struct {
	Category      string
	Total         float64
	CategoryError error
}

func (slice ProductSlice) TotalPrice2(category string) (total float64, err error) {
	var productCount int = 0

	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		err = fmt.Errorf("cannot find category : %v", category) // Le %v est remplacé par le 2ème argument de la méthode Errorf
	}
	return
}

// Signalement d'erreurs via les canaux
func (slice ProductSlice) TotalPriceAsync2(categories []string, channel chan<- ChannelMessage2) {
	for _, c := range categories {
		total, err := slice.TotalPrice2(c)
		channel <- ChannelMessage2{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
