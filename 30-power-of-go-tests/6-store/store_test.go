package store_test

import (
	"testing"

	"store"

	"github.com/google/go-cmp/cmp"
)

func TestOpen_ReturnsStoreObject(t *testing.T) {
	t.Parallel()

	want := &store.Store{}
	got, err := store.Open("testdata/godlen.txt") // Le chemin correct du fichier est : testdata/golden.txt

	if err != nil {
		t.Fatalf("unexpected error reading golden file: %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestOpen_GivesErrUnopenableForBogusFile(t *testing.T) {
	t.Parallel()

	_, err := store.Open("bogus")
	if err != store.ErrUnopenable {
		t.Errorf("wrong error: %v", err)
	}
}

// Lors des tests, ignorer les erreurs est une erreur. Aussi par rapport à un test sur un comportement précis si une erreur survient
// alors il faut utiliser la fonction t.Fatalf pour stopper immédiatement le test.

// 1- Si une erreur est destinée à signaler quelque chose à l'utilisateur, une erreur sentinelle ne sera probablement pas
// très utile, car sa valeur est fixe. Nous ne pouvons pas l'utiliser pour transmettre des informations dynamiques susceptibles d'aider
// l'utilisateur à résoudre le problème.
// 2- Si le système n'a besoin de rien savoir de l'erreur, si ce n'est qu'elle n'est pas nulle, alors une erreur sentinelle n'est pas
// nécessaire du tout.
