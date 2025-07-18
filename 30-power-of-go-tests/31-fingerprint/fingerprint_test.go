package fingerprint_test

import (
	"testing"

	"fingerprint"
)

func TestHash_GivesSameUniqueHashForSameData(t *testing.T) {
	t.Parallel()

	data := []byte("These pretzels are making me thirsty.")
	orig := fingerprint.Hash(data)
	same := fingerprint.Hash(data)
	different := fingerprint.Hash([]byte("Hello, Newman"))

	if orig != same {
		t.Error("same data produced different hash")
	}
	if orig == different {
		t.Error("different data produced same hash")
	}
}
