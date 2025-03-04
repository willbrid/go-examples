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
		{
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		{
			input: []string{},
			want:  "",
		},
	}

	for _, tc := range cases {
		got := game.ListItems(tc.input)
		if tc.want != got {
			t.Errorf(cmp.Diff(tc.want, got))
		}
	}
}

// Comme le code est correct et que tous les cas sont passés, nous pouvons maintenant refactoriser le code correct pour le rendre
// plus agréable.
// La définition de « refactoring » consiste à modifier le code sans modifier son comportement de manière significative.
// Étant donné que le test définit tous les comportements que nous considérons comme pertinents, nous pouvons modifier
// le code en toute liberté, en nous fiant au test pour nous indiquer le moment où le code commence à se comporter différemment.
