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
}
