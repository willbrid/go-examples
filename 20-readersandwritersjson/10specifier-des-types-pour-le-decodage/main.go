package main

import (
	"encoding/json"
	"strings"
)

/**
Si l'on connait la structure des données JSON à décoder, l'on peut indiquer à l'instruction de décodage d'utiliser des
types Go spécifiques en utilisant des variables de ce type pour recevoir une valeur décodée.
Le processus de décodage renverra une erreur s'il ne parvient pas à décoder une valeur JSON dans un type spécifié. Cette technique ne doit être
utilisée que si nous sommes certain de bien comprendre les données JSON à décoder.
**/

func main() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	var bval bool
	var sval string
	var fpval float64
	var ival int
	vals := []any{&bval, &sval, &fpval, &ival}
	decoder := json.NewDecoder(reader)
	for i := range vals {
		err := decoder.Decode(vals[i])
		if err != nil {
			Printfln("Error: %v", err.Error())
			break
		}
	}
	Printfln("Decoded (%T): %v", bval, bval)
	Printfln("Decoded (%T): %v", sval, sval)
	Printfln("Decoded (%T): %v", fpval, fpval)
	Printfln("Decoded (%T): %v", ival, ival)
}
