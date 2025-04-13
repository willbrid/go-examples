package runes

import "unicode/utf8"

// L'opérateur range parcourt une chaîne par runes, et non par octets. Ainsi, à chaque exécution successive de cette boucle,
// r sera la rune suivante de la chaîne, en commençant par la première. Mais dès cette première itération, nous appliquons
// l'instruction return dans le corps de la boucle et renvoyons immédiatement la première rune.
func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}

	return utf8.RuneError
}

/**
// Third implementation : the fuzz test failed
func FirstRune(s string) rune {
	if s == "" {
		return utf8.RuneError
	}

	return rune(s[0])
}

Il est déconseillé de mettre à jour automatiquement les cas de test à partir du comportement actuel, car cela pourrait cacher des bugs.
En revanche, il est utile de générer automatiquement des tests qui échouent, car cela attire l’attention sur un problème réel.
Le fuzzing est un bon moyen de découvrir des bugs en générant des entrées inattendues. Cependant :
- il doit être exécuté manuellement (par exemple en local), car il modifie les cas de test stockés dans le code source.
- il prend du temps, donc pas idéal à intégrer dans l'intégration continue.
- il est particulièrement utile une fois que les tests classiques sont en place.
- un test fuzz continue indéfiniment s’il ne trouve pas d’erreur, ce qui est normal car l’espace d’entrées est très vaste.

// Second implementation : the fuzz test failed
func FirstRune(s string) rune {
	return rune(s[0])
}
Quand le fuzzer démarre, il utilise tous les cœurs disponibles du processeur (par exemple 8 workers) pour générer des entrées aléatoires
en parallèle. Cela augmente l’efficacité du fuzzing.
Lorsqu’une entrée cause une erreur, le fuzzer cherche automatiquement à réduire cette entrée à sa version minimale défaillante.
Ce processus est appelé minimisation : il permet d’éliminer tout ce qui n’est pas essentiel au bug.
Cela facilite le diagnostic du bug pour les développeurs, car l’entrée est plus simple, plus courte et plus ciblée.

// First implementation : the fuzz test failed
func FirstRune(s string) rune {
	return 0
}
**/
