package user_test

import (
	"errors"
	"testing"

	"user"
)

func TestFindUser_GivesErrUserNotFoundForBogusUser(t *testing.T) {
	t.Parallel()

	_, err := user.FindUser("bogus user")
	if !errors.As(err, &user.ErrUserNotFound{}) {
		t.Errorf("wrong error: %v", err)
	}
}

// Création d'un custom error (ErrUserNotFound) et comparaison avec la fonction errors.As qui recherche la première erreur dans l'arbre
// d'erreurs correspondant à la cible. Si elle est trouvée, il définit la cible sur cette valeur d'erreur et renvoie true.
// Sinon, il renvoie « false ».

// Les types d'erreurs personnalisés fonctionnent de manière similaire à l'encapsulation des erreurs avec fmt.Errorf et le verbe %w :
// ils permettent d'ajouter des informations dynamiques tout en conservant la possibilité de comparer l'erreur à une sentinelle
// (ici, un type sentinelle).

// Cependant, cette approche présente quelques inconvénients. Définir une structure distincte et une méthode Error associée pour
// chaque type d'erreur peut rapidement devenir fastidieux et redondant, surtout lorsqu'il y en a un grand nombre.

// À l'inverse, l'encapsulation des erreurs est bien plus simple : il suffit d'utiliser fmt.Errorf en remplaçant le verbe %v par %w.
// Un autre avantage de l'encapsulation est qu'une même erreur peut être encapsulée plusieurs fois tout en permettant à errors.Is de
// retrouver la sentinelle d'origine, ce qui n'est pas possible avec les types d'erreurs personnalisés.
