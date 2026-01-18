package main

import "fmt"

/**
Les fonctions dont le nom commence par `Benchmark`, suivi d'un terme commençant par une lettre majuscule sont des benchmarks, dont
l'exécution est temporisée. Les fonctions benchmark reçoivent un pointeur vers la structure `testing.B`, qui définit le champ :
- `N` : ce champ int spécifie le nombre de fois que la fonction benchmark doit exécuter le code à mesurer.

La valeur de `N` est utilisée dans une boucle for au sein de la fonction benchmark pour répéter le code dont les performances sont mesurées.
Les outils benchmark peuvent appeler la fonction benchmark à plusieurs reprises, en utilisant différentes valeurs de N,
pour établir une mesure stable.

Pour effectuer le benchmark, nous exécutons la commande : `go test -bench . -run notest`

La période qui suit l'argument `-bench` entraîne l'exécution de tous les tests de performance découverts par l'outil `go test`.
Le point peut être remplacé par une expression régulière pour sélectionner des repères spécifiques. Par défaut, les tests unitaires
sont également effectués, mais nous avons utilisé l'argument `-run` pour spécifier une valeur qui ne correspondra à aucun des
noms de fonction de test dans le projet, avec pour résultat que seuls les tests de performance seront effectués.

Le nom de la fonction de référence est suivi du nombre de processeurs ou de cœurs.
Le champ suivant indique la valeur de `N` qui a été transmise à la fonction de référence pour générer ces résultats.
La valeur finale rend compte de la durée, en nanosecondes, nécessaire pour effectuer chaque itération de la boucle de référence.

Pour chaque itération de la boucle `for`, la fonction BenchmarkSort doit générer des données aléatoires, et le temps nécessaire
pour produire ces données est inclus dans les résultats du benchmark. La classe `testing.B` définit les méthodes utilisées pour
contrôler la minuterie utilisée pour l'analyse comparative.
- StopTimer() : cette méthode arrête le chronomètre.
- StartTimer() : cette méthode démarre le minuteur.
- ResetTimer() : cette méthode réinitialise la minuterie.

La méthode `ResetTimer` est utile lorsqu'un benchmark nécessite une configuration initiale, et les autres méthodes sont utiles lorsqu'il y a
une surcharge associée à chaque activité de benchmark.

Une fonction de benchmark peut exécuter des sous-benchmarks, tout comme une fonction de test peut exécuter des sous-tests.
La méthode utilisée pour exécuter un sous-benchmark :
- Run(name, func) : l'appel de cette méthode exécute la fonction spécifiée en tant que sous-benchmark.
                    La méthode se bloque pendant l'exécution du benchmark.
**/

func main() {
	nums := []int{100, 20, 1, 7, 84}
	sorted, total := SortAndTotal(nums)
	fmt.Println("Sorted Data:", sorted)
	fmt.Println("Total:", total)
}
