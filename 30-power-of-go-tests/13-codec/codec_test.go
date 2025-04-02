package codec_test

import (
	"math/rand/v2"
	"testing"

	"codec"
)

func TestEncodeFollowedByDecode_GivesStartingValue(t *testing.T) {
	t.Parallel()

	input := rand.IntN(10) + 1
	encoded := codec.Encode(input)
	t.Logf("encoded value: %#v", encoded)
	want := input
	got := codec.Decode(encoded)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

// Lorsqu'on teste des fonctions comme Encoder et Decoder, qui devraient être réversibles (encoder une valeur puis la décoder doit
// donner la valeur d'origine), on peut utiliser des entrées de test aléatoires.
// Cependant, si les valeurs générées sont réellement aléatoires, le test peut devenir instable : un bug déclenché uniquement par
// certaines entrées pourrait parfois passer inaperçu.
// Pour éviter cela, on peut fixer la graine du générateur de nombres aléatoires. Ainsi, la séquence de nombres générés restera
// toujours la même, rendant les tests déterministes et reproductibles. Une approche consiste à initialiser un générateur
// aléatoire spécifique aux tests avec une graine fixe, comme 1.
// var rng = rand.New(rand.NewSource(1))
// rng.IntN(10) + 1

// Permutation aléatoire d'un ensemble d'entrées connues
// Une bonne façon d'utiliser l'aléatoire sans provoquer de tests aléatoires est de permuter un ensemble d'entrées, c'est-à-dire de
// les réorganiser, dans un ordre aléatoire.
// Par exemple, le code suivant génère une tranche contenant les entiers de 0 à 99, ordonnés aléatoirement :
// inputs := rand.Perm(100)
