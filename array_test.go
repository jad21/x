package x_test

import (
	"reflect"
	"testing"

	. "github.com/jad21/x"
)

func TestFind(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	cb := func(item int) bool { return item == 3 }

	result := Find(array, cb)
	if result == nil || *result != 3 {
		t.Errorf("Find() = %v; want 3", result)
	}

	cb = func(item int) bool { return item == 6 }
	result = Find(array, cb)
	if result != nil {
		t.Errorf("Find() = %v; want nil", result)
	}
}

func TestReverse(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	result := Reverse(array)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse() = %v; want %v", result, expected)
	}
}

func TestContains(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}

	if !Contains(array, 3) {
		t.Errorf("Contains() = false; want true")
	}

	if Contains(array, 6) {
		t.Errorf("Contains() = true; want false")
	}
}

func TestMap(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	f := func(i int) int { return i * 2 }
	expected := []int{2, 4, 6, 8, 10}

	result := Map(array, f)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v; want %v", result, expected)
	}
}

func TestRepeat(t *testing.T) {
	value := 5
	count := 3
	expected := []int{5, 5, 5}

	result := Repeat(value, count)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Repeat() = %v; want %v", result, expected)
	}
}

func TestUnique(t *testing.T) {
	array := []int{1, 2, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	result := Unique(array)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Unique() = %v; want %v", result, expected)
	}
}
