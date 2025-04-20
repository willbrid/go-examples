package even_test

import (
	"strconv"
	"testing"

	"even"
)

func TestIsEven_IsFalseForOddNumbers(t *testing.T) {
	t.Parallel()
	for i := 1; i < 100; i += 2 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if even.IsEven(i) {
				t.Error(true)
			}
		})
	}
}
