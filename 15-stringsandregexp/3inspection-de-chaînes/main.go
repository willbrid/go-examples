package main

import (
	"fmt"
	"strings"
)

/**
Le package `strings` fournit des fonctions pour l'inspection des chaînes de caractères.

Count(s, sub) : Cette fonction renvoie un entier indiquant le nombre d'occurrences de la sous-chaîne sub spécifiée dans la chaîne s.

Index(s, sub) et LastIndex(s, sub) : Ces fonctions renvoient l'indice de la première ou de la dernière occurrence d'une sous-chaîne spécifiée
dans la chaîne s, ou -1 si elle n'est pas trouvée.

IndexAny(s, chars) et LastIndexAny(s, chars) : Ces fonctions renvoient l'indice de la première ou de la dernière occurrence de n'importe
quel caractère dans la chaîne spécifiée au sein de la chaîne s, ou -1 si elle n'est pas trouvée.

IndexByte(s, b) et LastIndexByte(s, b) : Ces fonctions renvoient l'indice de la première ou de la dernière occurrence d'un octet spécifié dans
la chaîne s, ou -1 si elle n'est pas trouvée.

IndexFunc(s, func) et LastIndexFunc(s, func) : Ces fonctions renvoient l'indice de la première ou de la dernière occurrence du caractère dans
la chaîne s pour lequel la fonction spécifiée renvoie vrai
**/

func main() {
	description := "A boat for sailing"

	fmt.Println("Count :", strings.Count(description, "o"))
	fmt.Println("Index :", strings.Index(description, "o"))
	fmt.Println("LastIndex :", strings.LastIndex(description, "o"))
	fmt.Println("IndexAny :", strings.IndexAny(description, "abcd"))
	fmt.Println("LastIndexAny :", strings.LastIndexAny(description, "abcd"))
	fmt.Println("LastIndexAny :", strings.LastIndexAny(description, "o"))

	/**
	Les fonctions personnalisées reçoivent une rune et renvoient un résultat booléen indiquant si le caractère remplit la condition souhaitée.
	La fonction IndexFunc appelle la fonction personnalisée pour chaque caractère de la chaîne jusqu'à obtenir un résultat vrai, auquel cas
	l'index est renvoyé.

	La variable isLetterB se voit attribuer une fonction personnalisée qui reçoit une rune et renvoie vrai si la rune est un B majuscule ou minuscule.
	**/
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	fmt.Println("IndexFunc :", strings.IndexFunc(description, isLetterB))
	fmt.Println("LastIndexFunc :", strings.LastIndexFunc(description, isLetterB))
}
