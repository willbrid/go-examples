package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Hello, Collections")

	fmt.Println("Travailler avec des tableaux")
	var names [3]string
	fmt.Println(names)
	names[0] = "Kayak"
	names[1] = "Lifejacket"
	names[2] = "Paddle"
	fmt.Println(names)

	var names1 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names1)

	names2 := [3]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names2)

	names3 := [5]string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names3)

	var coords [3][3]int
	coords[0][1] = 7
	coords[1][2] = 10
	fmt.Println(coords)

	var names4 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names5 [3]string = names4

	names[0] = "Canoe"
	fmt.Println(names)
	fmt.Println(names5)

	var names6 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var pointerNames6 *[3]string = &names6
	names6[0] = "Canoe"
	fmt.Println("Names : ", names6)
	fmt.Println("Pointer : ", *pointerNames6)

	var names7 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var pointerNames7Element1 *string = &names7[1]
	fmt.Println("Pointer to Element 1 Before : ", *pointerNames7Element1)
	names7[1] = "Canoe"
	fmt.Println("Names : ", names7)
	fmt.Println("Pointer to Element 1 After : ", *pointerNames7Element1)

	var names8 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var names9 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	var isSame bool = names8 == names9
	fmt.Println("Comparaison : ", isSame)

	var names10 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for index, value := range names10 {
		fmt.Println("Index : ", index, " - Value : ", value)
	}

	var names11 [3]string = [3]string{"Kayak", "Lifejacket", "Paddle"}
	for _, value := range names11 {
		fmt.Println("Value : ", value)
	}

	fmt.Println("Travailler avec Slices : Tableau dont on ne connait pas sa longueur en avance ou dont sa longueur est variable.")
	var names12 []string = make([]string, 3)
	names12[0] = "Kayak"
	names12[1] = "Lifejacket"
	names12[2] = "Paddle"
	fmt.Println(names12)

	names13 := []string{"Kayak", "Lifejacket", "Paddle"}
	fmt.Println(names13)

	names14 := []string{"Kayak", "Lifejacket", "Paddle"}
	names14 = append(names14, "Hat", "Gloves")
	fmt.Println(names14)

	var names15 []string = []string{"Kayak", "Lifejacket", "Paddle"}
	var appendedNames15 []string = append(names15, "Hat", "Gloves")
	names15[0] = "Canoe"
	fmt.Println(names15)
	fmt.Println(appendedNames15)

	var names16 []string = make([]string, 3, 7)
	names16[0] = "Kayak"
	names16[1] = "Lifejacket"
	names16[2] = "Paddle"
	fmt.Println("Tableau : ", names16)
	fmt.Println("Longueur : ", len(names16))
	fmt.Println("Capacité : ", cap(names16))

	var names17 []string = make([]string, 3, 7)
	names17[0] = "Kayak"
	names17[1] = "Lifejacket"
	names17[2] = "Paddle"
	moreName := []string{"Hat Gloves"}
	appendedNames17 := append(names17, moreName...)
	fmt.Println("AppendedNames : ", appendedNames17)

	var products [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames := products[1:3]
	allNames := products[:]
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))
	someNames = append(someNames, "Gloves")
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))
	someNames = append(someNames, "Boots")
	fmt.Println("SomeNames : ", someNames)
	fmt.Println("SomeNames len : ", len(someNames), " cap : ", cap(someNames))
	fmt.Println("AllNames : ", allNames)
	fmt.Println("AllNames len : ", len(allNames), " cap : ", cap(allNames))

	var products1 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	someNames1 := products1[1:3:3]
	allNames1 := products1[:]
	someNames1 = append(someNames1, "Gloves")
	fmt.Println("SomeNames : ", someNames1)
	fmt.Println("SomeNames len : ", len(someNames1), " cap : ", cap(someNames1))
	fmt.Println("AllNames : ", allNames1)
	fmt.Println("AllNames len : ", len(allNames1), " cap : ", cap(allNames1))

	var products2 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames2 := products2[1:]
	someNames2 := make([]string, 2)
	copy(someNames2, allNames2)
	fmt.Println("SomeNames2 : ", someNames2)
	fmt.Println("AllNames2 : ", allNames2)

	var products3 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames3 := products3[1:]
	var someNames3 []string // Variable non initialisée : la copie n'est pas effective
	copy(someNames3, allNames3)
	fmt.Println("SomeNames3 : ", someNames3)
	fmt.Println("AllNames3 : ", allNames3)

	var products4 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	allNames4 := products4[1:]
	someNames4 := []string{"Boots", "Canoe"}
	copy(someNames4[1:], allNames4[2:3]) // Le tableau source allNames4 sera copié à partir de la position 2 et
	// La copie va être positionné à partir de la position 1 du tableau destination someNames4
	fmt.Println("SomeNames4 : ", someNames4)
	fmt.Println("AllNames4 : ", allNames4)

	var products5 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts := []string{"Canoe", "Boots"}
	copy(products5[:], replacementProducts)
	fmt.Println("products5 : ", products5) // La copie sera effective uniquement sur les deux premiers éléments du tableau products5

	var products6 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	replacementProducts1 := []string{"Canoe", "Boots"}
	copy(products6[0:1], replacementProducts1)
	fmt.Println("products6 : ", products6) // La copie sera effective uniquement sur le premier élément du tableau products6

	var products7 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	deleted := append(products7[:2], products7[3:]...)
	fmt.Println("Deleted : ", deleted) // On forme le tableau deleted en ajoutant au deux premiers éléments (0 et 1),
	// l'élément numéro 3 : d'où cela supprime l'élément numéro 2

	var products8 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	for index, value := range products8[2:] {
		fmt.Println("Index : ", index, " - Value: ", value)
	}

	var products9 [4]string = [4]string{"Kayak", "Lifejacket", "Paddle", "Hat"}
	sort.Strings(products9[:])
	for index, value := range products9 {
		fmt.Println("Index : ", index, " - Value: ", value)
	}

	products10 := products9
	fmt.Println("Equal : ", reflect.DeepEqual(products9, products10)) // Comparaison de tableau avec la fonction DeepEqual du package reflect

	fmt.Println("Travailler avec Maps : tableau associatif clé-valeur")
	var products11 map[string]float64 = make(map[string]float64, 10) // string représente le type de la clé et float64 le type de la valeur
	products11["Kayak"] = 279
	products11["Lifejacket"] = 48.95
	fmt.Println("Map size: ", len(products11))
	fmt.Println("Price: ", products11["Kayak"])
	fmt.Println("Price: ", products11["Hat"])

	var products12 map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Map size: ", len(products12))
	fmt.Println("Price: ", products12["Kayak"])
	fmt.Println("Price: ", products12["Hat"])

	var products13 map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	value, ok := products13["Hat"]
	if ok {
		fmt.Println("Stored value: ", value)
	} else {
		fmt.Println("No stored value")
	}
	delete(products13, "Hat") // Supprimer un élément du tableau associatif à partir de sa clé
	if value1, ok1 := products13["Hat"]; ok1 {
		fmt.Println("Stored value: ", value1)
	} else {
		fmt.Println("No stored value")
	}

	for key, value := range products13 {
		fmt.Println("Key : ", key, " - Value: ", value)
	}

	var products14 map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	var keys []string = make([]string, 0, len(products14))
	for key := range products14 {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("Key : ", key, " - Valeur : ", products14[key])
	}

	var price string = "$48.95"
	var currencyByte byte = price[0] // Byte est l'alias à uint8
	var currencyString string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("Currency Byte : ", currencyByte)
	fmt.Println("Currency String : ", currencyString)
	fmt.Println("Length : ", len(price))
	if parseErr == nil {
		fmt.Println("Amount : ", amount)
	} else {
		fmt.Println("Parse Error : ", parseErr)
	}

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

	var price3 string = "€48.95"
	for index, char := range []byte(price3) {
		// Le symbole € en byte c'est 3 nombres : 226, 130 et 172
		fmt.Println("Index : ", index, " - value : ", char)
	}
}
