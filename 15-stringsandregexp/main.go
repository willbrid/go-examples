package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	/** Comparaison des chaines de caractères **/
	product := "Kayak"

	// Cette fonction renvoie true si la chaîne s contient une sous-chaine substr et false sinon.
	fmt.Println("Contains : ", strings.Contains(product, "yak"))
	// Cette fonction renvoie true si la chaîne s contient l'un des caractères contenu dans la sous-chaine substr.
	fmt.Println("ContainsAny : ", strings.ContainsAny(product, "abc"))
	// Cette fonction renvoie true si la chaîne s contient une rune spécifique.
	fmt.Println("ContainsRune : ", strings.ContainsRune(product, 'K'))
	// Cette fonction effectue une comparaison insensible à la casse et renvoie true si les chaînes s1 et s2 sont identiques.
	fmt.Println("EqualFold : ", strings.EqualFold(product, "KAYAK"))
	// Cette fonction renvoie true si la chaîne se termine par le prefixe de chaîne
	fmt.Println("HasPrefix : ", strings.HasPrefix(product, "Ka"))
	// Cette fonction renvoie true si la chaîne se termine par le suffixe de chaîne
	fmt.Println("HasSuffix : ", strings.HasSuffix(product, "yak"))

	/**
	Pour toutes les fonctions du package strings, qui opèrent sur des caractères,
	il existe une fonction correspondante dans le package bytes qui opère sur une tranche d'octet.
	**/
	price := "€100"
	fmt.Println("Strings Prefix : ", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix : ", bytes.HasPrefix([]byte(price), []byte{226, 130}))

	/** Conversion de la casse de chaîne **/
	description := "A boat for sailing"
	fmt.Println("Original : ", description)
	fmt.Println("Title : ", strings.Title(description)) // la fonction Title est obsolète depuis Go 1.18
	fmt.Println("ToUpper : ", strings.ToUpper(description))
	fmt.Println("ToLower : ", strings.ToLower(description))
	fmt.Println("ToTitle : ", strings.ToTitle(description))

	specialChar := "\u01c9"
	fmt.Println("Original : ", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper : ", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title : ", titleChar, []byte(titleChar))

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

	description1 := "A boat for sailing"
	//Permet de compter le nombre de fois la sous chaine "o" apparaît dans la chaine description1
	fmt.Println("Count : ", strings.Count(description1, "o"))
	// Permet de renvoyer l'index de la première occurrence de la sous chaine "o" dans la chaine description1 ou -1 s'il y'a aucune occurrence
	fmt.Println("Index : ", strings.Index(description1, "o"))
	// Permet de renvoyer l'index de la dernière occurrence de la sous chaine "o" dans la chaine description1 ou -1 s'il y'a aucune occurrence
	fmt.Println("LastIndex:", strings.LastIndex(description1, "o"))
	// Permet de renvoyer l'index de la première occurrence de n'importe quel caractère de la chaîne description1, ou -1 s'il n'y a pas d'occurrence.
	fmt.Println("IndexAny:", strings.IndexAny(description1, "abcd"))
	// Permet de renvoyer l'index de la dernière occurrence de n'importe quel caractère de la chaîne description1, ou -1 s'il n'y a pas d'occurrence.
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description1, "abcd"))
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description1, "o"))

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	// Permet de renvoyer l'index de la première occurrence du caractère dans la chaîne description1 pour laquelle la fonction spécifiée renvoie true.
	fmt.Println("IndexFunc:", strings.IndexFunc(description1, isLetterB))
	// Permet de renvoyer l'index de la dernière occurrence du caractère dans la chaîne description1 pour laquelle la fonction spécifiée renvoie true.
	fmt.Println("LastIndexFunc:", strings.LastIndexFunc(description1, isLetterB))

	description2 := "A boat for sailing"
	// Cette fonction Fields divise une chaîne sur les caractères d'espacement et renvoie un slice contenant les éléments non blanches de la chaîne description2
	fieldsResult := strings.Fields(description2)
	fmt.Println("Fields : ", fieldsResult, " - Longueur : ", len(fieldsResult))
	// Cette fonction Split divise la chaîne description2 sur une sous-chaine "a" et renvoie un slice contenant les éléments ne contenant pas la sous-chaine
	splitResult := strings.Split(description2, "a")
	fmt.Println("Split : ", splitResult, " - Longueur : ", len(splitResult))
	// Cette fonction Split divise la chaîne description2 sur une sous-chaine "a" et renvoie un slice contenant un nombre maximal de 2 éléments.
	// Le dernier élément peut contenir la sous-chaine "a" : dans ce cas il est non splité
	splitNResult := strings.SplitN(description2, "a", 2)
	fmt.Println("SplitN : ", splitNResult, " - Longueur : ", len(splitNResult))
	splitAfterResult := strings.SplitAfter(description2, "i")
	// Cette fonction Split divise la chaîne description2 sur une sous-chaine "i" et renvoie un slice contenant les éléments contenant la sous-chaine "i"
	fmt.Println("SplitAfter : ", splitAfterResult, " - Longueur : ", len(splitAfterResult))
	splitAfterNResult := strings.SplitAfterN(description2, "i", 2)
	fmt.Println("SplitAfterN : ", splitAfterNResult, " - Longueur : ", len(splitAfterNResult))

	description3 := "This  is double  spaced"
	// La fonction Fields ne prend pas en charge une limite sur le nombre de résultats mais gère correctement les doubles espaces.
	splitResult1 := strings.Fields(description3)
	fmt.Println("Split : ", splitResult1, " - Longueur : ", len(splitResult1))
	splitter := func(r rune) bool {
		return r == ' '
	}
	fieldsFuncResult := strings.FieldsFunc(description3, splitter)
	fmt.Println("FieldsFunc : ", fieldsFuncResult, " - Longueur : ", len(fieldsFuncResult))

	username := " Alice"
	// Cette fonction permet de supprimer tous les caractères d'espacement de début ou de fin.
	trimSpaceResult := strings.TrimSpace(username)
	fmt.Println("Trimmed : ", ">>"+trimSpaceResult+"<<")
	description4 := "A boat for one person"
	// Cette fonction renvoie une chaîne à partir de laquelle tous les caractères de début ou de fin contenus dans la
	// chaîne "Asno " sont supprimés de la chaîne description4
	trimResult := strings.Trim(description4, "Asno ")
	// Ici la suppression commence à partir de la gauche
	trimLeftResult := strings.TrimLeft(description4, "Asno ")
	// Ici la suppression commence à partir de la droite
	trimRightResult := strings.TrimRight(description4, "Asno ")
	fmt.Println("Trim : ", trimResult)
	fmt.Println("TrimLeft : ", trimLeftResult)
	fmt.Println("TrimRight : ", trimRightResult)
	// Cette fonction supprime la sous-chaine "A boat " au début de la chaine description4 et renvoie le reste de la chaine
	trimPrefixResult := strings.TrimPrefix(description4, "A boat ")
	// Cette fonction supprime la sous-chaine "son" à la fin de la chaine description4 et renvoie le reste de la chaine
	trimSuffixResult := strings.TrimSuffix(description4, "son")
	fmt.Println("TrimPrefix : ", trimPrefixResult)
	fmt.Println("TrimSuffix : ", trimSuffixResult)
	// La fonction personnalisée est appelée pour les caractères au début et à la fin de la chaîne,
	// et les caractères seront coupés jusqu'à ce que la fonction renvoie false.
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimFuncResult := strings.TrimFunc(description4, trimmer)
	fmt.Println("TrimFunc : ", trimFuncResult)
}
