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
		{
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
	}

	for _, tc := range cases {
		got := game.ListItems(tc.input)
		if tc.want != got {
			t.Errorf(cmp.Diff(tc.want, got))
		}
	}
}

// Adoption de la règle « nouveau comportement, nouveau test ».
// Chaque fois que nous pensons à un nouveau comportement que nous souhaitons,
// nous devons écrire un test pour celui-ci, ou au moins étendre un test de réussite existant pour qu’il échoue dans
// le cas qui nous intéresse.
// Cas: Gestion du cas où nous avons un slice d'entrées à 2 éléments
