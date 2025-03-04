package game_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"game"
)

func TestListItems_GivesCorrectResultForInput(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input []string
		want  string
	}

	cases := []testCase{
		{
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
	}

	for _, tc := range cases {
		got := game.ListItems(tc.input)
		if tc.want != got {
			t.Errorf(cmp.Diff(tc.want, got))
		}
	}
}

// Chaque fois que nous voulons effectuer la même opération de manière répétée, simplement avec des données différentes à chaque fois,
// nous pouvons exprimer cette idée à l'aide d'une boucle au niveau de la même fonction de test.
// C'est pourquoi le code de test a été refactorisé pour prendre en charge facilement de nouvelles entrées de données à tester.
