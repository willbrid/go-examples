package tps_test

import (
	"bytes"
	"os"
	"testing"

	"tps"
)

func TestWriteReportFile_ProducesCorrectOutputFile(t *testing.T) {
	t.Parallel()

	output := t.TempDir() + "/" + t.Name()
	tps.WriteReportFile(output, "Hello world !")

	want, err := os.ReadFile("testdata/output.want")
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

// Les fichiers sont souvent utilisés comme fichiers de référence dans les tests.
// Un fichier de référence contient la sortie attendue d'un processus, et nous pouvons la comparer à la sortie réelle
// pour vérifier son exactitude.
