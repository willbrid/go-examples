package calculator_test

import (
	"calculator-project/calculator"

	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.5
	want := 6.0

	got := e.Add(x, y)

	if got != want {
		t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestSubtract(t *testing.T) {
	e := calculator.Engine{}
	x, y := 7.0, 5.0
	want := 2.0

	got := e.Subtract(x, y)

	if got != want {
		t.Errorf("Subtract(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestMultiply(t *testing.T) {
	e := calculator.Engine{}
	x, y := 2.0, 5.0
	want := 10.0

	got := e.Multiply(x, y)

	if got != want {
		t.Errorf("Multiply(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}

func TestDivide(t *testing.T) {
	e := calculator.Engine{}
	x, y := 10.0, 2.0
	want := 5.0

	got := e.Divide(x, y)

	if got != want {
		t.Errorf("Divide(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
	}
}
