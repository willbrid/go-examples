package main

import (
	"encoding/json"
	"io"
	"strings"
)

/**
Décodage des tableaux (arrays)
Le processus de décodage traite automatiquement les tableaux (arrays), mais il faut être prudent car JSON autorise les
tableaux à contenir des valeurs de types différents, ce qui entre en conflit avec les règles de typage strictes imposées par Go.

Dans l'exemple #1, les données JSON sources contiennent deux tableaux : l’un ne contient que des nombres, l’autre un mélange de nombres et
de chaînes de caractères. Le processus de décodage ne cherche pas à déterminer si un tableau JSON peut être représenté par un seul type Go et
décode chaque tableau en une slice d’interface vide. Chaque valeur est typée en fonction de sa valeur JSON, mais le type de la slice est l’interface vide.

Dans l'exemple #2, l'on connait à l’avance la structure des données JSON et que l'on décode un tableau contenant un seul type de données JSON,
l'on peut alors passer une slice Go du type souhaité à la méthode `Decode`.
L'on peut utiliser une slice d'entiers pour décoder le premier tableau des données JSON, car toutes les valeurs peuvent être représentées par
des entiers Go. Le second tableau contient des valeurs variées ; l'on doit donc spécifier l'interface vide comme type cible.
**/

func main() {
	Printfln("Exemple #1")
	reader := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	vals := []any{}
	decoder := json.NewDecoder(reader)
	for {
		var decodedVal any
		err := decoder.Decode(&decodedVal)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}
	for _, val := range vals {
		Printfln("Decoded (%T): %v", val, val)
	}

	Printfln("Exemple #2")
	reader1 := strings.NewReader(`[10,20,30]["Kayak","Lifejacket",279]`)
	ints := []int{}
	mixed := []any{}
	val1s := []any{&ints, &mixed}
	decoder1 := json.NewDecoder(reader1)
	for i := range val1s {
		err := decoder1.Decode(val1s[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", ints, ints)
	Printfln("Decoded (%T): %v", mixed, mixed)
}
