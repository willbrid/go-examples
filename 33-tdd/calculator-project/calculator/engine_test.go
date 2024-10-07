package calculator_test

import (
	"calculator-project/calculator"

	"log"
	"os"
	"testing"
)

func init() {
	log.Println("Init setup.")
}

func TestMain(m *testing.M) {
	setup()

	e := m.Run()

	teardown()

	os.Exit(e)
}

func setup() {
	log.Println("Setting up.")
}

func teardown() {
	log.Println("Teardown up.")
}

func TestAdd(t *testing.T) {
	defer func() {
		log.Println("Deferred tearing down.")
	}()

	// Arrange
	e := calculator.Engine{}

	actAssert := func(x, y, want float64) {
		got := e.Add(x, y)

		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 2.5, 3.5
		want := 6.0

		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.5, -3.5
		want := -6.0

		actAssert(x, y, want)
	})
}

func TestSubtract(t *testing.T) {
	e := calculator.Engine{}

	actAssert := func(x, y, want float64) {
		got := e.Subtract(x, y)

		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 7.0, 5.0
		want := 2.0

		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -7.0, -5.0
		want := -2.0

		actAssert(x, y, want)
	})
}

func TestMultiply(t *testing.T) {
	e := calculator.Engine{}

	actAssert := func(x, y, want float64) {
		got := e.Multiply(x, y)

		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 2.0, 5.0
		want := 10.0

		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -2.0, -5.0
		want := 10.0

		actAssert(x, y, want)
	})
}

func TestDivide(t *testing.T) {
	e := calculator.Engine{}

	actAssert := func(x, y, want float64) {
		got := e.Divide(x, y)

		if got != want {
			t.Errorf("Add(%.2f,%.2f) incorrect, got: %.2f, want: %.2f", x, y, got, want)
		}
	}

	t.Run("positive input", func(t *testing.T) {
		x, y := 10.0, 2.0
		want := 5.0

		actAssert(x, y, want)
	})

	t.Run("negative input", func(t *testing.T) {
		x, y := -10.0, -2.0
		want := 5.0

		actAssert(x, y, want)
	})
}
