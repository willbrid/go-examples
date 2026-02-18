package main

import (
	"fmt"
	"regexp"
)

/**
Les sous-expressions permettent d'accéder à des parties d'une expression régulière, ce qui peut faciliter l'extraction de sous-chaînes à
l'intérieur d'une région correspondante.

Les méthodes du type Regexp pour les sous-expressions :

FindStringSubmatch(s) : Cette méthode renvoie une tranche contenant la première correspondance trouvée par le motif et le texte des sous-expressions
définies par ce motif.

FindAllStringSubmatch(s, max) : Cette méthode renvoie une tranche contenant toutes les correspondances et le texte des sous-expressions.
L'argument entier permet de spécifier le nombre maximal de correspondances. La valeur -1 indique que toutes les correspondances sont prises en compte.

FindStringSubmatchIndex(s) : Cette méthode est équivalente à FindStringSubmatch, mais renvoie les indices au lieu des sous-chaînes.

FindAllStringSubmatchIndex(s, max) : Cette méthode est équivalente à FindAllStringSubmatch, mais renvoie les indices au lieu des sous-chaînes.

NumSubexp() : Cette méthode renvoie le nombre de sous-expressions.

SubexpIndex(name) : Cette méthode renvoie l'indice de la sous-expression portant le nom spécifié, ou -1 si aucune sous-expression ne correspond.

SubexpNames() : Cette méthode renvoie les noms des sous-expressions, exprimés dans l'ordre où elles sont définies.
**/

func main() {
	description := "Kayak. A boat for one person."

	/**
	Dans cet exemple, le modèle correspond à une structure de phrase spécifique, ce qui me permet d'extraire la partie de la chaîne voulue.
	Cependant, une grande partie de la structure de la phrase est statique, et les deux sections variables du modèle contiennent le contenu recherché.
	La méthode FindString est inefficace dans ce cas, car elle extrait l'intégralité du modèle, y compris les régions statiques.
	**/
	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
	str := pattern.FindString(description)
	fmt.Println("Match :", str)

	/**
	Nous pouvons ajouter des sous-expressions pour identifier les régions de contenu importantes au sein du modèle.
	Les sous-expressions sont indiquées par des parenthèses. Nous avons défini deux sous-expressions, chacune encadrant une portion variable du motif.
	La méthode `FindStringSubmatch` effectue la même opération que `FindString`, mais inclut également dans son résultat les sous-chaînes correspondant
	aux expressions.
	**/
	pattern1 := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	subs := pattern1.FindStringSubmatch(description)
	for _, s := range subs {
		fmt.Println("SubMatch :", s)
	}
}
