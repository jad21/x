package x_test

import (
	"fmt"
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

func TestSwitchMap(t *testing.T) {
	cases := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	tests := []struct {
		key          int
		expected     string
		defaultValue string
	}{
		{key: 1, expected: "one", defaultValue: "default"},
		{key: 2, expected: "two", defaultValue: "default"},
		{key: 3, expected: "three", defaultValue: "default"},
		{key: 4, expected: "default", defaultValue: "default"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("key=%d", tt.key), func(t *testing.T) {
			result := SwitchMap(tt.key, cases, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("SwitchMap(%d) = %v; want %v", tt.key, result, tt.expected)
			}
		})
	}
}
