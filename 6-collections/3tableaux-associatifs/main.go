package main

import (
	"fmt"
	"sort"
)

func main() {
	/** Travailler avec Maps : tableau associatif (dictionnaire) clé-valeur
	Les dictionnaires sont une structure de données intégrée qui associe des valeurs de données à des clés. Ils sont créées avec la
	fonction intégrée `make`, comme pour les tranches. Le type d'un dictionnaire est spécifié à l'aide du mot-clé `map`,
	suivi du type de clé entre crochets, puis du type de valeur.
	**/
	fmt.Println("Travailler avec Maps : tableau associatif clé-valeur")
	var products map[string]float64 = make(map[string]float64, 10) // string représente le type de la clé et float64 le type de la valeur
	products["Kayak"] = 279
	products["Lifejacket"] = 48.95
	fmt.Println("Map size: ", len(products))
	fmt.Println("Price: ", products["Kayak"])
	fmt.Println("Price: ", products["Hat"])

	// Les dictionnaires peuvent également être définies à l’aide d’une syntaxe littérale.
	var product1s map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
	}
	fmt.Println("Map size: ", len(product1s))
	fmt.Println("Price: ", product1s["Kayak"])
	fmt.Println("Price: ", product1s["Hat"])

	/**
	Les dictionnaires renvoient le zéro pour le type de valeur lors de la lecture d'une clé inexistante. Il peut alors être difficile
	de distinguer une valeur stockée avec le zéro du type de valeur d'une clé inexistante. Pour résoudre ce problème,
	les dictionnaires produisent deux valeurs lors de la lecture d'une valeur :
	--- la première valeur est soit la valeur associée à la clé spécifiée, soit le zéro du type de valeur en l'absence de clé.
	--- la seconde valeur est une valeur booléenne qui est vraie si le dictionnaire contient la clé spécifiée et fausse dans
	    le cas contraire. Cette seconde valeur est généralement affectée à une variable nommée ok, d'où le terme « comma ok ».
	**/
	var product2s map[string]float64 = map[string]float64{
		"Kayak":      279,
		"Lifejacket": 48.95,
		"Hat":        0,
	}
	value, ok := product2s["Hat"]
	if ok {
		fmt.Println("Stored value: ", value)
	} else {
		fmt.Println("No stored value")
	}
	delete(product2s, "Hat") // Supprimer un élément du tableau associatif à partir de sa clé
	if value1, ok1 := product2s["Hat"]; ok1 {
		fmt.Println("Stored value: ", value1)
	} else {
		fmt.Println("No stored value")
	}

	for key, value := range product2s {
		fmt.Println("Key : ", key, " - Value: ", value)
	}

	/**
	Il n'y a aucune garantie que le contenu d'un dictionnaire soit énuméré dans un ordre spécifique. Pour classer les valeurs d'un
	dictionnaire, la meilleure approche consiste à énumérer le dictionnaire, à créer une tranche contenant les clés,
	à trier la tranche, puis à énumérer la tranche pour lire les valeurs du dictionnaire.
	**/
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
}
