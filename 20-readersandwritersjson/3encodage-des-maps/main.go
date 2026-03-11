package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Encodage des maps

Les maps Go sont encodées sous forme d'objets JSON, les clés de la map servant de clés à l'objet. Les valeurs contenues dans la map sont
encodées en fonction de leur type. L'exemple ci-dessous illustre l'encodage d'une map contenant des valeurs de type float64.

Les maps peuvent également servir à créer des représentations JSON personnalisées de données Go.
**/

func main() {
	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(m)
	fmt.Println(writer.String())
}
