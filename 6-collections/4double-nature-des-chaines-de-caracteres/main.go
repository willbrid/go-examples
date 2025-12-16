package main

import (
	"fmt"
	"strconv"
)

func main() {
	/** Comprendre la double nature des chaines de caractères
	Go traite les chaînes de caractères comme des tableaux d'octets et prend en charge la notation d'index de tableau et de plage de tranches.
	--- Lorsque la notation d'index est utilisée, le résultat est un octet de l'emplacement spécifié dans la chaîne
	--- Lorsqu'une chaîne est découpée en tranches, la tranche est également décrite à l'aide d'octets, mais le résultat est une chaîne
	**/
	fmt.Println("Comprendre la double nature des chaines de caractères")
	var price string = "$48.95"
	var currencyByte byte = price[0] // Byte est l'alias à uint8
	var currencyString string = string(price[0])
	var amountString string = price[1:] // Ici la tranche est un aussi un tableau d'octets mais dont la combinaison est une chaîne.
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency Byte : ", currencyByte)
	fmt.Println("Currency String : ", currencyString)
	fmt.Println("Length : ", len(price))
	if parseErr == nil {
		fmt.Println("Amount : ", amount)
	} else {
		fmt.Println("Parse Error : ", parseErr)
	}

	/**
	Le type rune représente un point de code Unicode, qui est essentiellement un caractère unique. Pour éviter de découper
	les chaînes au milieu des caractères, une conversion explicite en tranche de rune peut être effectuée.
	**/
	var price1 []rune = []rune("€48.95") // rune est l'alias à int32
	var currency1 string = string(price1[0])
	var amountString1 string = string(price1[1:])
	amount1, parseErr1 := strconv.ParseFloat(amountString1, 64)
	fmt.Println("Currency String : ", currency1)
	fmt.Println("Length : ", len(price1))
	if parseErr1 == nil {
		fmt.Println("Amount : ", amount1)
	} else {
		fmt.Println("Parse Error : ", parseErr1)
	}

	var price2 string = "$48.95"
	for index, char := range price2 {
		// Value correspondance en byte et string(char) correspondance en caractère
		fmt.Println("Index : ", index, " - value : ", char, " - value string : ", string(char))
	}

	/**
	Notez que les valeurs d'index ne sont pas séquentielles. La boucle for traite la chaîne comme une séquence de caractères dérivée
	de la séquence d'octets sous-jacente. Les valeurs d'index correspondent au premier octet composant chaque caractère.
	La deuxième valeur d'index est 3, car le premier caractère de la chaîne est composé d'octets en positions 0, 1 et 2.
	**/
	var price3 string = "€48.95"
	for index, char := range price3 {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("price3 Index : ", index, " - value : ", char, " - value string : ", string(char))
	}

	/**
	Pour énumérer les octets sous-jacents sans les convertir en caractères, nous pouvons effectuer une conversion explicite
	en tranche d'octets.
	**/
	var price4 string = "€48.95"
	for index, char := range []byte(price4) {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("price4 Index : ", index, " - value : ", char, " - value string : ", string(char))
	}

	/**
	Pour énumérer les caractères de la chaine, nous pouvons aussi effectuer une conversion explicite en tranche de rune.
	**/
	var price5 []rune = []rune("€48.95")
	for index, char := range price5 {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("price5 Index : ", index, " - value : ", char, " - value string : ", string(char))
	}
}
