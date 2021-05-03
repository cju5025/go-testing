package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sums an array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := SumArray(numbers)
		want := 15

		if got != want {
			t.Errorf("given %v - got %d, want %d", numbers, got, want)
		}
	})

	t.Run("sums elements of a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlice(numbers)
		want := 6

		if got != want {
			t.Errorf("given %v - got %d, want %d", numbers, got, want)
		}
	})
}

// go test -cover to see coverage percent

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
