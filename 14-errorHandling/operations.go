package main

type CategoryError struct {
	requestedCategory string
}

/**
cette méthode implémente l'interface native :
type error interface {
    Error() string
}
**/
func (e *CategoryError) Error() string {
	return "Category " + e.requestedCategory + " does not exist"
}

type ChannelMessage struct {
	Category string
	Total    float64
	*CategoryError
}

func (slice ProductSlice) TotalPrice(category string) (total float64, err *CategoryError) {
	var productCount int = 0

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

// Signalement d'erreurs via les canaux
func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}
