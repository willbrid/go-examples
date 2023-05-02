package main

import "fmt"

func main() {
	// Traiter les erreurs récupérables
	var categories []string = []string{"Watersports", "Chess", "Running"}
	var total float64
	var err *CategoryError

	for _, cat := range categories {
		total, err = products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, " Total : ", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}

	// Reception d'erreur via un canal
	var channel chan ChannelMessage = make(chan ChannelMessage, 10)
	go products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}
}
