package main

import (
	"fmt"
	"regexp"
)

/**
Cet ensemble de méthodes RegExp est utilisé pour remplacer les sous-chaînes correspondant à une expression régulière.

`ReplaceAllString(s, template)` : Cette méthode remplace la portion correspondante de la chaîne `s` par le modèle spécifié,
qui est développé avant d'être inclus dans le résultat afin d'y intégrer les sous-expressions.

`ReplaceAllLiteralString(s, sub)` : Cette méthode remplace la portion correspondante de la chaîne `s` par le contenu spécifié, qui est inclus
dans le résultat sans être développé pour les sous-expressions.

`ReplaceAllStringFunc(s, func)` : Cette méthode remplace la portion correspondante de la chaîne `s` par le résultat produit par la fonction spécifiée.
**/

func main() {
	description := "Kayak. A boat for one person."

	/**
	La méthode `ReplaceAllString` est utilisée pour remplacer la portion d'une chaîne de caractères correspondant à l'expression régulière par
	un modèle, qui peut faire référence à des sous-expressions.
	Notons que le modèle n'est responsable que d'une partie du résultat de la méthode ReplaceAl`lString. La première partie de
	la chaîne description (le mot « Kayak », suivi d'un point et d'un espace) n'est pas reconnue par l'expression régulière et est donc
	incluse dans le résultat sans modification.

	Note: Utilisons la méthode `ReplaceAllLiteralString` si nous souhaitons remplacer du contenu sans que la nouvelle sous-chaîne
	soit interprétée pour les sous-expressions.
	**/
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	template := "(type: ${type}, capacity: ${capacity})"
	replaced1 := pattern.ReplaceAllString(description, template)
	replaced2 := pattern.ReplaceAllLiteralString(description, "Good")
	fmt.Println("ReplaceAllString :", replaced1)
	fmt.Println("ReplaceAllLiteralString :", replaced2)

	/**
	La méthode `ReplaceAllStringFunc` remplace la section correspondante d'une chaîne de caractères par un contenu généré par une fonction.
	**/
	pattern1 := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	replaced3 := pattern1.ReplaceAllStringFunc(description, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println("ReplaceAllStringFunc :", replaced3)
}
