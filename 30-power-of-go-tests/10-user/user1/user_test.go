package user_test

import (
	"errors"
	"testing"

	"user"
)

func TestFindUser_GivesErrUserNotFoundForBogusUser(t *testing.T) {
	t.Parallel()

	_, err := user.FindUser("bogus user")
	if !errors.Is(err, user.ErrUserNotFound) {
		t.Errorf("wrong error: %v", err)
	}
}

// Le verbe %w indique à fmt.Errorf de créer une erreur encapsulée : une erreur qui contient les informations dynamiques que nous lui
// avons fournies, mais qui se souvient également de l'erreur d'origine: ErrUserNotFound.
// Encapsuler des erreurs de cette manière est toujours sûr, même si le code à l'autre extrémité ignore que l'erreur contient des
// informations supplémentaires. Si vous comparez cette valeur d'erreur à nil, ou si vous l'affichez, elle se comporte comme
// une erreur normale.
// Cependant, il existe un moyen de désencapsuler l'erreur et de récupérer sa valeur sentinelle d'origine : la fonction errors.Is

// Ce mécanisme d'encapsulation des erreurs est élégant, car il permet d'inclure des informations dynamiques utiles dans une erreur,
// tout en facilitant la détermination dans le code du type d'erreur qu'elle représente.
