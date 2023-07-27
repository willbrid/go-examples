package main

import (
	"fmt"
	"sort"
)

/**
Les tests unitaires sont définis dans des fichiers dont le nom se termine par _test.go

La bibliothèque standard Go prend en charge l'écriture de tests unitaires via le package testing. Les tests unitaires sont exprimés sous forme de
fonctions dont le nom commence par Test, suivi d'un terme commençant par une lettre majuscule, comme TestSum.
(La lettre majuscule est importante car les outils de test ne reconnaîtront pas un nom de fonction tel que Testsum comme test unitaire.)

Les fonctions de test unitaire reçoivent un pointeur vers une classe T (testing.T), qui définit les méthodes de gestion des tests et de rapport des résultats
des tests. Les tests Go ne reposent pas sur des assertions et sont écrits à l'aide d'instructions de code standard. Tout ce dont les outils de test
se soucient, c'est de savoir si le test échoue, ce qui est signalé à l'aide des méthodes :
- Log(...vals) : cette méthode écrit les valeurs spécifiées dans le journal des erreurs de test.
- Logf(template, ...vals) : cette méthode utilise le modèle et les valeurs spécifiés pour écrire un message dans le journal des erreurs de test.
- Fail() : l'appel de cette méthode marque le test comme ayant échoué mais poursuit l'exécution du test.
- FailNow() : l'appel de cette méthode marque le test comme ayant échoué et arrête l'exécution du test.
- Failed() ; cette méthode renvoie true si le test a échoué.
- Error(...errs) : l'appel de cette méthode équivaut à appeler la méthode Log, suivie de la méthode Fail.
- Errorf(template, ...vals) : l'appel de cette méthode équivaut à appeler la méthode Logf, suivie de la méthode Fail.
- Fatal(...vals) : l'appel de cette méthode équivaut à appeler la méthode Log, suivie de la méthode FailNow.
- Fatalf(template, ...vals) : l'appel de cette méthode équivaut à appeler la méthode Logf, suivie de la méthode FailNow.

Un fichier de test peut contenir plusieurs tests.

La commande <go test> ne rapporte aucun détail par défaut, mais plus d'informations peuvent être générées en exécutant la commande : <go test -v>

La commande go test peut être utilisée pour exécuter des tests sélectionnés par leur nom : <go test -v -run "um">
Les tests sont sélectionnés avec une expression régulière. Le seul test dont le nom correspond à l'expression est <TestSum>.

La classe T (testing.T) fournit également un ensemble de méthodes pour gérer l'exécution des tests :
- Run(name, func) : l'appel de cette méthode exécute la fonction spécifiée en tant que sous-test. La méthode se bloque pendant que le test est exécuté
                    dans sa propre goroutine et renvoie un booléen qui indique si le test a réussi.
					La méthode Run est utilisée pour exécuter un sous-test, ce qui est un moyen pratique d'exécuter une série de tests liés à partir
					d'une seule fonction.
- SkipNow() : l'appel de cette méthode arrête l'exécution du test et le marque comme ignoré.
- Skip(...args) : cette méthode équivaut à appeler la méthode Log, suivie de la méthode SkipNow.
- Skipf(template, ...args) : cette méthode équivaut à appeler la méthode Logf, suivie de la méthode SkipNow.
- Skipped() : cette méthode renvoie true si le test a été ignoré.

Les tests peuvent être ignorés, ce qui peut être utile lorsqu'un échec d'un test signifie qu'il est inutile d'effectuer des tests connexes.


# Code d'analyse comparative
Les fonctions dont le nom commence par Benchmark, suivi d'un terme commençant par une lettre majuscule sont des benchmarks, dont l'exécution est temporisée.
Les fonctions de benchmark reçoivent un pointeur vers la classe testing.B, qui définit le champ :
- N : ce champ int spécifie le nombre de fois que la fonction de benchmark doit exécuter le code à mesurer.

La valeur de N est utilisée dans une boucle for au sein de la fonction benchmark pour répéter le code dont les performances sont mesurées.
Les outils benchmark peuvent appeler la fonction benchmark à plusieurs reprises, en utilisant différentes valeurs de N, pour établir une mesure stable.

Pour effectuer le benchmark, nous exécutons la commande : go test -bench . -run notest

La période qui suit l'argument -bench entraîne l'exécution de tous les tests de performance découverts par l'outil de test go. Le point peut être
remplacé par une expression régulière pour sélectionner des repères spécifiques. Par défaut, les tests unitaires sont également effectués, mais nous avons
utilisé l'argument -run pour spécifier une valeur qui ne correspondra à aucun des noms de fonction de test dans le projet, avec pour résultat
que seuls les tests de performance seront effectués.

Le nom de la fonction de référence est suivi du nombre de processeurs ou de cœurs.
Le champ suivant indique la valeur de N qui a été transmise à la fonction de référence pour générer ces résultats.
La valeur finale rend compte de la durée, en nanosecondes, nécessaire pour effectuer chaque itération de la boucle de référence.

Pour chaque itération de la boucle for, la fonction BenchmarkSort doit générer des données aléatoires, et le temps nécessaire pour produire ces données
est inclus dans les résultats du benchmark. La classe testing.B définit les méthodes utilisées pour contrôler la minuterie utilisée pour l'analyse comparative.
- StopTimer() : cette méthode arrête le chronomètre.
- StartTimer() : cette méthode démarre le minuteur.
- ResetTimer() : cette méthode réinitialise la minuterie.

La méthode ResetTimer est utile lorsqu'un benchmark nécessite une configuration initiale, et les autres méthodes sont utiles lorsqu'il y a
une surcharge associée à chaque activité de benchmark.

Une fonction de benchmark peut exécuter des sous-benchmarks, tout comme une fonction de test peut exécuter des sous-tests.
La méthode utilisée pour exécuter un sous-benchmark :
- Run(name, func) : l'appel de cette méthode exécute la fonction spécifiée en tant que sous-benchmark. La méthode se bloque pendant l'exécution du benchmark.
**/

func sortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, val := range sorted {
		total += val
		// total++ // Ici on fixe l'erreur qui a conduit à un échec de test
	}
	return
}

func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := sortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)
}
