package main

import "fmt"

/**
Go facilite l'expression des conditions exceptionnelles, ce qui permet à une fonction ou une méthode d'indiquer au code appelant
qu'un problème est survenu.

Go fournit une interface prédéfinie nommée `error`. Cette interface requiert que les erreurs définissent une méthode nommée `Error`,
qui renvoie une chaîne de caractères.
type error interface {
    Error() string
}
**/

func main() {
	categories := []string{"Watersports", "Chess"}
	for _, cat := range categories {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total :", ToCurrency(total))
	}

	categorie1s := []string{"Watersports", "Chess", "Running"}
	for _, cat := range categorie1s {
		total := Products.TotalPrice(cat)
		fmt.Println(cat, "Total :", ToCurrency(total))
	}

	categorie2s := []string{"Watersports", "Chess", "Running"}
	for _, cat := range categorie2s {
		total, err := Products.TotalPriceWithCategoryError(cat)
		if err == nil {
			fmt.Println(cat, "Total :", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}

	categorie3s := []string{"Watersports", "Chess", "Running"}
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categorie3s, channel)
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total :", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}

	categorie4s := []string{"Watersports", "Chess", "Running"}
	channel1 := make(chan ChannelResult, 10)
	go Products.TotalPriceResultAsync(categorie4s, channel1)
	for message := range channel1 {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total :", ToCurrency(message.Total))
		} else {
			fmt.Println(message.Category, "(no such category)")
		}
	}
}
