package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Les slices et les tableaux Go sont encodés sous forme de tableaux JSON, à l'exception des slices d'octets qui sont exprimées sous forme de
chaînes encodées en base64. Les tableaux d'octets, quant à eux, sont encodés sous forme de tableau de nombres JSON.

L'encodeur exprime chaque tableau selon la syntaxe JSON, à l'exception des slices d'octets.
Notons que les tableaux d'octets et les slices d'octets sont traités différemment, même si leur contenu est identique.
**/

func main() {
	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)

	fmt.Println(writer.String())
}
