package main

import "fmt"

func main() {
	/**
	La meilleure façon de concevoir les tranches est de les considérer comme un tableau de longueur variable, car elles sont utiles
	lorsque nous ne connaissons pas le nombre de valeurs à stocker ou lorsque ce nombre évolue dans le temps.
	Une façon de définir une tranche est d'utiliser la fonction intégrée `make`. La fonction make accepte des arguments qui spécifient le type
	et la longueur de la tranche.
	**/

	names := make([]string, 3)
	names[0] = "kayak"
	names[1] = "lifejacket"
	names[2] = "paddle"
	fmt.Println("names :", names)

	// Syntaxe littérale de tranche
	name1s := []string{"kayak", "lifejacket", "paddle"}
	fmt.Println("name1s :", name1s)

	/**
	La combinaison du type et de la longueur d'une tranche permet de créer un tableau, qui sert de zone de stockage pour cette tranche.
	Une tranche est une structure de données qui contient trois valeurs : un pointeur vers le tableau, sa longueur et sa capacité.
	La longueur d'une tranche correspond au nombre d'éléments qu'elle peut stocker, et sa capacité correspond au nombre total
	d'éléments pouvant être stockés dans le tableau.
	Les tranches prennent en charge la notation d'index de type tableau, ce qui permet d'accéder aux éléments du tableau sous-jacent.
	**/

	/**
	La fonction intégrée `append` accepte une tranche et un ou plusieurs éléments à ajouter à la tranche, séparés par des virgules.
	La fonction `append` crée un tableau suffisamment grand pour contenir les nouveaux éléments, copie le tableau existant et
	y ajoute les nouvelles valeurs. Le résultat de la fonction `append` est une tranche qui est mappée sur le nouveau tableau.
	**/
	name2s := []string{"kayak", "lifejacket", "paddle"}
	name2s = append(name2s, "Hat", "Gloves")
	fmt.Println("name2s :", name2s)

	// La tranche originale et son tableau sous-jacent existent toujours et peuvent être utilisés.
	name3s := []string{"kayak", "lifejacket", "paddle"}
	appendedNames := append(name3s, "Hat", "Gloves")
	name3s[0] = "Canoe"
	fmt.Println("name3s :", name3s)
	fmt.Println("appendedNames :", appendedNames)

	/**
	Comme précisé plus haut, Les tranches ont une longueur et une capacité. La longueur d'une tranche correspond au nombre de valeurs
	qu'elle peut contenir, tandis que la capacité représente le nombre d'éléments pouvant être stockés dans le tableau sous-jacent
	avant que la tranche ne doive être redimensionnée et un nouveau tableau créé. La capacité est toujours au moins égale à la longueur,
	mais peut être supérieure si de la capacité supplémentaire a été allouée avec la fonction `make`.

	Les fonctions intégrées `len` et `cap` renvoient la longueur et la capacité d'une tranche.

	Si nous définissons une variable de type tranche sans l'initialiser, la tranche résultante aura une longueur et une capacité nulles,
	ce qui provoquera une erreur lors de l'ajout d'un élément.
	**/
	name4s := make([]string, 3, 6)
	name4s[0] = "kayak"
	name4s[1] = "lifejacket"
	name4s[2] = "paddle"
	fmt.Println("name4s len :", len(name4s))
	fmt.Println("name4s cap :", cap(name4s))

	/**
	Le tableau sous-jacent n'est pas remplacé lorsque la fonction `append` est appelée sur une tranche ayant une capacité suffisante
	pour accueillir les nouveaux éléments.
	Dans cet exemple, la fonction `append` a pour résultat une tranche dont la longueur a augmenté, mais qui repose toujours sur
	le même tableau sous-jacent. La tranche d'origine existe toujours et repose sur le même tableau, ce qui a pour effet
	qu'il existe désormais deux vues d'un même tableau.
	**/
	name5s := make([]string, 3, 6)
	name5s[0] = "kayak"
	name5s[1] = "lifejacket"
	name5s[2] = "paddle"
	appendedName1s := append(name5s, "Hat", "Gloves")
	name5s[0] = "Canoe"
	fmt.Println("name5s :", name5s)
	fmt.Println("appendedName1s :", appendedName1s)
}
