package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Evaluate the sum of all the values of an Array", func(t *testing.T) {
		numbers := [...]int{1, 2, 3, 4, 5}

		result := Add(numbers)
		expected := 15

		if result != expected {
			t.Errorf("Evaluated %d, expected %d, given %v", result, expected, numbers)
		}
	})

	t.Run("Evaluate the sum of all the values of a Slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		result := AddSlice(numbers)
		expected := 21

		if result != expected {
			t.Errorf("Evaluated %d, expected %d, given %v", result, expected, numbers)
		}
	})
}

func ExampleAdd() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Add(numbers))
	// Output: 15
}

func ExampleAddSlice() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(AddSlice(numbers))
	// Output: 21
}
