package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Contrôle de flux")

	/** If, else if et else **/
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

	// For : utilisation de for avec l'instruction range pour parcourir les chaines, les arrays, les slices et les maps
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

	/**
	    switch : on peut utiliser l'instruction break parmi les instructions d'un cas pour sortir immédiatement du switch
	**/
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

	/**
	    Switch : l'instruction fallthrough exécutée dans un cas oblige l'exécution des instructions du prochain cas
	**/
	var product3 string = "Kayak"
	for index, character := range product3 {
		switch character {
		case 'K':
			fmt.Println("Uppercase character")
			fallthrough
		case 'k':
			fmt.Println("k at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}

	/**
	    Switch : on peut avoir plusieurs valeurs de cas; elles doivent être séparées par une virgule
		L'instruction default est exécutée lorsque tous les autres cas évalués ne sont correspondants
	**/
	var product4 string = "Kayak"
	for index, character := range product4 {
		switch character {
		case 'K', 'k':
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("k at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}

	/**
	  Switch : il y'a duplication de l'expression -> counter5 / 2
	**/
	for counter5 := 0; counter5 < 20; counter5++ {
		switch counter5 / 2 {
		case 2, 3, 5, 7:
			fmt.Println("Prime value: ", counter5/2)
		default:
			fmt.Println("Non-prime value: ", counter5/2)
		}
	}

	/**
	  Switch : on utilise une variable d'initialisaiton (déclaration courte de variable) pour éviter les duplications d'expression.
	**/
	for counter5 := 0; counter5 < 20; counter5++ {
		switch val := counter5 / 2; val {
		case 2, 3, 5, 7:
			fmt.Println("Prime value: ", val)
		default:
			fmt.Println("Non-prime value: ", val)
		}
	}

	/**
		Switch : lorsque la valeur de comparaison est omise au niveau du switch, on utilise les expressions de comparaison dans les cas
	**/
	for counter6 := 0; counter6 < 10; counter6++ {
		switch {
		case counter6 == 0:
			fmt.Println("Zero value")
		case counter6 < 3:
			fmt.Println(counter6, "is < 3")
		case counter6 >= 3 && counter6 < 7:
			fmt.Println(counter6, "is >= 3 && < 7")
		default:
			fmt.Println(counter6, "is >= 7")
		}
	}

	// Target : les instructions d'étiquette permettent à l'exécution de passer à un point différent
	var counter7 int = 0
target:
	fmt.Println("Counter : ", counter7)
	counter7++
	if counter7 < 5 {
		goto target
	}
}
