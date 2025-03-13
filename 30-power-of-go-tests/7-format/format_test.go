package format_test

import (
	"testing"

	"format"
)

const (
	invalidInput        string = "invalid input"
	validInput          string = "valid input"
	validInputFormatted string = "valid input formatted"
)

func TestFormatData_ErrorsOnInvalidInput(t *testing.T) {
	t.Parallel()

	_, err := format.FormatData(invalidInput)
	if err == nil {
		t.Error("want error for invalid input")
	}
}

func TestFormatData_IsCorrectForValidInput(t *testing.T) {
	t.Parallel()

	want := validInputFormatted
	got, err := format.FormatData(validInput)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

// Au minimum, nos tests doivent vérifier deux choses :
// premièrement, que le système génère une erreur lorsqu’il le devrait;
// deuxièmement, qu’il n’en génère pas lorsqu’il ne le devrait pas.

//- Pour le test TestFormatData_ErrorsOnInvalidInput, nous ignorons l’autre valeur de retour en utilisant l’identifiant vide _.
// Nous ne prévoyons pas de vérifier cette valeur, car la seule chose qui importe à ce test est que err ne soit pas nul dans ces circonstances.

//- Pour le test TestFormatData_IsCorrectForValidInput même si nous n’attendons pas d’erreur dans ce cas, nous la recevons et la vérifions
// quand même, et nous utilisons t.Fatal pour court-circuiter le test si elle n’est pas nulle. En effet, ce test s’intéresse non seulement à ce
// que la fonction ne renvoie pas d’erreur, mais aussi à ce qu’elle renvoie le bon résultat.

// Vérifier l’erreur ici permet de détecter les bogues dans le comportement de la fonction, ce qui est important, mais ce n’est pas la
// seule raison de le faire. Il est également toujours possible que le test soit erroné. Un bug dans le système est déjà assez grave,
// mais un bug dans le test est bien pire. Cela signifie que nous ne pouvons plus détecter de manière fiable les bugs du système.
