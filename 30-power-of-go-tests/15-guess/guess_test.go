package guess_test

import (
	"testing"

	"guess"
)

func FuzzGuess(f *testing.F) {
	f.Fuzz(func(t *testing.T, input int) {
		guess.Guess(input)
	})
}

// Les tests fuzz permettent de détecter des bugs en testant une fonction avec des entrées générées aléatoirement. Ce type de test est
// utile pour explorer un large éventail de cas d'entrée et identifier des comportements inattendus.
// Introduits dans Go 1.18, les tests fuzz peuvent aider à trouver des valeurs spécifiques qui déclenchent des erreurs, comme une panique
// dans une fonction. L'idée est de générer un grand nombre de valeurs aléatoires et d'exécuter la fonction avec chacune d'elles.
// Plutôt que d’écrire manuellement un test pour chaque cas possible, on peut utiliser le package de test standard de Go pour automatiser
// ce processus et augmenter les chances de repérer des bugs cachés.

// En Go, les noms des tests fuzz commencent par le mot « Fuzz ». Et le paramètre qu'ils prennent est un *testing.F, et qui possède de
// nombreuses méthodes similaires.

// Nous appelons la méthode Fuzz du fuzzer, en lui passant une fonction. C'est la fonction que le fuzzer (le testeur de fuzz) appellera à
// plusieurs reprises avec les différentes valeurs qu'il génère. Nous appelons cette fonction la cible fuzz, et sa signature est :
// func(t *testing.T, input int)

// Bien que la fontion cible peut prendre un *testing.T comme premier argument, elle peut également accepter un nombre illimité
// d'autres arguments.
// Ces arguments supplémentaires à la cible fuzz représentent nos entrées, que le fuzzer va générer aléatoirement. Autrement dit, le rôle
// du fuzzer est d'appeler cette fonction (la cible fuzz) avec de nombreuses valeurs différentes en entrée, et d'observer le résultat.

// Exécution de tests en mode fuzzing
// Pour lancer le fuzzing, nous utilisons la commande go test, en ajoutant l'indicateur -fuzz :
// go test -fuzz .
// Remarque : le point après l'indicateur -fuzz est significatif. Tout comme l'indicateur -run utilise une expression régulière
// pour spécifier les tests à exécuter, l'indicateur -fuzz fait de même. Dans ce cas, nous utilisons l'expression régulière «.»,
// qui correspond à tous les tests fuzz.

// Les tests fuzz peuvent aussi agir comme des tests classiques. Lorsqu'ils sont exécutés sans l'option -fuzz, ils n’utilisent pas
// d’entrées aléatoires, mais rejouent les valeurs problématiques déjà découvertes lors de précédents tests de fuzzing.
// Ces valeurs sont stockées dans des fichiers spéciaux sous le dossier testdata/fuzz.
// Ainsi, une fois qu’un bug est détecté par le fuzzer, cette entrée sera testée à chaque exécution de go test,
// sauf si nous supprimons manuellement les fichiers de test dans testdata.
// C'est une fonctionnalité intéressante, car elle permet au fuzzer de créer automatiquement des tests de régression. Si un bug
// devait survenir ultérieurement et entraîner l'échec de la fonction avec la même entrée, go test le détecterait. Chaque fois
// qu'un nouveau cas de test défaillant est détecté par le fuzzer, il est intégré à notre suite de tests standard.
