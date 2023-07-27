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
