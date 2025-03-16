package reader_test

import (
	"io"
	"testing"

	"reader"
)

type errReader struct{}

func (e errReader) Read([]byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}

func TestReadAll_ReturnsAnyReadError(t *testing.T) {
	input := errReader{}

	_, err := reader.ReadAll(input)
	if err == nil {
		t.Error("want error for broken reader, got nil")
	}
}

// Il peut souvent être difficile de créer une valeur d'input pour implémenter un test précis.
// Dans notre cas nous souhaitons tester que notre fonction ReadAll retourne une erreur à partir d'un valeur invalide d'input.
// L'astuce effectuée sera de créer nous-même un type d'input qui retournera toujour une erreur.

// En fait, nous n'avons même pas besoin d'implémenter ce reader (errReader) d'erreurs nous-mêmes. Il est fourni dans le package iotest de
// la bibliothèque standard sous le nom ErrReader, avec d'autres readers de test utiles tels que TimeoutReader et HalfReader.
