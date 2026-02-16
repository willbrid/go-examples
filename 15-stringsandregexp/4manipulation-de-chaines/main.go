package main

import (
	"fmt"
	"strings"
)

/**
Le package `strings` fournit des fonctions utiles pour l'édition de chaînes de caractères, notamment la possibilité de remplacer tout ou partie
des caractères ou de supprimer les espaces.

- Le premier ensemble de fonctions sert à découper des chaînes de caractères.

Fields(s) : Cette fonction découpe une chaîne en fonction des espaces et renvoie un tableau contenant les parties de la chaîne s ne contenant pas
d'espaces.

FieldsFunc(s, func) : Cette fonction découpe la chaîne s en fonction des caractères pour lesquels une fonction personnalisée renvoie vrai et
renvoie un tableau contenant les parties restantes de la chaîne.

Split(s, sub) : Cette fonction découpe la chaîne s à chaque occurrence de la sous-chaîne spécifiée et renvoie un tableau de chaînes.
Si le séparateur est une chaîne vide, le tableau contiendra une chaîne pour chaque caractère.

SplitN(s, sub, max) : Cette fonction est similaire à Split, mais accepte un argument entier supplémentaire spécifiant le nombre maximal de sous-chaînes
à renvoyer. La dernière sous-chaîne du tableau de résultats contiendra la partie non découpée de la chaîne source.

SplitAfter(s, sub) : Cette fonction est similaire à Split, mais inclut la sous-chaîne utilisée dans les résultats.

SplitAfterN(s, sub, max) : Cette fonction est similaire à SplitAfter, mais accepte un argument entier supplémentaire qui spécifie le nombre maximal
de sous-chaînes à renvoyer.
**/

func main() {
	description := "A boat for sailing"

	fields := strings.Fields(description)
	fmt.Println("Fields :", fields, " - Longueur : ", len(fields))
	splits := strings.Split(description, "a")
	fmt.Println("Split :", splits, " - Longueur : ", len(splits))
	splitNs := strings.SplitN(description, "a", 2)
	fmt.Println("SplitN :", splitNs, " - Longueur : ", len(splitNs))
	splitAfters := strings.SplitAfter(description, "i")
	fmt.Println("SplitAfter :", splitAfters, " - Longueur : ", len(splitAfters))
	splitAfterNs := strings.SplitAfterN(description, "i", 2)
	fmt.Println("SplitAfterN :", splitAfterNs, " - Longueur : ", len(splitAfterNs))

	phrase := "This  is double  spaced"
	// La fonction Fields ne prend pas en charge une limite sur le nombre de résultats mais gère correctement les doubles espaces.
	fieldPhraseTab := strings.Fields(phrase)
	fmt.Println("Fields (for double space) :", fieldPhraseTab, " - Longueur :", len(fieldPhraseTab))

	splitter := func(r rune) bool {
		return r == ' '
	}
	fieldsFuncs := strings.FieldsFunc(phrase, splitter)
	fmt.Println("FieldsFunc : ", fieldsFuncs, " - Longueur : ", len(fieldsFuncs))
}
