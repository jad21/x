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

// TestSortInt prueba la función Sort con un array de enteros
func TestSortInt(t *testing.T) {
	numbers := []int{5, 3, 4, 1, 2}
	expected := []int{1, 2, 3, 4, 5}

	Sort(numbers, func(i, j int) bool {
		return i < j
	})

	if !reflect.DeepEqual(numbers, expected) {
		t.Errorf("expected %v, got %v", expected, numbers)
	}
}

// TestSortString prueba la función Sort con un array de strings
func TestSortString(t *testing.T) {
	words := []string{"banana", "apple", "cherry"}
	expected := []string{"apple", "banana", "cherry"}

	Sort(words, func(i, j string) bool {
		return i < j
	})

	if !reflect.DeepEqual(words, expected) {
		t.Errorf("expected %v, got %v", expected, words)
	}
}

// TestSortStruct prueba la función Sort con un array de estructuras personalizadas
func TestSortStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	expected := []Person{
		{"Bob", 25},
		{"Alice", 30},
		{"Charlie", 35},
	}

	Sort(people, func(i, j Person) bool {
		return i.Age < j.Age
	})

	if !reflect.DeepEqual(people, expected) {
		t.Errorf("expected %v, got %v", expected, people)
	}
}

// TestFilterInt prueba la función Filter con un array de enteros
func TestFilterInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{4, 5}

	result := Filter(numbers, func(n int) bool {
		return n > 3
	})

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestFilterString prueba la función Filter con un array de strings
func TestFilterString(t *testing.T) {
	words := []string{"go", "is", "awesome"}
	longerThanTwo := func(s string) bool {
		return len(s) > 2
	}
	expected := []string{"awesome"}
	result := Filter(words, longerThanTwo)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestUniqueBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
		{"Alice", 25},
	}

	uniqueByName := UniqueBy(people, func(p Person) string {
		return p.Name
	})
	expectedByName := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 30},
	}
	if !reflect.DeepEqual(uniqueByName, expectedByName) {
		t.Errorf("expected %v, got %v", expectedByName, uniqueByName)
	}

	uniqueByAge := UniqueBy(people, func(p Person) int {
		return p.Age
	})
	expectedByAge := []Person{
		{"Alice", 30},
		{"Bob", 25},
	}
	if !reflect.DeepEqual(uniqueByAge, expectedByAge) {
		t.Errorf("expected %v, got %v", expectedByAge, uniqueByAge)
	}
}
