package store_test

import (
	"testing"

	"store"

	"github.com/google/go-cmp/cmp"
)

func TestOpen_ReturnsStoreObject(t *testing.T) {
	want := &store.Store{}
	got, err := store.Open("testdata/godlen.txt") // Le chemin correct du fichier est : testdata/golden.txt

	if err != nil {
		t.Fatalf("unexpected error reading golden file: %v", err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

// Lors des tests, ignorer les erreurs est une erreur. Aussi par rapport à un test sur un comportement précis si une erreur survient
// alors il faut utiliser la fonction t.Fatalf pour stopper immédiatement le test.
