package square_test

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"square"
)

func TestSquare_GivesNonNegativeResult(t *testing.T) {
	t.Parallel()

	inputs := rand.Perm(100)

	for _, n := range inputs {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			got := square.Square(n)
			if got < 0 {
				t.Errorf("Square(%d) is negative: %d", n, got)
			}
		})
	}
}

// Bien que nous ne puissions pas prédire le résultat exact si le système est correct, nous pouvons néanmoins identifier certaines
// propriétés qu'il devrait posséder.
// Nous pourrions donc écrire un test appelant Square avec de nombreuses entrées différentes et vérifiant que le résultat
// n'est jamais négatif.
// Cette approche est parfois appelée test basé sur les propriétés.
// On peut aussi considérer que les tests basés sur les propriétés décrivent le comportement du système, non pas en termes de valeurs exactes,
// mais en termes d'invariants : des éléments qui ne changent pas dans le résultat, quelle que soit l'entrée. Dans le cas de Square,
// par exemple, son résultat devrait toujours être positif.
// Les tests aléatoires basés sur les propriétés permettent de résoudre le problème suivant : la fonction ne fonctionne peut-être que
// pour les exemples spécifiques auxquels nous avons pensé. Et même si nous générons les entrées de manière aléatoire, dès que nous
// trouvons une valeur qui déclenche un bug, celle-ci devrait simplement être intégrée à nos tests basés sur les exemples classiques.
// Nous pourrions y parvenir en ajoutant manuellement ces valeurs à l'ensemble des entrées d'un test de table, par exemple.
