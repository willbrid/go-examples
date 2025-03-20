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

	cases := map[string]testCase{
		"no items": {
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
		"one item": {
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
		"two items": {
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		"three items": {
			input: []string{},
			want:  "",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := game.ListItems(tc.input)
			if tc.want != got {
				t.Error(cmp.Diff(tc.want, got))
			}
		})
	}
}

// Nous voyons que le test parent est TestListItems et nous voyons la fonction t.Run qui permet de créer des sous-tests via son 2ème argument.
// Le nom d'un sous-test décrit souvent les entrées, mais nous pouvons l'utiliser pour ajouter des informations d'étiquette utiles.
// Les sous-tests nommés sont mieux organisés sous forme de map que de slice.
