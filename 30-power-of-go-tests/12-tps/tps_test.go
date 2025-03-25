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

// Gestion des fins de ligne multiplateformes
// Une solution consiste à ajouter un fichier .gitattributes à la racine de notre référentiel ou dans le repertoire testdata,
// avec le contenu suivant : * -text
// Ceci indique à Git de traiter tous les fichiers (*) du référentiel ou du repertoire testdata comme non textuels (-text).
// Autrement dit, de traiter tous les fichiers comme binaires, sans traduction des fins de ligne.
// Cela devrait résoudre le problème, et les utilisateurs disposant de versions à jour de Windows et d'éditeurs de texte
// comme le Bloc-notes pourront facilement modifier les fichiers s'ils le souhaitent.
