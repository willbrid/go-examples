package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Encodage d'interfaces
Le processus d'encodage JSON peut être utilisé sur les valeurs affectées aux variables d'interface, mais c'est le type dynamique qui est encodé.

La slice `namedItems` contient différents types dynamiques.
Aucun aspect de l'interface n'est utilisé pour adapter le JSON, et tous les champs exportés de chaque valeur de la slice sont inclus dans le JSON.
Cela peut s'avérer utile, mais il convient d'être prudent lors du décodage de ce type de JSON, car chaque valeur peut avoir un ensemble de
champs différent.
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
