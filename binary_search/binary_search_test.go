package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15}
	target := 5
	actual := BinarySearch(list, target, 0, len(list)-1)
	expected := 4

	if actual < len(list) && list[actual] == target {
		if actual != expected {
			t.Errorf("got %v\nwant %v", actual, expected)
		}
	} else {
		t.Errorf("search failed. got %v want %v", actual, expected)
	}
}

func TestBinarySearchOver(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15}
  expected := 11
  actual := BinarySearch(list, 9999, 0, len(list)-1)

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

func TestBinarySearchTooMin(t *testing.T) {
  list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15}
  expected := 0
  actual := BinarySearch(list, 0, 0, len(list)-1)

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

// func Benchmark_Factorial(b *testing.B){
//   for i := 0; i < b.N; i++ {
//     Factorial(i)
//   }
// }
