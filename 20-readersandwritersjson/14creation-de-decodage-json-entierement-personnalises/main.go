package main

import (
	"encoding/json"
	"io"
	"strings"
)

/**
Création de décodeurs JSON entièrement personnalisés

Le processus de décodage vérifie si une structure implémente l'interface `Unmarshaler`, qui désigne un type possédant un décodage personnalisé
et qui définit la méthode :
`UnmarshalJSON(byteSlice)` : Cette méthode est appelée pour décoder les données JSON contenues dans la slice d’octets spécifiée. Le résultat est
une erreur indiquant des problèmes de décodage.
**/

func main() {
	reader := strings.NewReader(`{"Name":"Kayak","Category":"Watersports","Price":279, "Offer": "10"}`)
	decoder := json.NewDecoder(reader)
	for {
		var val DiscountedProduct
		err := decoder.Decode(&val)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		} else {
			Printfln("Name: %v, Category: %v, Price: %v, Discount: %v", val.Name, val.Category, val.Price, val.Discount)
		}
	}
}
