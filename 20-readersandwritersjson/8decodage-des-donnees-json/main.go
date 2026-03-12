package main

import (
	"encoding/json"
	"io"
	"strings"
)

/**
Décodage des données JSON

La fonction constructeur `NewDecoder` crée une valeur de type `Decoder` permettant de décoder les données JSON obtenues d'un reader.
Les méthodes des valeurs de type `Decoder` :

Decode(value) : Cette méthode lit et décode des données afin de créer la valeur spécifiée. Elle renvoie une erreur indiquant un
problème de décodage des données vers le type requis ou EOF.

DisallowUnknownFields() : Par défaut, lors du décodage d'un type structure, le processus de décodage ignore toute clé des données JSON
pour laquelle il n'existe pas de champ structure correspondant. L'appel de cette méthode force le décodage à renvoyer une erreur
au lieu d'ignorer la clé.

UseNumber() : Par défaut, les valeurs numériques JSON sont décodées en valeurs float64. L'appel de cette méthode utilise le type `Number`.


Dans l'exemple ci-dessous, l'on commence par créer un `Reader` qui produit des données à partir d'une chaîne de caractères contenant une séquence de
valeurs séparées par des espaces (la spécification JSON autorise les espaces ou les sauts de ligne pour séparer les valeurs).
Puis un `Decoder` qui lit depuis ce Reader. Ensuite, la méthode `Decode` est appelée dans une boucle `for` pour lire successivement chaque valeur JSON.
Pour permettre au décodeur de choisir automatiquement le type Go approprié, on passe un pointeur vers une interface vide (interface{}) à Decode.
La méthode renvoie une erreur pour signaler soit un problème de décodage, soit la fin des données avec `io.EOF`.
La première boucle lit les valeurs jusqu’à atteindre `EOF`, puis une seconde boucle affiche le type et la valeur de chaque donnée décodée.
**/

func main() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
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
}
