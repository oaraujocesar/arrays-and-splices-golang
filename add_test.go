package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5}

	result := Add(numbers)
	expected := 15

	if result != expected {
		t.Errorf("Evaluated %d, expected %d, given %v", result, expected, numbers)
	}
}

func ExampleAdd() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Add(numbers))
	// Output: 15
}
