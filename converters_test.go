package x_test

import (
	"testing"

	. "github.com/jad21/x"
)

func TestToListInterfaces(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}

	result := ToListInterfaces(array)

	if len(result) != len(array) {
		t.Errorf("ToListInterfaces() = %v; want same len", result)
	}
}
