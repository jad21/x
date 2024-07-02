package x_test

import (
	"testing"

	. "github.com/jad21/x"
)

func TestIf(t *testing.T) {
	if got := If(true, "yes", "no"); got != "yes" {
		t.Errorf("If(true, 'yes', 'no') = %v; want 'yes'", got)
	}

	if got := If(false, "yes", "no"); got != "no" {
		t.Errorf("If(false, 'yes', 'no') = %v; want 'no'", got)
	}

	if got := If(true, 1, 2); got != 1 {
		t.Errorf("If(true, 1, 2) = %v; want 1", got)
	}

	if got := If(false, 1, 2); got != 2 {
		t.Errorf("If(false, 1, 2) = %v; want 2", got)
	}
}
