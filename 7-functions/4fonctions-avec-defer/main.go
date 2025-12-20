package main

import "fmt"

func main() {
	products := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	printPrice("Kayak", 275, 0.2)
	_, total := calcTotalPrice(products)
	fmt.Println("Total :", total)
}
