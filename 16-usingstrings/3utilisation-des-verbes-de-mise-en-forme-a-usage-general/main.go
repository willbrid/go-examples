package main

import "fmt"

/**
Les verbes à usage général peuvent être utilisés pour afficher n'importe quelle valeur.

%v : Ce verbe affiche le format par défaut de la valeur. En ajoutant un signe plus (+) (%+v), les noms de champs sont inclus lors de l’écriture des
valeurs de structure.

%#v : Ce verbe affiche une valeur dans un format permettant de la recréer dans un fichier de code Go.

%T : Ce verbe affiche le type Go d’une valeur.


Le package `fmt` prend en charge la mise en forme personnalisée des structures via une interface nommée `Stringer`, définie comme suit
type Stringer interface {
    String() string
}

Si nous définissons une méthode `GoString` qui renvoie une chaîne de caractères, alors notre type sera conforme à l'interface `GoStringer`,
qui permet une mise en forme personnalisée pour le verbe %#v.

Lorsque les tableaux et les tranches sont représentés sous forme de chaînes de caractères, le résultat est un ensemble de crochets, à l'intérieur
desquels se trouvent les éléments individuels, comme ceci : [Kayak Lifejacket Paddle]. Notons qu'aucune virgule ne sépare les éléments.
Lorsque les dictionnaires sont représentés sous forme de chaînes de caractères, les paires clé-valeur sont affichées entre crochets,
précédées du mot-clé `map`, comme ceci : map[1:Kayak 2:Lifejacket 3:Paddle]
L'interface `Stringer` peut être utilisée pour modifier le format utilisé pour les types de données personnalisés contenus dans un tableau,
une tranche ou un map.
**/

func Printfln(template string, values ...any) {
	fmt.Printf(template+"\n", values...)
}

/**
La méthode `String` sera automatiquement appelée lorsqu'une représentation sous forme de chaîne de caractères de la valeur
d'une structure (dans cet exemple `Product`) est requise : lors de l'utilisation du verbe %v et %+v.
**/

func (p Product) String() string {
	return fmt.Sprintf("Product: %v, Price: $%4.2f", p.Name, p.Price)
}

func main() {
	/**
	Go possède un format par défaut pour tous les types de données utilisés par le verbe `%v`. Pour les structures, la valeur par défaut liste
	les valeurs des champs entre accolades. On peut modifier ce format par défaut avec un signe plus pour inclure les noms des champs dans le résultat.
	**/
	Printfln("Value : %v", kayak)
	Printfln("Value with fields : %+v", kayak)
	Printfln("GO syntaxe : %#v", kayak)
	Printfln("Type : %T", kayak)
}
