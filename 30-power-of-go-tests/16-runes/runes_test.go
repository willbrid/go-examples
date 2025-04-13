package runes_test

import (
	"testing"
	"unicode/utf8"

	"runes"
)

func FuzzFirstRune(f *testing.F) {
	f.Add("Hello")
	f.Add("world")

	f.Fuzz(func(t *testing.T, s string) {
		got := runes.FirstRune(s)
		want, _ := utf8.DecodeRuneInString(s)

		if want == utf8.RuneError {
			t.Skip()
		}
		if want != got {
			t.Errorf("given %q (0x%[1]x): want '%c' (0x%[2]x)", s, want)
			t.Errorf("got '%c' (0x%[1]x)", got)
		}
	})
}

// Supposons, par exemple, que nous souhaitions écrire une fonction FirstRune qui prenne une chaîne et renvoie sa première rune.
// Étant donnée la chaîne < Hello >, par exemple, elle devrait renvoyer la rune < H >. Dans un langage de programmation où les runes
// ne sont que des octets, ce serait simple. Nous pourrions simplement traiter la chaîne comme un octet [] et renvoyer son premier élément.
// Nous savons que cela ne fonctionnera pas en Go, à moins de garantir que l'entrée est limitée au texte ASCII.
// Nous pouvons utiliser des tests de fuzz pour détecter ce type d'erreur.

// Le corpus du fuzzer est l'ensemble des données d'entraînement que nous lui fournissons. Il peut toujours effectuer du fuzzing même avec un
// corpus vide, mais il est utile de lui fournir quelques exemples de valeurs de départ. Ces valeurs seront transmises à notre cible de fuzzing
// et testées. Le fuzzer générera ensuite d'autres entrées à partir de ces valeurs en les modifiant aléatoirement.

// Le corpus inclut également tous les cas de test précédemment générés et échoués, stockés dans le dossier testdata/fuzz. Au fil du temps,
// nous allons donc constituer un corpus de plus en plus volumineux, ce qui permettra au fuzzer de travailler avec davantage de ressources.

// Lors du fuzzing d'une fonction comme FirstRune, toutes les chaînes générées aléatoirement ne sont pas forcément valides, notamment
// du point de vue de l'encodage UTF-8. Par exemple, une chaîne vide ou une séquence d’octets non valides ne représente pas une rune correcte.
// Pour éviter de tester des entrées inutiles, il faut préfiltrer les chaînes invalides. Ainsi si la fonction utf8.DecodeRuneInString
// retourne  une constante < utf8.RuneError >, cela signifie que la chaîne n’est pas une entrée valide, et on ignore ce cas avec < t.Skip >.

// Lorsqu’un test échoue à cause d’une chaîne d’entrée, il est important d’afficher cette chaîne fautive. Mais comme certaines runes
// peuvent ne pas être imprimables, on affiche aussi les octets bruts de la chaîne en hexadécimal avec %x.
// t.Errorf("étant donné %q (0x%x) ...", s, s)
// Mais cela répète deux fois la même variable s. Pour éviter cela, on peut utiliser un index d’argument explicite dans le format, comme
// t.Errorf("given %q (0x%[1]x): want '%c' (0x%[2]x)", s, want)

// Nous pouvons également exécuter le fuzzer pendant une durée déterminée, en utilisant l'option -fuzztime. Par exemple :
// go test -fuzz . -fuzztime=10m

// Il s'avère que le fuzzer ne se contente pas de générer des entrées aléatoires. Il analyse également les chemins de code exécutés par
// ces entrées. Il considère une entrée comme < intéressante > si elle provoque l'exécution d'un nouveau chemin de code dans le système testé.
// Après tout, si un chemin de code n'est exécuté par aucun autre cas de test, des bugs pourraient bien s'y cacher. Le fuzzer tente donc
// de découvrir les entrées qui provoquent l'exécution de nouveaux chemins de code. S'il en trouve, il les utilise comme base pour générer
// des entrées similaires, espérant en trouver une qui déclenche un échec.
