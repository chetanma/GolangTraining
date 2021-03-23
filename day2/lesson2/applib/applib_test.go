package applib

import "testing"

func TestVersion(t *testing.T) {
	ver := Version()

	if ver > 1.1 {
		t.Error("Incorrect version")

	}
}
