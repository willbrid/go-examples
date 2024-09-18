package main

import "fmt"

func main() {
	shoe, _ := getProduct("shoe")
	shirt, _ := getProduct("shirt")

	printDetails(shoe)
	printDetails(shirt)
}

func printDetails(p IProduct) {
	fmt.Printf("Product: %s", p.getName())
	fmt.Println()
	fmt.Printf("Price: %d", p.getPrice())
	fmt.Println()
}
