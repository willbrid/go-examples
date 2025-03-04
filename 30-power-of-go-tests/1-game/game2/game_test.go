package game_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"game"
)

func TestListItems_GivesCorrectResultForInput(t *testing.T) {
	t.Parallel()
	input := []string{
		"a battery",
		"a key",
		"a tourist map",
	}
	want := "You can see here a battery, a key, and a tourist map."
	got := game.ListItems(input)
	if want != got {
		t.Errorf(cmp.Diff(want, got))
	}
}

// Utilisation du package go-cmp - https://pkg.go.dev/github.com/google/go-cmp/cmp
// Installation : go get -u github.com/google/go-cmp/cmp

// Le package cmp détermine l'égalité des valeurs.
// Ce package est destiné à être une alternative plus puissante et plus sûre à reflect.DeepEqual pour comparer si deux valeurs sont
// sémantiquement égales. Il est destiné à être utilisé uniquement dans les tests, car les performances ne sont pas un objectif et il
// peut paniquer s'il ne peut pas comparer les valeurs. Sa propension à paniquer signifie qu'il n'est pas adapté aux environnements de
// production où une panique intempestive peut être fatale.
