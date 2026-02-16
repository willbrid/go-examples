package main

import (
	"fmt"
	"strings"
)

/**
Le package `strings` fournit des fonctions permettant de modifier le contenu des chaînes de caractères.

`Replace(s, old, new, n)` : Cette fonction modifie la chaîne `s` en remplaçant les occurrences de `old` par `new`.
Le nombre maximal d'occurrences à remplacer est spécifié par l'argument entier `n`.

`ReplaceAll(s, old, new)` : Cette fonction modifie la chaîne `s` en remplaçant toutes les occurrences de `old` par `new`.
Contrairement à la fonction `Replace`, le nombre d'occurrences à remplacer est illimité.

`Map(func, s)` : Cette fonction génère une chaîne en appelant la fonction personnalisée pour chaque caractère de la chaîne `s` et
en concaténant les résultats. Si la fonction produit une valeur négative, le caractère courant est ignoré.

`Replace(s)` : Cette méthode renvoie une chaîne de caractères après application de toutes les substitutions spécifiées dans le constructeur
à la chaîne s. Une méthode de `*strings.Replacer`

`WriteString(writer, s)` : Cette méthode permet d'effectuer les substitutions spécifiées dans le constructeur et d'écrire le résultat
dans un objet io.Writer.

Une fonction constructeur nommée `NewReplacer` est utilisée pour créer un `Replacer` et accepte des paires d'arguments qui spécifient les
sous-chaînes et leurs remplacements.
**/

func main() {
	text := "It was a boat. A small boat."
	// Cette fonction modifie la chaîne text en remplaçant les occurrences de la chaîne "boat" par la chaîne "canoe".
	// Le nombre maximum d'occurrences qui seront remplacées est 1.
	replaceResult := strings.Replace(text, "boat", "canoe", 1)
	// // Cette fonction modifie la chaîne text en remplaçant toute les occurrences de la chaîne "boat" par la chaîne "truck".
	replaceAllResult := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace :", replaceResult)
	fmt.Println("ReplaceAll :", replaceAllResult)
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}

		return r
	}
	// Cette fonction se base sur la fonction personalisée mapper pour remplacer toutes les caractères 'b' par 'c'
	mapResult := strings.Map(mapper, text)
	fmt.Println("Map :", mapResult)

	text1 := "It was a boat. A small boat."
	// Le contructeur strings.NewReplacer permet de définir des pairs d'arguments donc l'élément gauche de la pair sera remplacé par l'élément droit
	// dans notre cas : "boat" remplacé par "kayak" et "small" remplacé par "huge"
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	// Cette méthode retourne une chaîne pour laquelle tous les remplacements spécifiés avec le constructeur ont été effectués sur la chaîne text1.
	replaced := replacer.Replace(text1)
	fmt.Println("Replaced : ", replaced)
}
