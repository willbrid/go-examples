package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
Le package strings fournit les fonctions pour modifier la casse des chaînes de caractères.

ToLower(str) : Cette fonction renvoie une nouvelle chaîne contenant les caractères de la chaîne spécifiée, en minuscules.
ToUpper(str) : Cette fonction renvoie une nouvelle chaîne contenant les caractères de la chaîne spécifiée, en minuscules.
ToTitle(str) : Cette fonction renvoie une nouvelle chaîne contenant les caractères de la chaîne spécifiée, en respectant la casse des titres.


Le paquet `unicode` fournit des fonctions qui peuvent être utilisées pour déterminer ou modifier la casse des caractères individuels.
IsLower(rune) : Cette fonction renvoie vrai si la rune spécifiée est en minuscules.
ToLower(rune) : Cette fonction renvoie la rune en minuscules associée à la rune spécifiée.
IsUpper(rune) : Cette fonction renvoie vrai si la rune spécifiée est en majuscules.
ToUpper(rune) : Cette fonction renvoie la rune en majuscules associée à la rune spécifiée.
IsTitle(rune) : Cette fonction renvoie vrai si la rune spécifiée est en casse de titre.
ToTitle(rune) : Cette fonction renvoie la rune en casse de titre associée à la rune spécifiée.
**/

func main() {
	description := "A boat for sailing"
	fmt.Println("Original :", description)
	fmt.Println("ToUpper :", strings.ToUpper(description))
	fmt.Println("ToLower :", strings.ToLower(description))
	fmt.Println("ToTitle :", strings.ToTitle(description))

	/**
	Dans certaines langues, l'apparence de certains caractères change lorsqu'ils sont utilisés dans un titre. Unicode définit trois états pour
	chaque caractère : minuscule, majuscule et casse de titre. La fonction ToTitle renvoie une chaîne contenant uniquement des caractères en
	casse de titre. Cela a le même effet que la fonction ToUpper pour l'anglais, mais peut produire des résultats différents dans d'autres langues.
	**/
	specialChar := "\u01c9"
	fmt.Println("Original :", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper :", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title :", titleChar, []byte(titleChar))

	/**
	le type rune est un alias pour le type int32. Il représente un point de code Unicode unique et peut être utilisé pour représenter
	des caractères Unicode dans une chaîne ou un flux de caractères.
	**/
	product1 := "Kayak"
	for _, char := range product1 {
		fmt.Println(string(char), " Upper case : ", unicode.IsUpper(char)) // Renvoie si le caractère rune est en majuscule
		fmt.Println(string(char), " Lower case : ", unicode.IsLower(char)) // Renvoie si le caractère rune est en minuscule
		fmt.Println(string(char), " Title case : ", unicode.IsTitle(char)) // Renvoie si le caractère rune est la casse du titre
	}
}
