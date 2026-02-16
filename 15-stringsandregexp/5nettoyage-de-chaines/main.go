package main

import (
	"fmt"
	"strings"
)

/**
Le processus de nettoyage supprime les caractères de début et de fin d'une chaîne de caractères et est le plus souvent utilisé pour
supprimer les espaces.

TrimSpace(s) : Cette fonction renvoie la chaîne s sans les espaces de début ni de fin.

Trim(s, set) : Cette fonction renvoie une chaîne dont les caractères de début et de fin contenus dans le set sont supprimés.

TrimLeft(s, set) : Cette fonction renvoie la chaîne s sans aucun caractère de début contenu dans le set.
Elle accepte n'importe quel caractère spécifié ; utilisez la fonction TrimPrefix pour supprimer une sous-chaîne complète.

TrimRight(s, set) : Cette fonction renvoie la chaîne s sans aucun caractère de fin contenu dans le set.
Elle accepte n'importe quel caractère spécifié ; utilisez la fonction TrimSuffix pour supprimer une sous-chaîne complète.

TrimPrefix(s, prefix) : Cette fonction renvoie la chaîne s après suppression du préfixe spécifié.
Elle supprime l'intégralité du préfixe ; utilisez la fonction TrimLeft pour supprimer des caractères d'un ensemble.

TrimSuffix(s, suffix) : Cette fonction renvoie la chaîne s après suppression du suffixe spécifié.
Cette fonction supprime la chaîne de suffixe complète. Utilisez la fonction TrimRight pour supprimer des caractères d'un ensemble.

TrimFunc(s, func) : Cette fonction renvoie la chaîne s dont on a supprimé tout caractère de début ou de fin pour lequel
une fonction personnalisée renvoie vrai.

TrimLeftFunc(s, func) : Cette fonction renvoie la chaîne s dont on a supprimé tout caractère de début pour lequel
une fonction personnalisée renvoie vrai.

TrimRightFunc(s, func) : Cette fonction renvoie la chaîne s dont on a supprimé tout caractère de fin pour lequel
une fonction personnalisée renvoie vrai.
**/

func main() {
	username := " Alice"
	trimSpaceResult := strings.TrimSpace(username)
	fmt.Println("Trimmed :", ">>"+trimSpaceResult+"<<")

	description := "A boat for one person"
	trimResult := strings.Trim(description, "Asno ")
	trimLeftResult := strings.TrimLeft(description, "Asno ")
	trimRightResult := strings.TrimRight(description, "Asno ")
	fmt.Println("Trim :", trimResult)
	fmt.Println("TrimLeft :", trimLeftResult)
	fmt.Println("TrimRight :", trimRightResult)
	trimPrefixResult := strings.TrimPrefix(description, "A boat ")
	trimSuffixResult := strings.TrimSuffix(description, "son")
	fmt.Println("TrimPrefix :", trimPrefixResult)
	fmt.Println("TrimSuffix :", trimSuffixResult)
	// La fonction personnalisée est appelée pour les caractères au début et à la fin de la chaîne,
	// et les caractères seront coupés jusqu'à ce que la fonction renvoie false.
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	trimFuncResult := strings.TrimFunc(description, trimmer)
	fmt.Println("TrimFunc :", trimFuncResult)
}
