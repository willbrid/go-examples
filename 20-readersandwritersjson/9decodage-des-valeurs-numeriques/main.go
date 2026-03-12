package main

import (
	"encoding/json"
	"io"
	"strings"
)

/**
Décodage des valeurs numériques

JSON utilise un seul type de données pour représenter les valeurs à virgule flottante et les entiers. Le type `Decoder` convertit ces
valeurs numériques en valeurs float`64.
Ce comportement peut être modifié en appelant la méthode `UseNumber` du type `Decoder`, ce qui force le décodage des valeurs numériques JSON
vers le type `Number`, défini dans le package `encoding/json`.
Les méthodes définies par le type `Number` (du package `encoding/json`) :
- Int64() : Cette méthode renvoie la valeur décodée sous forme d'entier 64 bits (int64) et un message d'erreur indiquant
si la conversion est impossible.
- Float64() : Cette méthode renvoie la valeur décodée sous forme de nombre à virgule flottante 64 bits (float64) et un message d'erreur indiquant
si la conversion est impossible.
- String() : Cette méthode renvoie la chaîne de caractères non convertie issue des données JSON.

De ce qui en ressort de l'exemple ci-dessous : toutes les valeurs numériques JSON ne peuvent pas être exprimées en tant que valeurs Go `int64`,
c'est pourquoi cette méthode est généralement appelée en premier. Si la conversion en entier échoue, la méthode `Float64` peut être appelée.
Si un nombre ne peut être converti en aucun des deux types Go, la méthode `String` peut être utilisée pour obtenir la chaîne de caractères
non convertie à partir des données JSON.
**/

func main() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []any{}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
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
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded Floating Point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}
}
