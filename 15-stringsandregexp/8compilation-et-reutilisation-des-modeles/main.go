package main

/**
Le package regexp offre une prise en charge des expressions régulières, qui permettent de trouver des motifs complexes dans les chaînes de caractères.

Match(pattern, b) : Cette fonction renvoie un booléen indiquant si le motif est trouvé par la tranche d'octets b.

MatchString(pattern, s) : Cette fonction renvoie un booléen indiquant si le motif est trouvé par la chaîne de caractères s.
Cette méthode est le moyen le plus simple de déterminer si une chaîne de caractères correspond à une expression régulière.
La fonction MatchString accepte une expression régulière et la chaîne à rechercher. Elle renvoie un booléen : vrai en cas de correspondance,
nul en cas d’erreur. Les erreurs d’expression régulière surviennent généralement lorsque le modèle ne peut pas être traité.

Compile(pattern) : Cette fonction renvoie une expression régulière permettant d'effectuer des correspondances répétées avec le motif spécifié.
Cette méthode est plus efficace car le modèle n'a besoin d'être compilé qu'une seule fois. Le résultat de la fonction `Compile` est une instance du
type `RegExp`.

MustCompile(pattern) : Cette fonction offre les mêmes fonctionnalités que Compile, mais génère une erreur si le motif spécifié ne peut pas être compilé.


Le type `RegExp` fournit également des méthodes permettant de traiter des tranches d'octets et des méthodes gérant les readers, qui font partie
de la prise en charge des E/S en Go. ci-dessous les méhodes de `RegExp` :

MatchString(s) : Cette méthode renvoie vrai si la chaîne s correspond au motif compilé.

FindStringIndex(s) : Cette méthode renvoie un tableau d'entiers contenant l'emplacement de la première occurrence du motif compilé dans la chaîne s.
Un résultat nul indique qu'aucune occurrence n'a été trouvée.

FindAllStringIndex(s, max) : Cette méthode renvoie un tableau de tableaux d'entiers contenant l'emplacement de toutes les occurrences du motif compilé
dans la chaîne s. Un résultat nul indique qu'aucune occurrence n'a été trouvée.

FindString(s) : Cette méthode renvoie une chaîne contenant la première occurrence du motif compilé dans la chaîne s. Une chaîne vide est renvoyée si
aucune occurrence n'est trouvée.

FindAllString(s, max) : Cette méthode renvoie un tableau de chaînes contenant toutes les occurrences du motif compilé dans la chaîne s.
L'argument entier max spécifie le nombre maximal d'occurrences, -1 indiquant une limite nulle. Un résultat nul est renvoyé si
aucune occurrence n'est trouvée.

Split(s, max) : Cette méthode divise la chaîne s en utilisant les correspondances du modèle compilé comme séparateurs et renvoie une tranche
contenant les sous-chaînes divisées.
**/

import (
	"fmt"
	"regexp"
)

func getSubstring(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func main() {
	description := "A boat for one person"

	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match :", match)
	} else {
		fmt.Println("Error :", err)
	}

	pattern, compileErr := regexp.Compile("[A-z]oat")
	question := "Is that a goat ?"
	preference := "I like oats"
	if compileErr == nil {
		fmt.Println("Description :", pattern.MatchString(description))
		fmt.Println("Question :", pattern.MatchString(question))
		fmt.Println("Preference :", pattern.MatchString(preference))
	} else {
		fmt.Println("Error :", compileErr)
	}

	/**
	L'expression régulière trouvera deux correspondances avec la chaîne de description.

	La méthode `FindStringIndex` ne renvoie que la première correspondance, en parcourant la chaîne de gauche à droite.
	La correspondance est exprimée sous forme de tranche d'entiers, où la première valeur indique la position de départ de la correspondance
	dans la chaîne, et la seconde le nombre de caractères correspondants.

	La méthode `FindAllStringIndex` renvoie plusieurs correspondances et est appelée avec l'argument `-1`, indiquant que toutes les correspondances
	doivent être renvoyées. Les correspondances sont renvoyées dans une tranche de tranches d'entiers
	(chaque valeur de la tranche résultante est une tranche d'entiers), chacune décrivant une correspondance. Les indices sont utilisés pour
	extraire des régions de la chaîne à l'aide de la fonction `getSubstring`.
	**/
	description1 := "Kayak. A boat for one person."
	pattern1 := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	firstIndex := pattern1.FindStringIndex(description1)
	allIndices := pattern1.FindAllStringIndex(description1, -1)
	fmt.Println("Result First index :", firstIndex)
	fmt.Println("First index ", firstIndex[0], "-", firstIndex[1], "=", getSubstring(description, firstIndex))
	fmt.Println("Result All indices :", allIndices)
	for i, idx := range allIndices {
		fmt.Println("Index ", i, "=", idx[0], "-", idx[1], "=", getSubstring(description, idx))
	}

	/**
	Si vous n’avez pas besoin de connaître l’emplacement des correspondances, les méthodes FindString et FindAllString sont plus utiles car
	leurs résultats sont les sous-chaînes correspondant à l’expression régulière.
	**/
	description2 := "Kayak. A boat for one person."
	pattern2 := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	firstMatch := pattern2.FindString(description2)
	allMatches := pattern2.FindAllString(description2, -1)
	fmt.Println("First match :", firstMatch)
	for i, m := range allMatches {
		fmt.Println("Match :", i, "=", m)
	}

	/**
	La méthode Split divise une chaîne de caractères en fonction des correspondances trouvées par une expression régulière.
	**/
	description3 := "Kayak. A boat for one person."
	pattern3 := regexp.MustCompile(" |boat|one")
	splits := pattern3.Split(description3, -1)
	for _, s := range splits {
		if s != "" {
			fmt.Println("Split substring :", s)
		}
	}
}
