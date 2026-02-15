package main

import (
	"bytes"
	"fmt"
	"strings"
)

/**
Le package `strings` fournit un ensemble de fonctions pour le traitement des chaînes de caractères.

`Contains(s, substr)` : Cette fonction renvoie `true` si la chaîne `s` contient `substr`, et `false` sinon.
`ContainsAny(s, substr)` : Cette fonction renvoie `true` si la chaîne `s` contient au moins un des caractères de la chaîne `substr`.
`ContainsRune(s, rune)` : Cette fonction renvoie `true` si la chaîne `s` contient une rune spécifique.
`EqualFold(s1, s2)` : Cette fonction effectue une comparaison insensible à la casse et renvoie `true` si les chaînes `s1` et `s2` sont identiques.
`HasPrefix(s, prefix)` : Cette fonction renvoie `true` si la chaîne `s` commence par le préfixe spécifié.
`HasSuffix(s, suffix)` : Cette fonction renvoie `true` si la chaîne `s` se termine par le suffixe spécifié.

Pour toutes les fonctions du package `strings`, qui opèrent sur des caractères, il existe une fonction correspondante dans le package `bytes`
qui opère sur une tranche d'octets.
**/

func main() {
	product := "kayak"
	fmt.Println("Product :", product)

	fmt.Println("Contains :", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny :", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune :", strings.ContainsRune(product, 'K'))
	fmt.Println("EqualFold :", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix :", strings.HasPrefix(product, "Ka"))
	fmt.Println("HasSuffix :", strings.HasSuffix(product, "yak"))

	/**
	Cet exemple illustre l'utilisation de la fonction `HasPrefix` fournie par les deux packages. La version du package `strings` opère sur les caractères
	et vérifie le préfixe, indépendamment du nombre d'octets utilisés par ces caractères. Cela permet de déterminer si la chaîne de prix commence par le
	symbole de l'euro. La version de la fonction `bytes` permet de déterminer si la variable `price` commence par une séquence d'octets spécifique,
	indépendamment de la manière dont ces octets correspondent à un caractère.
	**/
	price := "€100"
	fmt.Println("Strings Prefix :", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix :", bytes.HasPrefix([]byte(price), []byte{226, 130}))
}
