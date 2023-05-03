package main

import "fmt"

func main() {
	/**------------- Traiter les erreurs récupérables -----------------**/
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

	// Utilisation de l'intégration du package errors
	var channel1 chan ChannelMessage1 = make(chan ChannelMessage1, 10)
	go products.TotalPriceAsync1(categories, channel1)
	for message := range channel1 {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}

	// Utilisation de l'intégration du package errors
	var channel2 chan ChannelMessage2 = make(chan ChannelMessage2, 10)
	go products.TotalPriceAsync2(categories, channel2)
	for message := range channel2 {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}

	/**------------- Traiter les erreurs non récupérables -----------------**/
}
