package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Contrôle de flux")

	/** If else if et else **/
	var kayakPrice float32 = 275.00
	fmt.Println("Price : ", kayakPrice)

	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}

	var kayakPrice1 float32 = 275.00
	if kayakPrice1 > 500 {
		fmt.Println("Price is greater than 500")
	} else if kayakPrice1 < 300 {
		fmt.Println("Price is less than 300")
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}

	var kayakPrice2 float32 = 275.00
	if kayakPrice2 > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice2 < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar)
	}

	priceString := "275"
	if kayakPrice3, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price: ", kayakPrice3)
	} else {
		fmt.Println("Error: ", err)
	}

	/** Utilisation de la boucle For **/
	var counter int = 0
	for {
		fmt.Println("Counter: ", counter)
		counter++
		if counter > 3 {
			break
		}
	}

	var counter1 int = 0
	for counter1 <= 3 {
		fmt.Println("Counter1: ", counter1)
		counter1++
	}

	for counter2 := 0; counter2 <= 3; counter2++ {
		fmt.Println("Counter2: ", counter2)
	}

	for counter3 := 0; true; counter3++ {
		fmt.Println("Counter3: ", counter3)
		if counter3 > 3 {
			break
		}
	}

	for counter4 := 0; counter4 <= 3; counter4++ {
		if counter4 == 1 {
			continue
		}
		fmt.Println("Counter4: ", counter4)
	}

	var product string = "Kayak"
	for index, character := range product {
		fmt.Println("Index: ", index, " Character: ", string(character))
	}

	for index := range product {
		fmt.Println("Index: ", index)
	}

	for _, character := range product {
		fmt.Println("Character: ", string(character))
	}

	var products []string = []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element:", element)
	}

	/** Utilisation de l'instruction switch **/
	var product1 string = "Kayak"
	for index, character := range product1 {
		switch character {
		case 'K':
			fmt.Println("K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		case 'a', 'k':
			fmt.Println("a or k at position", index)
		}
	}

	var product2 string = "Kayak"
	for index, character := range product2 {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}
