package main

import (
	"fmt"
	"sort"
	"testing"
)

/**
Le fichier de test utilise le mot-clé package pour spécifier le package main. Comme les tests sont écrits en Go standard, cela signifie que les tests
de ce fichier ont accès à toutes les fonctionnalités définies dans le package principal, y compris celles qui ne sont pas exportées en dehors du package.
Si nous souhaitons écrire des tests qui n'ont accès qu'aux fonctionnalités exportées, nous pouvons utiliser l'instruction package pour spécifier le
package main_test. Le suffixe _test ne causera pas de problèmes de compilation et permet d'écrire des tests qui n'ont accès qu'aux fonctionnalités
exportées du package en cours de test.

ÉCRITURE DE mocks POUR LES TESTS UNITAIRES
La seule façon de créer des implémentations mock pour les tests unitaires est de créer des implémentations d'interface, qui permettent de définir
des méthodes personnalisées qui produisent les résultats requis pour un test. Si nous souhaitons utiliser des mocks pour nos tests unitaires,
nous devons écrire nos API afin qu'elles acceptent les types d'interface.
Mais même si l'utilisation de mocks est limitée aux interfaces, il est généralement possible de créer des valeurs de structure dont les champs se voient
attribuer des valeurs spécifiques que nous pouvons tester.
**/

type SumTest struct {
	testValues     []int
	expectedResult int
}

/*
*
Le test a appelé la fonction sumAndTotal avec un ensemble de valeurs et a comparé le résultat au résultat attendu à l'aide d'un opérateur de
comparaison Go standard. Si le résultat n'est pas égal à la valeur attendue, la méthode Fatalf est appelée, qui signale l'échec du test et arrête
l'exécution de toutes les instructions restantes dans le test unitaire (bien qu'il n'y ait pas d'instructions restantes dans cet exemple).
*
*/
func TestSum(t *testing.T) {
	testValues := []int{10, 20, 30}
	_, sum := sortAndTotal(testValues)
	expected := 60
	if sum != expected {
		t.Fatalf("Expected %v, Got %v", expected, sum)
	}
}

func TestSort(t *testing.T) {
	testValues := []int{1, 279, 48, 12, 3}
	sorted, _ := sortAndTotal(testValues)
	if !sort.IntsAreSorted(sorted) { // IntsAreSorted signale si la tranche x est triée par ordre croissant.
		t.Fatalf("Unsorted data %v", sorted)
	}
}

func TestSortWithRun(t *testing.T) {
	slices := [][]int{
		{1, 279, 48, 12, 3},
		{-10, 0, -10},
		{1, 2, 3, 4, 5, 6, 7},
		{1},
	}

	for index, data := range slices {
		t.Run(fmt.Sprintf("Sort  #%v", index), func(subT *testing.T) {
			sorted, _ := sortAndTotal(data)
			if !sort.IntsAreSorted(sorted) { // IntsAreSorted signale si la tranche x est triée par ordre croissant.
				subT.Fatalf("Unsorted data %v", sorted)
			}
		})
	}
}

func TestSumWithSkip(t *testing.T) {
	testVals := []SumTest{
		{testValues: []int{10, 20, 30}, expectedResult: 10},
		{testValues: []int{-10, 0, -10}, expectedResult: -20},
		{testValues: []int{-10, 0, -10}, expectedResult: -20},
	}
	for index, testVal := range testVals {
		t.Run(fmt.Sprintf("Sum  #%v", index), func(subT *testing.T) {
			if t.Failed() {
				subT.SkipNow()
			}

			_, sum := sortAndTotal(testVal.testValues)
			if sum != testVal.expectedResult {
				subT.Fatalf("Expected %v, Got %v", testVal.expectedResult, sum)
			}
		})
	}
}
