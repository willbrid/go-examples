package main

/**
Décodage des maps
Les objets JavaScript sont exprimés sous forme de paires clé-valeur, ce qui facilite leur décodage en maps Go.
L'approche la plus sûre consiste à définir une map avec des clés de type chaîne de caractères et des valeurs d'interface vides,
ce qui garantit que toutes les paires clé-valeur des données JSON peuvent être décodées dans la `map`.
Une fois le JSON décodé, une boucle « for » est utilisée pour parcourir le contenu de la `map`.

Un seul objet JSON peut contenir plusieurs types de données comme valeurs, mais si nous savons à l'avance que nous allons décoder un objet JSON
ne contenant qu'un seul type de valeur, nous pouvons être plus précis lors de la définition de la map dans laquelle les données seront décodées.
**/

import (
	"encoding/json"
	"strings"
)

func main() {
	reader := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m := map[string]any{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&m)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m, m)
		for k, v := range m {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}

	reader1 := strings.NewReader(`{"Kayak" : 279, "Lifejacket" : 49.95}`)
	m1 := map[string]float64{}
	decoder1 := json.NewDecoder(reader1)
	err1 := decoder1.Decode(&m1)
	if err1 != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Map: %T, %v", m1, m1)
		for k, v := range m1 {
			Printfln("Key: %v, Value: %v", k, v)
		}
	}
}
