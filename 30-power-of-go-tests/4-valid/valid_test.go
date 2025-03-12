package valid_test

import (
	"testing"

	"valid"
)

func TestValid_IsTrueForValidInput(t *testing.T) {
	t.Parallel()

	if !valid.Valid("valid input") {
		t.Error(false)
	}
}

func TestValid_IsFalseForInvalidInput(t *testing.T) {
	t.Parallel()

	if valid.Valid("invalid input") {
		t.Error(true)
	}
}

// Un nom de test doit décrire le comportement du système que le test vérifie. En d’autres termes, un nom du test est le comportement qu’il
// est conçu pour réfuter.

// Les noms de test doivent être ACE c.a.d ils doivent inclure : Action, Condition et Expectation.
// Par exemple, dans le test de notre fonction Valid, ils sont les suivants :
// - Action : appel de Valid
// - Condition : avec une entrée valide
// - Expectation : renvoie true

// Il est possible de partir d’une intention très claire, mais de finir par écrire un test qui ne fait pas exactement ce que nous voulions.

// Détecteur de bugs : doit détecter non seulement les bugs que nous écrivons aujourd’hui, mais aussi les bugs que nous,
// ou quelqu’un d’autre, pourrions écrire à l’avenir.

// Il existe un outil appelé gotestdox qui exécute les tests et signale les résultats tout en traduisant simultanément
// les noms des tests en phrases lisibles.
//
// Installation du binaire gotestdox
// go install github.com/bitfield/gotestdox/cmd/gotestdox@latest
//
// Execution
// gotestdox ./...
//
// Exemple de résultat
// valid:
//  ✔ Valid is false for invalid input (0.00s)
//  ✔ Valid is true for valid input (0.00s)
//
// Nous pouvons garder nos phrases de test courtes et concises en omettant des mots comme should, must et will, puis
// décrire simplement ce qu'elles font.
// En d'autres termes, une unité bien conçue ne devrait pas avoir plus de comportement que ce qui peut être exprimé en quelques
// phrases courtes, chacune pouvant être directement traduite en test. L'information contenue dans une seule phrase doit correspondre
// assez bien au comportement qu'une unité devrait avoir.
//
// Une bonne conception ne se résume pas seulement à ce que nous avons l'imagination d'y intégrer, mais aussi à ce que
// nous avons la clarté d'intention d'omettre.
// Abstractions superficielles: Est-ce que ça en vaut vraiment la peine ?
//
// Ne nous soucions pas des noms de test longs.
// Créer des messages d'échec informatifs
// if err != nil {
//     t.Error("input %d caused unexpected error: %v", input, err)
// }
//
// if want != got {
// 	   t.Errorf("want %v, got %v", want, got)
// }
