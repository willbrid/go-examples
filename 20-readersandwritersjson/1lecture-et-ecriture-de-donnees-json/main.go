package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/**
Lecture et écriture de données JSON
Le package `encoding/json` permet d'encoder et de décoder des données JSON. Ci-dessous les fonctions constructeur utilisées pour créer
les structures d'encodage et de décodage des données JSON.

- NewEncoder(writer) : Cette fonction renvoie une valeur de type `Encoder` permettant d’encoder des données JSON et de les écrire dans le writer spécifié.
- NewDecoder(reader) : Cette fonction renvoie une valeur de type `Decoder` permettant de lire des données JSON depuis le reader spécifié et de les décoder.

Le package `encoding/json` fournit également des fonctions pour l'encodage et le décodage JSON sans utiliser de reader ou de writer.
- Marshal(value) : Cette fonction encode la valeur spécifiée au format JSON. Elle renvoie le contenu JSON sous forme de slice d'octets,
ainsi qu'une erreur indiquant tout problème d'encodage.
- Unmarshal(byteSlice, val) : Cette fonction analyse les données JSON contenues dans la slice d'octets spécifiée et assigne
le résultat à la valeur spécifiée.


Encodage des données JSON
La fonction constructeur `NewEncoder` permet de créer une valeur de type `Encoder`, qui peut être utilisé pour écrire des données JSON dans un writer.
Les méthodes des valeurs de type `Encoder` :
- Encode(val) : Cette méthode encode la valeur spécifiée au format JSON et l'écrit dans le Writer.
- SetEscapeHTML(on) : Cette méthode accepte un argument booléen qui, s'il est vrai, encode les caractères dangereux pour le HTML en les échappant.
Par défaut, ces caractères sont échappés.
- SetIndent(prefix, indent) : Cette méthode spécifie un préfixe et une indentation appliqués au nom de chaque champ dans la sortie JSON.


Représentation des types de données de base de Go en JSON
- bool : Les valeurs booléennes Go sont exprimées sous forme de valeurs JSON true ou false.
- string : Les valeurs de type chaîne de caractères Go sont exprimées sous forme de chaînes JSON. Par défaut, les caractères HTML non sécurisés
sont échappés.
- float32, float64 : Les valeurs à virgule flottante Go sont exprimées sous forme de nombres JSON.
- int, int<size> : Les valeurs entières Go sont exprimées sous forme de nombres JSON.
- uint, uint<size> : Les valeurs entières Go sont exprimées sous forme de nombres JSON.
- byte : Les octets Go sont exprimés sous forme de nombres JSON.
- rune : Les runes Go sont exprimées sous forme de nombres JSON.
- nil : La valeur nil Go est exprimée sous forme de valeur nulle JSON.
- Pointeurs : L'encodeur JSON suit les pointeurs et encode la valeur à l'adresse du pointeur.

Cet exemple ci-dessous définit une série de variables de différents types de base. Le constructeur NewEncoder sert à créer une valeur de type Encoder,
et une boucle for est utilisée pour encoder chaque valeur au format JSON. Les données sont écrites dans une valeur de type `Builder`, dont
la méthode `String` est appelée pour afficher le JSON.
**/

func main() {
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var irune rune = 'c'
	var ibyte []byte = []byte("cool")
	var pointer *int = &ival

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	for _, val := range []any{b, str, fval, ival, irune, ibyte, pointer} {
		encoder.Encode(val)
	}
	fmt.Println(writer.String())
}
