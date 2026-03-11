package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Création d'encodages JSON entièrement personnalisés

Le processus d'encodage vérifie si une structure implémente l'interface `Marshaler`, qui désigne un type possédant un encodage personnalisé
et qui définit la méthode
- MarshalJSON() : Cette méthode est appelée pour créer une représentation JSON d'une valeur et renvoie une slice d'octets contenant le JSON et
une erreur indiquant les problèmes d'encodage.
**/

func main() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct{
		Product:  &kayak,
		Discount: 10.50,
	}
	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	fmt.Println(writer.String())
}
