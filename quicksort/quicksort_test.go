package QuickSort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	actual := []int{3, 4, 5, 99, 1, 2, 3000, 5}
	expected := make([]int, len(actual))
	copy(expected, actual)
	QuickSort(actual, 0, len(actual)-1)
	sort.Sort(sort.IntSlice(expected))

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestRandArray(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	actual := []int{}
	for i := 0; i < 9999999; i++ {
		actual = append(actual, rand.Intn(100000))
	}
	expected := make([]int, len(actual))
	copy(expected, actual)

	QuickSort(actual, 0, len(actual)-1)
	sort.Sort(sort.IntSlice(expected))

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	fmt.Println(actual)
}

// func Benchmark_Factorial(b *testing.B){
//   for i := 0; i < b.N; i++ {
//     Factorial(i)
//   }
// }
