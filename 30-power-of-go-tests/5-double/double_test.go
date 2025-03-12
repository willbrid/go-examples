package double_test

import (
	"fmt"
	"testing"

	"double"
)

func TestDouble_2Return4(t *testing.T) {
	t.Parallel()

	want := 4
	got := double.Double(2)
	if want != got {
		t.Errorf("Double(2): want %d, got %d", want, got)
	}
}

func ExampleDouble() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_with2() {
	fmt.Println(double.Double(2))
	// Output:
	// 4
}

func ExampleDouble_with3() {
	fmt.Println(double.Double(3))
	// Output:
	// 6
}

// Nous souhaitons pouvoir écrire du code qui sera inclus dans la documentation autogénérée et dont le comportement pourra être vérifié
// automatiquement, de la même manière que les tests. C'est exactement ce que propose Go, avec une fonctionnalité
// appelée exemples exécutables.
// Un exemple exécutable est similaire à une fonction de test, mais en plus simple. Son nom doit commencer par Example,
// mais il ne prend aucun paramètre et ne renvoie rien.
// Si nous souhaitons écrire un exemple pour l'ensemble du package, plutôt que pour une fonction spécifique,
// nommons-le simplement "Example" :
// func Example() {
//    this demonstrates how to use the entire package
//	  ...
//}
